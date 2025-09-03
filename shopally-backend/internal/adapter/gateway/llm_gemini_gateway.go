package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/shopally-ai/internal/contextkeys"
	"github.com/shopally-ai/pkg/domain"
	"github.com/shopally-ai/pkg/util"
)

// GeminiLLMGateway implements domain.LLMGateway using Google Generative Language API (Gemini).
type GeminiLLMGateway struct {
	apiKey   string
	modelURL string
	client   *http.Client
	fx       domain.IFXClient
}

// CompareProducts implements domain.LLMGateway.
// CompareProducts implements domain.LLMGateway.
// CompareProducts implements domain.LLMGateway.
func (g *GeminiLLMGateway) CompareProducts(ctx context.Context, productDetails []*domain.Product) (*domain.ComparisonResult, error) {
	if len(productDetails) < 2 {
		return nil, fmt.Errorf("at least one product is required")
	}

	req := struct {
		Products []*domain.Product `json:"products"`
	}{Products: productDetails}

	b, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal products: %w", err)
	}

	lang := "en"
	if v := ctx.Value(contextkeys.RespLang); v != nil {
		if s, ok := v.(string); ok && s != "" {
			lang = s
		}
	}

	// Extract delivery info from first product for the prompt
	deliveryInfo := "varies"
	if len(productDetails) > 0 && productDetails[0].DeliveryEstimate != "" {
		deliveryInfo = productDetails[0].DeliveryEstimate
	}

	prompt := fmt.Sprintf(`You are an expert e-commerce product comparison assistant. Analyze and compare %d products thoroughly.

CRITICAL INSTRUCTIONS:
1. Return STRICT JSON ONLY, no additional text or commentary
2. JSON structure must be exactly this format:
{
  "products": [
    {
      "product": { /* complete original product object for product 1 */ },
      "synthesis": {
        "pros": ["advantage 1", "advantage 2", ...],
        "cons": ["disadvantage 1", "disadvantage 2", ...],
        "isBestValue": true/false,
        "features": {
          "priceScore": "excellent/good/fair/poor",
          "qualityScore": "excellent/good/fair/poor",
          "valueForMoney": "excellent/good/fair/poor",
          "deliveryRating": "excellent/good/fair/poor"
        }
      }
    },
    {
      "product": { /* complete original product object for product 2 */ },
      "synthesis": {
        "pros": ["advantage 1", "advantage 2", ...],
        "cons": ["disadvantage 1", "disadvantage 2", ...],
        "isBestValue": true/false,
        "features": {
          "priceScore": "excellent/good/fair/poor",
          "qualityScore": "excellent/good/fair/poor",
          "valueForMoney": "excellent/good/fair/poor",
          "deliveryRating": "excellent/good/fair/poor"
        }
      }
    }
    /* include all %d products in the same format */
  ]
}

FEATURE COMPARISON GUIDELINES:
- For each product, identify features that MAKE IT DIFFERENT from others
- Highlight what makes ONE PRODUCT BETTER than others in specific areas
- Focus on COMPETITIVE ADVANTAGES and UNIQUE SELLING POINTS
- Compare relative strengths and weaknesses across products

SPECIFIC AREAS TO COMPARE:
1. PRICE: Which offers best value? Consider ETB price, USD price, discounts, tax
2. QUALITY: Compare product ratings (%v/5 vs others) and customer reviews
3. SELLER: Compare seller scores (/100) and trust indicators
4. DELIVERY: Compare delivery estimates ("%s" vs others)
5. POPULARITY: Compare number sold (%d units) and market acceptance
6. FEATURES: Compare summary bullets and unique offerings


RESPONSE LANGUAGE: %s

PRODUCTS DATA:
%s`, len(productDetails), len(productDetails), len(productDetails), deliveryInfo, len(productDetails), lang, string(b))

	// Call LLM
	text, err := g.call(ctx, prompt)
	log.Println("LLM comparison error:", err)

	if err != nil {
		return nil, fmt.Errorf("LLM API call failed: %w", err)
	}

	clean := extractJSON(text)

	// Parse into ComparisonResult structure
	var result domain.ComparisonResult
	if err := json.Unmarshal([]byte(clean), &result); err != nil {
		return nil, fmt.Errorf("failed to parse LLM response: %w", err)
	}

	// Validate that we got the expected number of comparisons
	if len(result.Products) != len(productDetails) {
		return nil, fmt.Errorf("LLM returned %d comparisons but expected %d", len(result.Products), len(productDetails))
	}

	return &result, nil
}

// NewGeminiLLMGateway creates a new gateway using the GEMINI_API_KEY from env if apiKey is empty.
func NewGeminiLLMGateway(apiKey string, fx domain.IFXClient) domain.LLMGateway {
	if apiKey == "" {
		apiKey = os.Getenv("GEMINI_API_KEY")
	}

	return &GeminiLLMGateway{
		apiKey:   apiKey,
		modelURL: "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent",
		client:   &http.Client{Timeout: 12 * time.Second},
		fx:       fx,
	}
}

func (g *GeminiLLMGateway) call(ctx context.Context, prompt string) (string, error) {
	if g.apiKey == "" {
		return "", errors.New("missing GEMINI_API_KEY")
	}
	reqBody := domain.GeminiRequest{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{Parts: []struct {
				Text string `json:"text"`
			}{{Text: prompt}}},
		},
	}
	b, _ := json.Marshal(reqBody)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, g.modelURL+"?key="+g.apiKey, bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := g.client.Do(req)

	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", errors.New("gemini http status: " + resp.Status)
	}
	var gr domain.GeminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&gr); err != nil {
		return "", err
	}
	// Concatenate all parts to avoid returning partial code-fenced blocks like "```json"
	for _, c := range gr.Candidates {
		var b strings.Builder
		for _, p := range c.Content.Parts {
			if t := strings.TrimSpace(p.Text); t != "" {
				if b.Len() > 0 {
					b.WriteString("\n")
				}
				b.WriteString(t)
			}
		}
		if b.Len() > 0 {
			return b.String(), nil
		}
	}

	log.Println("[GeminiLLMGateway] Warning: empty response from Gemini API")

	return "", errors.New("gemini empty response")
}

// ParseIntent asks the model to extract a structured JSON of constraints.
// ParseIntent asks the model to extract a structured JSON of constraints.
// ParseIntent asks the model to extract a structured JSON of constraints.
func (g *GeminiLLMGateway) ParseIntent(ctx context.Context, query string) (map[string]interface{}, error) {
	requestID := ""
	if requestID == "" {
		requestID = "unknown"
	}

	normalizedQuery := strings.TrimSpace(query)

	// 2) Content moderation: Check for potentially harmful content
	if isPotentiallyHarmful(normalizedQuery) {
		log.Printf("[%s] Blocked query due to potentially harmful content: %s", requestID, normalizedQuery)
		return nil, errors.New("query contains potentially harmful or prohibited content")
	}

	prompt := fmt.Sprintf(`STRICT INSTRUCTIONS: OUTPUT ONLY RAW JSON, NO OTHER TEXT, NO EXPLANATIONS, NO CODE BLOCKS.

You are an advanced e-commerce search intent parser with expert budget extraction capabilities.

## CRITICAL RULES:
1. OUTPUT ONLY PURE JSON - NO OTHER TEXT WHATSOEVER
2. EXTRACT AND CONVERT NUMBER WORDS TO DIGITS (e.g., "five hundred" → 500, "ሁለት ሺህ" → 2000)
3. PRESERVE ORIGINAL CURRENCY VALUES - DO NOT CONVERT CURRENCIES
4. DETECT BOTH MINIMUM AND MAXIMUM PRICE RANGES
5. OUTPUT NUMERICAL VALUES EXACTLY AS EXTRACTED

## BUDGET EXTRACTION RULES:
- Convert number words to digits: "five" → 5, "twenty" → 20, "one hundred" → 100
- Convert Amharic numbers: "አንድ" → 1, "ሁለት" → 2, "አምስት" → 5, "መቶ" → 100, "ሺህ" → 1000
- Detect ranges: "between 100 and 200" → min=100, max=200
- Detect inequalities: "under 500", "less than 1000" → max=500, max=1000
- Detect inequalities: "over 200", "more than 300" → min=200, min=300
- Detect approximate: "around 1500", "about 2000" → use as max_sale_price
- Extract exact numbers: "500 ETB" → preserve as 500, is_etb=true

## CURRENCY HANDLING:
- PRESERVE ORIGINAL NUMERICAL VALUES - NO CURRENCY CONVERSION
- If user specifies "ETB", "birr", "ብር" → is_etb = true
- If user specifies "$", "USD", "dollars" → is_etb = false  
- If no currency specified → is_etb = true (default to ETB)
- Output prices as exact numbers without conversion

## LANGUAGE PROCESSING:
- Extract keywords in English regardless of input language
- Translate Amharic product names to English (e.g., "ስልክ" → "phone")
- Convert number words to digits in all languages
- Understand mixed language queries

## JSON SCHEMA:
{
  "keywords": "string",           // English keywords from any language input
  "category_ids": "string|null",
  "min_sale_price": number|null,  // Exact numerical value as mentioned
  "max_sale_price": number|null,  // Exact numerical value as mentioned  
  "delivery_days": number|null,
  "ship_to_country": "ET",
  "target_currency": "USD",
  "target_language": "en",
  "is_etb": boolean,              // Currency context (true=ETB, false=USD)
  "original_budget": "string|null" // Original budget phrase for reference
}

## EXAMPLES (User Query -> JSON Output):

"smartphone under five hundred birr" -> {"keywords":"smartphone","min_sale_price":null,"max_sale_price":500,"category_ids":null,"delivery_days":null,"ship_to_country":"ET","target_currency":"USD","target_language":"en","is_etb":true,"original_budget":"under five hundred birr"}

"laptop between one thousand and two thousand ETB" -> {"keywords":"laptop","min_sale_price":1000,"max_sale_price":2000,"category_ids":null,"delivery_days":null,"ship_to_country":"ET","target_currency":"USD","target_language":"en","is_etb":true,"original_budget":"between one thousand and two thousand ETB"}

"watch over two hundred dollars" -> {"keywords":"watch","min_sale_price":200,"max_sale_price":null,"category_ids":null,"delivery_days":null,"ship_to_country":"ET","target_currency":"USD","target_language":"en","is_etb":false,"original_budget":"over two hundred dollars"}

"ስልክ ከአምስት መቶ እስከ ሺህ ብር" -> {"keywords":"phone","min_sale_price":500,"max_sale_price":1000,"category_ids":null,"delivery_days":null,"ship_to_country":"ET","target_currency":"USD","target_language":"en","is_etb":true,"original_budget":"ከአምስት መቶ እስከ ሺህ ብር"}

"computer around one thousand five hundred" -> {"keywords":"computer","min_sale_price":null,"max_sale_price":1500,"category_ids":null,"delivery_days":null,"ship_to_country":"ET","target_currency":"USD","target_language":"en","is_etb":true,"original_budget":"around one thousand five hundred"}

"expensive shoes under ሁለት ሺህ ብር" -> {"keywords":"shoes","min_sale_price":null,"max_sale_price":2000,"category_ids":null,"delivery_days":null,"ship_to_country":"ET","target_currency":"USD","target_language":"en","is_etb":true,"original_budget":"under ሁለት ሺህ ብር"}

"ተንሸራታች ከአምስት ሺህ ብር በላይ" -> {"keywords":"sneakers","min_sale_price":5000,"max_sale_price":null,"category_ids":null,"delivery_days":null,"ship_to_country":"ET","target_currency":"USD","target_language":"en","is_etb":true,"original_budget":"ከአምስት ሺህ ብር በላይ"}

## NUMBER WORD CONVERSION REFERENCE:
English: one=1, two=2, three=3, four=4, five=5, six=6, seven=7, eight=8, nine=9, ten=10,
         twenty=20, thirty=30, forty=40, fifty=50, sixty=60, seventy=70, eighty=80, ninety=90,
         hundred=100, thousand=1000, million=1000000

Amharic: አንድ=1, ሁለት=2, ሦስት=3, አራት=4, አምስት=5, ስድስት=6, ሰባት=7, ስምንት=8, ዘጠኝ=9, አስር=10,
         ሃያ=20, ሰላሳ=30, አርባ=40, አምሳ=50, ስልሳ=60, ሰባ=70, ሰማንያ=80, ዘጠና=90,
         መቶ=100, ሺህ=1000, ሚሊዮን=1000000

INPUT QUERY: "%s"
OUTPUT:`, normalizedQuery)

	log.Printf("[%s] Sending multi-language JSON prompt to LLM", requestID)

	text, err := g.call(ctx, prompt)
	if err != nil {
		return nil, err
	}

	// Extract and clean JSON
	clean := extractStrictJSON(text)
	log.Printf("[%s] Extracted JSON: %s", requestID, clean)

	// Parse the JSON response directly into map
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(clean), &m); err != nil {
		log.Printf("[%s] Failed to parse LLM JSON response: %v. Raw: %s", requestID, err, clean)
		// Fallback to minimal response with default is_etb = true
		m = map[string]interface{}{
			"keywords":        normalizedQuery,
			"category_ids":    nil,
			"min_sale_price":  nil,
			"max_sale_price":  nil,
			"delivery_days":   nil,
			"ship_to_country": "ET",
			"target_currency": "USD",
			"target_language": "en",
			"is_etb":          true, // Default to ETB
		}
	}

	// convert "min_sale_price" and "max_sale_price" to float64 if they are numbers

	// Enforce required fields
	m["ship_to_country"] = "ET"
	m["target_currency"] = "USD"
	m["target_language"] = "en"

	// Ensure is_etb field exists and is boolean, default to true if missing
	if _, exists := m["is_etb"]; !exists {
		m["is_etb"] = true
	}

	if isETB, ok := m["is_etb"].(bool); ok && isETB {
		// Convert prices from ETB to USD if is_etb is true and prices are present
		if minPrice, ok := m["min_sale_price"].(float64); ok && minPrice > 0 {
			usdPrice, _, err := util.USDToETB(minPrice)
			if err == nil {
				m["min_sale_price"] = minPrice / usdPrice
			}
		}
		if maxPrice, ok := m["max_sale_price"].(float64); ok && maxPrice > 0 {
			usdPrice, _, err := util.USDToETB(maxPrice)
			if err == nil {
				m["max_sale_price"] = maxPrice / usdPrice
			}
		}
	}

	// Ensure keywords exist and are in English (basic fallback)
	if keywords, ok := m["keywords"].(string); !ok || strings.TrimSpace(keywords) == "" {
		// If LLM failed to extract keywords, use original query but this should be rare
		m["keywords"] = normalizedQuery
	} else {
		m["keywords"] = strings.TrimSpace(keywords)
	}

	return m, nil
}

// extractStrictJSON aggressively extracts JSON from LLM response
func extractStrictJSON(s string) string {
	s = strings.TrimSpace(s)

	// Remove code fences and any surrounding text
	if strings.Contains(s, "```") {
		lines := strings.Split(s, "\n")
		var jsonLines []string
		inJson := false

		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "```") {
				if inJson {
					break // End of JSON block
				}
				inJson = true
				continue
			}
			if inJson && trimmed != "" {
				jsonLines = append(jsonLines, line)
			}
		}
		if len(jsonLines) > 0 {
			s = strings.Join(jsonLines, "\n")
		}
	}

	// Try to find JSON object boundaries
	start := strings.Index(s, "{")
	end := strings.LastIndex(s, "}")
	if start != -1 && end != -1 && end > start {
		s = s[start : end+1]
	}

	// Remove any non-JSON content before/after
	s = strings.TrimSpace(s)

	// Basic validation - must start with { and end with }
	if !strings.HasPrefix(s, "{") || !strings.HasSuffix(s, "}") {
		// Fallback: return empty JSON object with is_etb field
		return `{"keywords":null,"category_ids":null,"min_sale_price":null,"max_sale_price":null,"delivery_days":null,"ship_to_country":"ET","target_currency":"USD","target_language":"en","is_etb":true}`
	}

	return s
}

func extractJSON(s string) string {
	s = strings.TrimSpace(strings.ReplaceAll(s, "\r", ""))
	if strings.HasPrefix(s, "```") {
		// Split into lines and remove starting and ending fence lines only
		lines := strings.Split(s, "\n")
		// drop first line (``` or ```json)
		if len(lines) > 0 && strings.HasPrefix(strings.TrimSpace(lines[0]), "```") { // Fixed index here
			lines = lines[1:]
		}
		// drop trailing fence lines (```) if present
		for len(lines) > 0 && strings.HasPrefix(strings.TrimSpace(lines[len(lines)-1]), "```") {
			lines = lines[:len(lines)-1]
		}
		s = strings.Join(lines, "\n")
	}
	return strings.TrimSpace(s)
}

// func fmtInt(i int) string { return fmt.Sprintf("%d", i) }

// isPotentiallyHarmful checks a query for keywords associated with illegal or adult content.
// This is a basic implementation and should be expanded significantly for production use.
func isPotentiallyHarmful(query string) bool {
	lowerQuery := strings.ToLower(query)
	blacklist := []string{
		"drugs", "weapons", "firearms", "explosives", "contraband",
		"porn", "sex toys", "adult content", "erotic", "hentai",
		"illegal", "smuggled", "stolen goods", "counterfeit",
		"hate speech", "violence", "racist", "discriminatory",
	}

	for _, keyword := range blacklist {
		if strings.Contains(lowerQuery, keyword) {
			return true
		}
	}
	return false
}

// Heuristic Amharic detection: Unicode Ethiopic block or common tokens
func (g *GeminiLLMGateway) SummarizeProduct(ctx context.Context, p *domain.Product, userPrompt string) (*domain.Product, error) {
	lang, _ := ctx.Value(contextkeys.RespLang).(string)
	if lang == "" {
		lang = "en"
	}

	// Calculate AI match percentage based on product relevance to user's original query
	aiMatchPercentage := calculateAIMatchPercentage(p, userPrompt)

	// Generate enhanced content in the appropriate language
	enhancedProduct, err := g.enhanceProductContent(ctx, p, userPrompt, lang, aiMatchPercentage)
	log.Println("Enhancement error:", err)
	if err != nil {
		// If enhancement fails, return the original product with basic enhancements
		log.Printf("Product enhancement failed, returning original product: %v", err)
		return g.createBasicEnhancedProduct(p, userPrompt, lang, aiMatchPercentage), nil
	}

	return enhancedProduct, nil
}

func (g *GeminiLLMGateway) enhanceProductContent(ctx context.Context, p *domain.Product, userPrompt, lang string, aiMatchPercentage int) (*domain.Product, error) {
	prompt := fmt.Sprintf(`STRICT INSTRUCTIONS: OUTPUT ONLY RAW JSON, NO OTHER TEXT, NO EXPLANATIONS, NO CODE BLOCKS.

You are an expert e-commerce product content enhancer. Return the COMPLETE product JSON structure with enhanced text content.

## USER'S ORIGINAL REQUEST: "%s"

## LANGUAGE: %s
- Write ALL text fields in %s language
- Use appropriate cultural context

## RULES:
- Output the EXACT product JSON structure
- Enhance text fields to be more engaging and persuasive
- Keep ALL original field names, values, and structure
- Only modify: description, customerHighlights, customerReview, summaryBullets
- All other fields must remain EXACTLY the same
- Numerical values, URLs, IDs must not change

## ORIGINAL PRODUCT DATA:
%s

## ENHANCEMENT GUIDELINES FOR %s:
1. description: Make comprehensive yet engaging (3-4 sentences)
2. customerHighlights: Make more compelling and benefit-focused
3. customerReview: Make more natural and persuasive
4. summaryBullets: Create 3-5 bullet points with ejection-style formatting
5. title: Keep meaning but make more appealing if needed

## REQUIRED OUTPUT:
The complete product JSON with enhanced text fields in %s language.

OUTPUT:`, userPrompt, lang, lang, getProductJSONString(p), strings.ToUpper(lang), lang)

	log.Printf("Enhancing product content for language: %s", lang)

	text, err := g.call(ctx, prompt)
	if err != nil {
		return nil, err
	}

	clean := extractStrictJSON(text)
	log.Printf("Extracted enhanced product JSON: %s", clean)

	// Parse the enhanced product
	var enhancedProduct domain.Product
	if err := json.Unmarshal([]byte(clean), &enhancedProduct); err != nil {
		log.Printf("Failed to parse enhanced product JSON: %v", err)
		return nil, err
	}

	// Ensure critical fields remain unchanged
	enhancedProduct.ID = p.ID
	enhancedProduct.ImageURL = p.ImageURL
	enhancedProduct.AIMatchPercentage = aiMatchPercentage
	enhancedProduct.Price = p.Price
	enhancedProduct.ProductRating = p.ProductRating
	enhancedProduct.SellerScore = p.SellerScore
	enhancedProduct.DeliveryEstimate = p.DeliveryEstimate
	enhancedProduct.NumberSold = p.NumberSold
	enhancedProduct.DeeplinkURL = p.DeeplinkURL
	enhancedProduct.TaxRate = p.TaxRate
	enhancedProduct.Discount = p.Discount

	return &enhancedProduct, nil
}

// getProductJSONString returns the product as a JSON string for the prompt
func getProductJSONString(p *domain.Product) string {
	productMap := map[string]interface{}{
		"id":                 p.ID,
		"title":              p.Title,
		"imageUrl":           p.ImageURL,
		"aiMatchPercentage":  p.AIMatchPercentage,
		"price":              p.Price,
		"productRating":      p.ProductRating,
		"sellerScore":        p.SellerScore,
		"deliveryEstimate":   p.DeliveryEstimate,
		"description":        p.Description,
		"customerHighlights": p.CustomerHighlights,
		"customerReview":     p.CustomerReview,
		"numberSold":         p.NumberSold,
		"summaryBullets":     p.SummaryBullets,
		"deeplinkUrl":        p.DeeplinkURL,
		"taxRate":            p.TaxRate,
		"discount":           p.Discount,
	}

	jsonBytes, _ := json.MarshalIndent(productMap, "", "  ")
	return string(jsonBytes)
}

// createBasicEnhancedProduct creates enhanced content without LLM
func (g *GeminiLLMGateway) createBasicEnhancedProduct(p *domain.Product, userPrompt, lang string, aiMatchPercentage int) *domain.Product {
	enhanced := &domain.Product{
		ID:                 p.ID,
		Title:              p.Title,
		ImageURL:           p.ImageURL,
		AIMatchPercentage:  aiMatchPercentage,
		Price:              p.Price,
		ProductRating:      p.ProductRating,
		SellerScore:        p.SellerScore,
		DeliveryEstimate:   p.DeliveryEstimate,
		Description:        enhanceDescription(p.Description, lang),
		CustomerHighlights: enhanceHighlights(p.CustomerHighlights, lang),
		CustomerReview:     enhanceReview(p.CustomerReview, lang),
		NumberSold:         p.NumberSold,
		SummaryBullets:     createSummaryBullets(p, lang),
		DeeplinkURL:        p.DeeplinkURL,
		TaxRate:            p.TaxRate,
		Discount:           p.Discount,
	}
	return enhanced
}

func enhanceDescription(desc, lang string) string {
	if lang == "am" {
		return "ይህ ምርት በጥራቱ የታወቀ እና በደንበኞች የተወደደ ነው። ከፍተኛ ጥራት ያለው ዲዛይን እና አስተማማኝ አገልግሎት ይገልጻል። በተጠቃሚዎች አወንታዊ አስተያየት የተረጋገጠ የምርት ልምድ ያቀርባል።"
	}
	return "This high-quality product is known for its excellent performance and customer satisfaction. It features durable construction and reliable functionality that users appreciate. The product has received positive feedback for its consistent delivery on promises and overall value."
}

func enhanceHighlights(highlights, lang string) string {
	if lang == "am" {
		return "★ ከፍተኛ ጥራት ያለው ምርት\n→ በደንበኞች የተወደደ\n• አስተማማኝ አፈጻጸም\n→ ዘመናዊ ዲዛይን"
	}
	return "★ High-quality construction\n→ Customer favorite\n• Reliable performance\n→ Modern design"
}

func enhanceReview(review, lang string) string {
	if lang == "am" {
		return "ተጠቃሚዎች ይህን ምርት ለጥራቱ እና አስተማማኝነቱ ያነግራሉ። ከፍተኛ የደንበኛ እርካታ ያለው ምርት ነው።"
	}
	return "Users praise this product for its quality and reliability. It has generated high customer satisfaction and positive feedback across various platforms."
}

func createSummaryBullets(p *domain.Product, lang string) []string {
	if lang == "am" {
		return []string{
			"★ ከፍተኛ ጥራት ያለው ምርት",
			"→ በደንበኞች የተወደደ",
			"• አስተማማኝ አፈጻጸም",
			"→ ዘመናዊ ዲዛይን",
		}
	}
	return []string{
		"★ High-quality product construction",
		"→ Customer favorite with great reviews",
		"• Reliable performance and durability",
		"→ Modern and user-friendly design",
	}
}

// calculateAIMatchPercentage calculates relevance score based on product data and user prompt
func calculateAIMatchPercentage(p *domain.Product, userPrompt string) int {
	score := 0
	userPrompt = strings.ToLower(userPrompt)

	// Text relevance scoring
	score += calculateTextMatchScore(p.Title, userPrompt, 30)
	score += calculateTextMatchScore(p.Description, userPrompt, 20)
	score += calculateTextMatchScore(p.CustomerHighlights+" "+p.CustomerReview, userPrompt, 15)

	// Quality indicators
	if p.ProductRating >= 4.0 {
		score += 10
	}
	if p.SellerScore >= 90 {
		score += 8
	}
	if p.NumberSold > 1000 {
		score += 7
	}

	// Context matching
	if containsBudgetKeywords(userPrompt) {
		score += 10
	}
	if containsDeliveryKeywords(userPrompt) {
		score += 10
	}

	return min(score, 100)
}

func calculateTextMatchScore(text, userPrompt string, maxScore int) int {
	text = strings.ToLower(text)
	words := strings.Fields(userPrompt)
	if len(words) == 0 {
		return 0
	}

	matchCount := 0
	for _, word := range words {
		if len(word) > 3 && strings.Contains(text, word) {
			matchCount++
		}
	}

	return int(float64(matchCount) / float64(len(words)) * float64(maxScore))
}

func containsBudgetKeywords(prompt string) bool {
	keywords := []string{"price", "cost", "budget", "cheap", "expensive", "affordable", "$", "etb", "birr", "ብር"}
	for _, keyword := range keywords {
		if strings.Contains(prompt, keyword) {
			return true
		}
	}
	return false
}

func containsDeliveryKeywords(prompt string) bool {
	keywords := []string{"delivery", "shipping", "arrive", "receive", "days", "time", "fast", "quick", "slow"}
	for _, keyword := range keywords {
		if strings.Contains(prompt, keyword) {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
