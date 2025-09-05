package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/shopally-ai/internal/contextkeys"
	"github.com/shopally-ai/pkg/domain"
	"github.com/shopally-ai/pkg/util"
)

// GeminiLLMGateway implements domain.LLMGateway using Google Generative Language API (Gemini).
type GeminiLLMGateway struct {
	tokenManager *util.TokenManager
	modelURL     string
	client       *http.Client
	fx           domain.IFXClient
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

	log.Println("CompareProducts: calling LLM with", len(productDetails), "products in lang:", ctx.Value(contextkeys.RespLang).(string))
	lang, _ := ctx.Value(contextkeys.RespLang).(string)
	if lang == "" {
		lang = "en"
	}

	// Extract delivery info from first product for the prompt
	deliveryInfo := "varies"
	if len(productDetails) > 0 && productDetails[0].DeliveryEstimate != "" {
		deliveryInfo = productDetails[0].DeliveryEstimate
	}

	prompt := fmt.Sprintf(`You are an expert e-commerce product comparison assistant. Analyze and compare %d products thoroughly.

CRITICAL INSTRUCTIONS:
1. Return STRICT JSON ONLY, no additional text or commentary
2. JSON structure must exactly match the format below
3. If language is Amharic ("am"), **translate both feature keys and descriptive text to Amharic**

{
  "products": [
    {
      "product": { "id":"123", "title":"Sample Product", "imageUrl":"https://example.com/image.jpg", "price":{"etb":1000,"usd":20,"fxTimestamp":"2025-09-03T12:00:00Z"}, "productRating":4.5, "sellerScore":95, "deliveryEstimate":"3-5 days", "description":"Sample description", "summaryBullets":["bullet1","bullet2"], "deeplinkUrl":"https://example.com/product/123", "taxRate":0.1, "discount":10, "numberSold":150, "aiMatchPercentage":0 },
      "synthesis": {
        English (en)
		"pros": [
			"Affordable price",
			"High quality materials",
			"Fast delivery"
		],
		"cons": [
			"Limited color options"
		]

		Amharic (am)
		"pros": [
			"ተመጣጣኝ ዋጋ",
			"ከፍተኛ ጥራት ያላቸው ቁሳቁሶች",
			"ፈጣን እና ታማኝ አሰራር"
		],
		"cons": [
			"የቀለም አማራጮች ገደብ ተደርጓል"
		]
        "isBestValue": true,
        "features": {
          // English keys if lang="en"
          "Price & Value": "Cheaper than competitors with excellent value",
          "Quality & Durability": "Solid build with premium materials and long-lasting",
          "Seller Trust": "Highly rated seller with good reputation",
          "Delivery Speed": "Faster than most competitors",
          "Popularity & Demand": "Well-liked with high sales volume",
          "Unique Features": "Supports wireless charging and extra features",

          // Amharic keys if lang="am"
          "ዋጋ እና የዋጋ እኩልነት": "ከተወዳጅ ምርቶች በተመጣጣኝ ዋጋ ይገኛል",
          "ጥራት እና ቆይታ": "ጥሩ እና ረጅም ቆይታ ያለው ጥራት ተሸማች",
          "የሻጭ እርግጠኝነት": "ከፍተኛ ደረጃ ያለው ታማኝ ሻጭ",
          "የአሰራር ፍጥነት": "ከአብዛኛዎቹ ምርቶች ፈጣን የማድረስ ጊዜ",
          "ታዋቂነት እና ጥያቄ": "በከፍተኛ ብዛት የተሸጠ የታወቀ ምርት",
          "ብቸኛ ስለሆነ ባህሪዎች": "ያለ ግድ የሚከናወን ማስተካከያ ያለው እና ተጨማሪ ባህሪዎች ያሉበት"
        }
      }
    }
    /* repeat above block for all %d products */
  ],
  "overallComparison": {
    "bestValueProduct": "Sample Product",
    "bestValueLink": "https://example.com/product/123",
    "bestValuePrice": {
      "etb": 1000,
      "usd": 20,
      "fxTimestamp": "2025-09-03T12:00:00Z"
    },
    "keyHighlights": [
      "Most cost-effective option",
      "High-quality materials and fast delivery"
    ],
    "summary": "Sample Product offers the best value, balancing price, quality, and speed, while competitors may offer slightly better features but at higher costs."
  }
}

FEATURE COMPARISON GUIDELINES:
- Compare products relative to each other
- Highlight strengths, weaknesses, trade-offs, and unique selling points
- Use descriptive, human-readable phrases
- Tone: analytical, descriptive, detail-oriented, persuasive
- If lang='am', feature keys AND values must be in Amharic

SPECIFIC AREAS TO COMPARE:
1. PRICE: ETB, USD, discounts, tax
2. QUALITY: product ratings (%v/5) and customer reviews
3. SELLER: seller scores (/100) and trust indicators
4. DELIVERY: delivery estimates ("%s")
5. POPULARITY: number sold (%d units)
6. FEATURES: summary bullets and distinctive advantages

RESPONSE LANGUAGE: %s
- 'am' → provide all text, including feature keys, in Amharic
- 'en' → provide all text, including feature keys, in English

PRODUCTS DATA:
%s`, len(productDetails), len(productDetails), len(productDetails), deliveryInfo, len(productDetails), lang, string(b))

	// Call LLM
	text, err := g.call(context.Background(), prompt)
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
func NewGeminiLLMGateway(tm *util.TokenManager, fx domain.IFXClient) domain.LLMGateway {

	return &GeminiLLMGateway{
		tokenManager: tm,
		modelURL:     "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent",
		client:       &http.Client{Timeout: 12 * time.Second},
		fx:           fx,
	}
}

func (g *GeminiLLMGateway) call(ctx context.Context, prompt string) (string, error) {
	// tokenManager must be initialized with at least one token
	if g.tokenManager == nil {
		return "", errors.New("no Gemini API keys available")
	}
	token := g.tokenManager.GetNextToken()

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
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, g.modelURL+"?key="+token, bytes.NewReader(b))
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

You are an advanced multi-language e-commerce intent parser. Your task is to normalize user queries into structured JSON for product search.

## CRITICAL RULES:
1. OUTPUT ONLY PURE JSON
2. DETECT CURRENCY MENTIONED IN QUERY:
   - If query mentions "USD", "$", or "dollars" → set is_etb=false
   - If query mentions "ETB", "birr", "ብር" or no currency → set is_etb=true
3. COLLECT ALL TERMS (main product, synonyms, categories, brands, brand codes, gender, usage/function)
4. STEM WORDS TO BASE FORM (e.g., "running" → "run", "shoes" → "shoe")
5. TRANSLATE PRODUCT TERMS TO ENGLISH
6. MERGE EVERYTHING INTO A SINGLE SPACE-SEPARATED STRING → keywords
7. INCLUDE GENDER (male, female, unisex, kid) IF SPECIFIED
8. PRESERVE BRAND NAMES AND MODEL NUMBERS (e.g., "Nike AirMax 270" → "nike airmax 270")
9. CONVERT NUMBER WORDS (English & Amharic) TO DIGITS
10. PRESERVE ORIGINAL BUDGET PHRASE

## JSON SCHEMA:
{
  "keywords": "string",            // single string, space-separated, stemmed, includes all relevant terms
  "min_sale_price": number|null,   // budget lower bound
  "max_sale_price": number|null,   // budget upper bound
  "original_budget": "string|null",// raw budget phrase
  "delivery_days": number|null,    // expected delivery time if mentioned
  "ship_to_country": "ET",
  "target_currency": "USD",
  "target_language": "en",
  "is_etb": boolean                // true = ETB, false = USD
}

## EXAMPLES:

"user query: red nike running shoes under 3000 birr" ->
{
  "keywords": "red nike run shoe sneaker sport trainer footwear male",
  "min_sale_price": null,
  "max_sale_price": 3000,
  "original_budget": "under 3000 birr",
  "delivery_days": null,
  "ship_to_country": "ET",
  "target_currency": "USD",
  "target_language": "en",
  "is_etb": true
}

"user query: cheap gaming laptop around one thousand dollars" ->
{
  "keywords": "cheap game laptop notebook computer pc electronic",
  "min_sale_price": null,
  "max_sale_price": 1000,
  "original_budget": "around one thousand dollars",
  "delivery_days": null,
  "ship_to_country": "ET",
  "target_currency": "USD",
  "target_language": "en",
  "is_etb": false
}

"user query: ነጭ ቀሚስ የወንድ ከአምስት መቶ ብር በታች" ->
{
  "keywords": "white shirt men male clothing apparel fashion",
  "min_sale_price": null,
  "max_sale_price": 500,
  "original_budget": "ከአምስት መቶ ብር በታች",
  "delivery_days": null,
  "ship_to_country": "ET",
  "target_currency": "USD",
  "target_language": "en",
  "is_etb": true
}

"user query: samsung galaxy s23 ultra phone under 700 dollars" ->
{
  "keywords": "samsung galaxy s23 ultra phone smartphone mobile",
  "min_sale_price": null,
  "max_sale_price": 700,
  "original_budget": "under 700 dollars",
  "delivery_days": null,
  "ship_to_country": "ET",
  "target_currency": "USD",
  "target_language": "en",
  "is_etb": false
}

INPUT QUERY: "%s"
OUTPUT:`, normalizedQuery)

	log.Printf("[%s] Sending multi-language JSON prompt to LLM", requestID)

	text, err := g.call(context.TODO(), prompt)
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

	// Generate enhanced content in the appropriate language
	enhancedProduct, err := g.enhanceProductContent(ctx, p, userPrompt, lang)
	log.Println("Enhancement error:", err)
	if err != nil {
		// If enhancement fails, return the original product with basic enhancements
		log.Printf("Product enhancement failed, returning original product: %v", err)
		return g.createBasicEnhancedProduct(p, lang), nil
	}

	return enhancedProduct, nil
}

func (g *GeminiLLMGateway) enhanceProductContent(ctx context.Context, p *domain.Product, userPrompt, lang string) (*domain.Product, error) {
	prompt := fmt.Sprintf(`STRICT INSTRUCTIONS: OUTPUT ONLY RAW JSON. 
	NO explanations, NO extra text, NO markdown, NO code blocks. 

	You are an expert e-commerce product content enhancer and product relevance evaluator. 
	Your tasks are: (1) evaluate product relevance to the user’s prompt, (2) enhance content if relevant.

	======================================================================
	## USER INPUT
	- USER PROMPT: "%s"
	- TARGET LANGUAGE: %s

	======================================================================
	## PRODUCT DATA
	%s

	======================================================================
	## DECISION LOGIC
	1. First, evaluate RELEVANCE of this product to the USER PROMPT.
	- If product is **irrelevant, misleading, or unrelated**, output only:
		{ "removeProduct": true, "aiMatchPercentage": 0 }
	- Do NOT attempt to enhance such products.
	- Examples of irrelevance: wrong category, unrelated use-case, incorrect features, or major mismatch.

	2. If product is relevant:
	- Enhance the following text fields ONLY:
		- title
		- description
		- summaryBullets
	- Ensure text is persuasive, engaging, culturally appropriate, and localized to the target language.
	- Keep it concise but compelling, highlighting benefits and differentiators.
	- Do NOT fabricate features, prices, URLs, or attributes that are not present in the input product data.

	3. Always generate a numerical score:
	- aiMatchPercentage (0 to 100)
	- Based on closeness to the user's prompt, product details, features, and quality.
	- High match (>70): strong relevance
	- Medium match (40–70): somewhat relevant but not perfect
	- Low match (<30): not relevant → mark removeProduct=true

	======================================================================
	## OUTPUT RULES
	- Strictly JSON only. No additional commentary.
	- Keep numerical values, IDs, URLs, and other structured data unchanged.
	- Enhanced fields must not contradict existing data.
	- summaryBullets must contain 3–5 short, punchy points.

	======================================================================
	## REQUIRED OUTPUT FORMAT
	{
	"title": "string",
	"description": "string",
	"summaryBullets": ["string", ...],
	"aiMatchPercentage": number,   // 0–100
	"removeProduct": bool          // true if not relevant
	}
	======================================================================

	OUTPUT:`, userPrompt, lang, getProductJSONString(p))

	text, err := g.call(ctx, prompt)
	if err != nil {
		return nil, err
	}

	clean := extractStrictJSON(text)

	// Temporary struct to detect if product should be removed
	type aiResponse struct {
		Title          string   `json:"title"`
		Description    string   `json:"description"`
		SummaryBullets []string `json:"summaryBullets"`
		AiMatchPercent int      `json:"aiMatchPercentage"`
		RemoveProduct  bool     `json:"removeProduct"`
	}

	var resp aiResponse
	if err := json.Unmarshal([]byte(clean), &resp); err != nil {
		log.Printf("Failed to parse enhanced product JSON: %v", err)
		return nil, err
	}

	// If AI determined product is irrelevant → skip it
	if resp.RemoveProduct || resp.AiMatchPercent < 30 {
		log.Printf("Product marked as irrelevant (aiMatchPercentage=%d). Removing...", resp.AiMatchPercent)
		return nil, nil
	}

	// Build enhanced product
	enhancedProduct := *p
	enhancedProduct.Title = resp.Title
	enhancedProduct.Description = resp.Description
	enhancedProduct.SummaryBullets = resp.SummaryBullets
	enhancedProduct.AIMatchPercentage = resp.AiMatchPercent

	// Keep critical fields unchanged
	enhancedProduct.ID = p.ID
	enhancedProduct.ImageURL = p.ImageURL
	enhancedProduct.Price = p.Price
	enhancedProduct.ProductRating = p.ProductRating
	enhancedProduct.DeliveryEstimate = p.DeliveryEstimate
	enhancedProduct.DeeplinkURL = p.DeeplinkURL
	enhancedProduct.TaxRate = p.TaxRate
	enhancedProduct.Discount = p.Discount

	return &enhancedProduct, nil
}

// getProductJSONString returns the product as a JSON string for the prompt
func getProductJSONString(p *domain.Product) string {
	productMap := map[string]interface{}{
		"id":                p.ID,
		"title":             p.Title,
		"imageUrl":          p.ImageURL,
		"aiMatchPercentage": p.AIMatchPercentage,
		"price":             p.Price,
		"productRating":     p.ProductRating,
		"deliveryEstimate":  p.DeliveryEstimate,
		"description":       p.Description,
		"summaryBullets":    p.SummaryBullets,
		"deeplinkUrl":       p.DeeplinkURL,
		"taxRate":           p.TaxRate,
		"discount":          p.Discount,
	}

	jsonBytes, _ := json.MarshalIndent(productMap, "", "  ")
	return string(jsonBytes)
}

// createBasicEnhancedProduct creates enhanced content without LLM
func (g *GeminiLLMGateway) createBasicEnhancedProduct(p *domain.Product, lang string) *domain.Product {
	enhanced := &domain.Product{
		ID:               p.ID,
		Title:            p.Title,
		ImageURL:         p.ImageURL,
		Price:            p.Price,
		ProductRating:    p.ProductRating / 20,
		DeliveryEstimate: p.DeliveryEstimate,
		Description:      enhanceDescription(p.Description, lang),
		SummaryBullets:   createSummaryBullets(lang),
		DeeplinkURL:      p.DeeplinkURL,
		TaxRate:          p.TaxRate,
		Discount:         p.Discount,
	}
	return enhanced
}

func enhanceDescription(desc, lang string) string {
	if lang == "am" {
		return "ይህ ምርት በጥራቱ የታወቀ እና በደንበኞች የተወደደ ነው። ከፍተኛ ጥራት ያለው ዲዛይን እና አስተማማኝ አገልግሎት ይገልጻል። በተጠቃሚዎች አወንታዊ አስተያየት የተረጋገጠ የምርት ልምድ ያቀርባል።"
	}
	return "This high-quality product is known for its excellent performance and customer satisfaction. It features durable construction and reliable functionality that users appreciate. The product has received positive feedback for its consistent delivery on promises and overall value."
}

func createSummaryBullets(lang string) []string {
	if lang == "am" {
		return []string{
			"ከፍተኛ ጥራት ያለው ምርት",
			"በደንበኞች የተወደደ",
			"አስተማማኝ አፈጻጸም",
			"ዘመናዊ ዲዛይን",
		}
	}
	return []string{
		"High-quality product construction",
		"Customer favorite with great reviews",
		"Reliable performance and durability",
		"Modern and user-friendly design",
	}
}
