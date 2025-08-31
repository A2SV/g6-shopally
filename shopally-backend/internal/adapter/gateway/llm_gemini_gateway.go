package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/shopally-ai/internal/contextkeys"
	"github.com/shopally-ai/pkg/domain"
)

// GeminiLLMGateway implements domain.LLMGateway using Google Generative Language API (Gemini).
type GeminiLLMGateway struct {
	apiKey   string
	modelURL string
	client   *http.Client
	fx       domain.IFXClient
}

// CompareProducts implements domain.LLMGateway.
func (g *GeminiLLMGateway) CompareProducts(ctx context.Context, productDetails []*domain.Product) (map[string]interface{}, error) {
	panic("unimplemented")
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

type geminiRequest struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

type geminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func (g *GeminiLLMGateway) call(ctx context.Context, prompt string) (string, error) {
	if g.apiKey == "" {
		return "", errors.New("missing GEMINI_API_KEY")
	}
	reqBody := geminiRequest{Contents: []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	}{
		{Parts: []struct {
			Text string `json:"text"`
		}{{Text: prompt}}},
	}}
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
	var gr geminiResponse
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
	return "", errors.New("gemini empty response")
}

// ParseIntent asks the model to extract a structured JSON of constraints.
func (g *GeminiLLMGateway) ParseIntent(ctx context.Context, query string) (map[string]interface{}, error) {
	// 1) Language normalization: if Amharic present, translate to English first
	normalized := strings.TrimSpace(query)
	if isLikelyAmharic(normalized) {
		tprompt := "Translate the following user shopping query into concise English only (no extra words):\n" + normalized
		tr, err := g.call(ctx, tprompt)
		if err == nil && strings.TrimSpace(tr) != "" {
			normalized = strings.TrimSpace(extractJSON(tr))
		}
	}

	// 2) Currency normalization: detect ETB amounts and convert to USD via fx
	// Default values
	var approxUSD float64
	var haveUSD bool
	if amt, cur, ok := extractAmountCurrency(normalized); ok && strings.EqualFold(cur, "ETB") && g.fx != nil {
		if rate, err := g.fx.GetRate(ctx, "ETB", "USD"); err == nil && rate > 0 {
			approxUSD = amt * rate
			haveUSD = true
		}
	}

	// 3) Build strict parsing prompt
	// Always enforce target conditions: language = "en", currency = "USD", ship_to = "Ethiopia"
	sb := &strings.Builder{}
	sb.WriteString("You are an intent parser for e-commerce search.\n")
	sb.WriteString("Task: Output ONLY compact JSON (no prose, no code fences) with keys: \n")
	sb.WriteString("  category (string),\n  min_price (number|null),\n  max_price (number|null),\n  delivery_days_max (number|null),\n  ship_to (string),\n  currency (string),\n  language (string).\n")
	sb.WriteString("Rules: ship_to must be 'Ethiopia'; currency must be 'USD'; language must be 'en'.\n")
	if haveUSD {
		fmt.Fprintf(sb, "Normalized hint: approx_budget_usd=%.2f.\n", approxUSD)
	}
	sb.WriteString("User query (English-normalized if needed): \n")
	sb.WriteString(normalized)

	text, err := g.call(ctx, sb.String())
	if err != nil {
		return nil, err
	}
	// Extract JSON from possible code fences
	clean := extractJSON(text)
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(clean), &m); err != nil {
		// fallback minimal
		m = map[string]interface{}{"category": "", "min_price": nil, "max_price": nil}
	}
	// Hard enforce required fields per business rules
	m["ship_to"] = "Ethiopia"
	m["currency"] = "USD"
	m["language"] = "en"
	if _, ok := m["max_price"]; !ok && haveUSD {
		// Provide a rounded price if we computed it (USD)
		m["max_price"] = round2(approxUSD)
	}
	return m, nil
}

// SummarizeProduct requests a single JSON string summary from provided product data.
func (g *GeminiLLMGateway) SummarizeProduct(ctx context.Context, p *domain.Product) ([]string, error) {
	lang, _ := ctx.Value(contextkeys.RespLang).(string)
	if lang == "" {
		lang = "en"
	}

	instr := "You are a product summarization assistant. " +
		"Using ONLY the provided fields (Title, Description, CustomerHighlights, CustomerReview, SellerScore, ProductRating), " +
		"write a short natural-language product summary (2–3 sentences). " +
		"Incorporate information from CustomerHighlights, CustomerReview, and ProductRating to describe the product features and strengths naturally, " +
		"but do NOT mention phrases like 'customers say', 'appreciate', or explicitly state the rating. " +
		"Do NOT include Price, DeliveryEstimate, or NumberOfSoldItems. " +
		"Do NOT invent or infer any details beyond the provided fields. "

	if lang == "am" {
		instr += "Write in Amharic (am). "
	} else {
		instr += "Write in English (en). "
	}

	instr += "Output format: strictly a JSON array with exactly ONE element: the product summary string. " +
		"Ensure the JSON is syntactically valid."

	base := instr +
		" Title: " + p.Title +
		", Description: " + p.Description +
		", CustomerHighlights: " + p.CustomerHighlights +
		", SellerScore: " + fmtInt(p.SellerScore)

	text, err := g.call(ctx, base)
	if err != nil {
		return nil, err
	}

	clean := extractJSON(text)

	// ✅ Unmarshal directly into []string
	var out []string
	if err := json.Unmarshal([]byte(clean), &out); err == nil && len(out) > 0 {
		return out, nil
	}

	// fallback: extract first non-empty line if JSON parse fails
	lines := strings.Split(strings.ReplaceAll(clean, "\r", ""), "\n")
	for _, ln := range lines {
		ln = strings.TrimSpace(ln)
		if ln == "" || strings.HasPrefix(ln, "```") {
			continue
		}
		if len(ln) >= 2 && ((ln[0] == '"' && ln[len(ln)-1] == '"') || (ln[0] == '\'' && ln[len(ln)-1] == '\'')) {
			ln = strings.Trim(ln, "'\"")
		}
		if ln != "" {
			return []string{ln}, nil
		}
	}

	return []string{"Summary unavailable"}, nil
}

func extractJSON(s string) string {
	s = strings.TrimSpace(strings.ReplaceAll(s, "\r", ""))
	if strings.HasPrefix(s, "```") {
		// Split into lines and remove starting and ending fence lines only
		lines := strings.Split(s, "\n")
		// drop first line (``` or ```json)
		if len(lines) > 0 && strings.HasPrefix(strings.TrimSpace(lines[0]), "```") {
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

func fmtInt(i int) string { return fmt.Sprintf("%d", i) }

// Heuristic Amharic detection: Unicode Ethiopic block or common tokens
func isLikelyAmharic(s string) bool {
	s = strings.TrimSpace(s)
	for _, r := range s {
		if r >= 0x1200 && r <= 0x137F {
			return true
		}
	}
	low := strings.ToLower(s)
	if strings.Contains(low, "ብር") || strings.Contains(low, "ኢትዮ") || strings.Contains(low, "አማር") {
		return true
	}
	return false
}

var priceRe = regexp.MustCompile(`(?i)(?:\b(?:etb|birr|ብር)\b)\s*([0-9]+(?:[.,][0-9]{1,2})?)|([0-9]+(?:[.,][0-9]{1,2})?)\s*\b(?:etb|birr|ብር)\b`)

// extractAmountCurrency finds a single ETB price in the string, returns numeric amount and currency
func extractAmountCurrency(s string) (float64, string, bool) {
	m := priceRe.FindStringSubmatch(s)
	if len(m) == 0 {
		return 0, "", false
	}
	num := m[1]
	if num == "" && len(m) > 2 {
		num = m[2]
	}
	num = strings.ReplaceAll(num, ",", "")
	f, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, "", false
	}
	return f, "ETB", true
}

func round2(f float64) float64 { return math.Round(f*100) / 100 }
