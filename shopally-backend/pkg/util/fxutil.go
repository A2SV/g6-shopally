package util

import (
	"context"
	"errors"
	"strconv"
	"sync/atomic"

	"github.com/shopally-ai/pkg/domain"
)

// FXKeyUSDToETB is the canonical cache key used to store the USD->ETB rate.
// It matches the key shape used by the cached FX client: "fx:USD:ETB".
const FXKeyUSDToETB = "fx:USD:ETB"

// Package-level cache and context factory to avoid passing them everywhere.
var (
	fxCache atomic.Value // stores domain.ICachePort
	ctxFac  atomic.Value // stores func() context.Context
)

func init() {
	// default context factory
	ctxFac.Store(func() context.Context { return context.Background() })
}

// SetFXCache sets the cache implementation used by this package.
func SetFXCache(c domain.ICachePort) { fxCache.Store(c) }

// SetContextFactory overrides the context factory used for cache calls.
// If f is nil, it resets to the default background context.
func SetContextFactory(f func() context.Context) {
	if f == nil {
		ctxFac.Store(func() context.Context { return context.Background() })
		return
	}
	ctxFac.Store(f)
}

// USDToETB converts a USD amount to ETB using the rate stored in cache.
//
// Contract:
// - Reads the latest USD->ETB FX rate from the cache at FXKeyUSDToETB.
// - Returns (convertedETB, rateUsed, error).
// - If the key is missing or unparsable, returns an error.
func USDToETB(usd float64) (float64, float64, error) {
	c, _ := fxCache.Load().(domain.ICachePort)
	if c == nil {
		return 0, 0, errors.New("fx cache not initialized")
	}
	getCtx, _ := ctxFac.Load().(func() context.Context)
	if getCtx == nil {
		getCtx = func() context.Context { return context.Background() }
	}
	ctx := getCtx()

	val, ok, err := c.Get(ctx, FXKeyUSDToETB)
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
