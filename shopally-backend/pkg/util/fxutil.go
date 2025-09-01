package util

import (
	"context"
	"errors"
	"strconv"

	"github.com/shopally-ai/pkg/domain"
)

// FXKeyUSDToETB is the canonical cache key used to store the USD->ETB rate.
// It matches the key shape used by the cached FX client: "fx:USD:ETB".
const FXKeyUSDToETB = "fx:USD:ETB"

// USDToETB converts a USD amount to ETB using the rate stored in cache.
//
// Contract:
// - Reads the latest USD->ETB FX rate from the cache at FXKeyUSDToETB.
// - Returns (convertedETB, rateUsed, error).
// - If the key is missing or unparsable, returns an error.
func USDToETB(ctx context.Context, usd float64, cache domain.ICachePort) (float64, float64, error) {
	if cache == nil {
		return 0, 0, errors.New("cache is nil")
	}
	val, ok, err := cache.Get(ctx, FXKeyUSDToETB)
	if err != nil {
		return 0, 0, err
	}
	if !ok {
		return 0, 0, errors.New("usd->etb rate not found in cache")
	}
	rate, perr := strconv.ParseFloat(val, 64)
	if perr != nil {
		return 0, 0, perr
	}
	return usd * rate, rate, nil
}
