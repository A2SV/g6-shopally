package gateway

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/shopally-ai/pkg/domain"
)

// FXHTTPGateway is an outbound adapter that calls a configurable FX HTTP API.
// It implements usecase.IFXClient.
type FXHTTPGateway struct {
	APIURL     string
	APIKey     string
	HTTPClient *http.Client
}

var _ domain.IFXClient = (*FXHTTPGateway)(nil)

// NewFXHTTPGateway creates a new gateway. If httpClient is nil, a default client is used.
func NewFXHTTPGateway(apiURL, apiKey string, httpClient *http.Client) *FXHTTPGateway {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 15 * time.Second}
	}
	return &FXHTTPGateway{APIURL: apiURL, APIKey: apiKey, HTTPClient: httpClient}
}

// GetRate fetches the conversion rate from -> to using the configured provider URL/key.
// It supports multiple common provider response shapes.
func (g *FXHTTPGateway) GetRate(ctx context.Context, from, to string) (float64, error) {
	if strings.TrimSpace(from) == "" || strings.TrimSpace(to) == "" {
		return 0, errors.New("from/to required")
	}

	reqURL := g.buildRequestURL(strings.ToUpper(from), strings.ToUpper(to))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return 0, err
	}
	resp, err := g.HTTPClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return 0, fmt.Errorf("fx api non-ok: %d - %s", resp.StatusCode, string(body))
	}

	// Try multiple known shapes
	// exchangerate.host convert: {"success":true, "result": 56.78}
	var xhost struct {
		Result float64 `json:"result"`
	}
	if err := json.Unmarshal(body, &xhost); err == nil && xhost.Result != 0 {
		return xhost.Result, nil
	}

	// currencyfreaks latest: {"rates": {"ETB": "56.78"}}
	var cf struct {
		Rates map[string]string `json:"rates"`
	}
	if err := json.Unmarshal(body, &cf); err == nil {
		if v, ok := cf.Rates[strings.ToUpper(to)]; ok {
			if f, perr := parseFloat(v); perr == nil {
				return f, nil
			}
		}
	}

	// open.er-api.com: {"rates": {"ETB": 56.78}}
	var er struct {
		Rates map[string]float64 `json:"rates"`
	}
	if err := json.Unmarshal(body, &er); err == nil {
		if f, ok := er.Rates[strings.ToUpper(to)]; ok && f != 0 {
			return f, nil
		}
	}

	// exchangerate-api.com v6: {"conversion_rates": {"ETB": 56.78}}
	var era struct {
		ConversionRates map[string]float64 `json:"conversion_rates"`
	}
	if err := json.Unmarshal(body, &era); err == nil {
		if f, ok := era.ConversionRates[strings.ToUpper(to)]; ok && f != 0 {
			return f, nil
		}
	}

	return 0, fmt.Errorf("unrecognized fx response for %s", reqURL)
}

func (g *FXHTTPGateway) buildRequestURL(from, to string) string {
	apiURL := g.APIURL
	if apiURL == "" {
		apiURL = "https://api.exchangerate.host/convert?from={FROM}&to={TO}"
	}

	// Replace common placeholders
	u := apiURL
	u = strings.ReplaceAll(u, "{FROM}", url.QueryEscape(from))
	u = strings.ReplaceAll(u, "{TO}", url.QueryEscape(to))
	u = strings.ReplaceAll(u, "{APIKEY}", url.QueryEscape(g.APIKey))

	// Provider-specific fallbacks/normalizations
	if strings.Contains(u, "exchangerate.host") && !strings.Contains(u, "from=") {
		v := url.Values{}
		v.Set("from", from)
		v.Set("to", to)
		u = fmt.Sprintf("https://api.exchangerate.host/convert?%s", v.Encode())
	}

	if strings.Contains(u, "currencyfreaks.com") && !strings.Contains(u, "apikey=") {
		v := url.Values{}
		if g.APIKey != "" {
			v.Set("apikey", g.APIKey)
		}
		v.Set("symbols", to)
		u = fmt.Sprintf("https://api.currencyfreaks.com/latest?%s", v.Encode())
	}

	if strings.Contains(u, "open.er-api.com") && !strings.Contains(u, "/latest/") {
		u = fmt.Sprintf("https://open.er-api.com/v6/latest/%s", from)
	}

	if strings.Contains(u, "v6.exchangerate-api.com") && !strings.Contains(u, "/latest/") {
		u = fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/%s", url.PathEscape(g.APIKey), from)
	}

	return u
}

func parseFloat(s string) (float64, error) {
	s = strings.TrimSpace(strings.Trim(s, "\"'"))
	if s == "" {
		return 0, errors.New("empty number string")
	}
	return strconv.ParseFloat(s, 64)
}
