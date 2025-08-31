package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/shopally-ai/pkg/domain"
)

type FXHandler struct {
	FX domain.IFXClient
}

func NewFXHandler(fx domain.IFXClient) *FXHandler {
	return &FXHandler{FX: fx}
}

type fxResponse struct {
	From      string  `json:"from"`
	To        string  `json:"to"`
	Rate      float64 `json:"rate"`
	Amount    float64 `json:"amount"`
	Converted float64 `json:"converted"`
}

func (h *FXHandler) GetFX(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	from := strings.ToUpper(strings.Trim(strings.TrimSpace(r.URL.Query().Get("from")), "\"'"))
	to := strings.ToUpper(strings.Trim(strings.TrimSpace(r.URL.Query().Get("to")), "\"'"))
	if from == "" {
		from = "USD"
	}
	if to == "" {
		to = "ETB"
	}

	amount := 1.0
	if s := strings.TrimSpace(r.URL.Query().Get("amount")); s != "" {
		a, err := strconv.ParseFloat(s, 64)
		if err != nil || a < 0 {
			http.Error(w, "invalid amount", http.StatusBadRequest)
			return
		}
		amount = a
	}

	rate, err := h.FX.GetRate(ctx, from, to)
	if err != nil {
		http.Error(w, "fx error: "+err.Error(), http.StatusBadGateway)
		return
	}

	resp := fxResponse{
		From:      from,
		To:        to,
		Rate:      rate,
		Amount:    amount,
		Converted: rate * amount,
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
