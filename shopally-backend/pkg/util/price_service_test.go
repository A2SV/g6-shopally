package util_test

import (
	"context"
	"errors"
	"testing"

	"github.com/shopally-ai/internal/testmocks"
	"github.com/shopally-ai/pkg/util"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUpdatePriceIfChanged_Changed(t *testing.T) {
	ff := testmocks.NewPriceFetcher(t)
	ff.On("FetchPrices", mock.Anything, []string{"p1"}).Return(
		map[string]util.PriceAmounts{"p1": {USD: 20}}, nil,
	).Once()

	svc := util.NewWithFetcher(ff)
	updated, changed, err := svc.UpdatePriceIfChanged(context.Background(), "p1", 10.0)
	require.NoError(t, err)
	require.Equal(t, 20.0, updated)
	require.True(t, changed)
}

func TestUpdatePriceIfChanged_Unchanged(t *testing.T) {
	ff := testmocks.NewPriceFetcher(t)
	ff.On("FetchPrices", mock.Anything, []string{"p2"}).Return(
		map[string]util.PriceAmounts{"p2": {USD: 15.5}}, nil,
	).Once()

	svc := util.NewWithFetcher(ff)
	updated, changed, err := svc.UpdatePriceIfChanged(context.Background(), "p2", 15.5)
	require.NoError(t, err)
	require.Equal(t, 15.5, updated)
	require.False(t, changed)
}

func TestUpdatePriceIfChanged_ProductNotFound(t *testing.T) {
	ff := testmocks.NewPriceFetcher(t)
	ff.On("FetchPrices", mock.Anything, []string{"missing"}).Return(
		map[string]util.PriceAmounts{}, nil,
	).Once()

	svc := util.NewWithFetcher(ff)
	_, _, err := svc.UpdatePriceIfChanged(context.Background(), "missing", 0)
	require.Error(t, err)
}

func TestUpdatePriceIfChanged_GatewayError(t *testing.T) {
	ff := testmocks.NewPriceFetcher(t)
	ff.On("FetchPrices", mock.Anything, []string{"p3"}).Return(
		nil, errors.New("upstream failure"),
	).Once()

	svc := util.NewWithFetcher(ff)
	_, _, err := svc.UpdatePriceIfChanged(context.Background(), "p3", 0)
	require.Error(t, err)
}

func TestUpdatePricesIfChangedBatch_ReturnsFoundPrices(t *testing.T) {
	ff := testmocks.NewPriceFetcher(t)
	// Input: ["x", "y", "y", "z", ""] → After dedupe: ["x", "y", "z"]
	ff.On("FetchPrices", mock.Anything, []string{"x", "y", "z"}).Return(
		map[string]util.PriceAmounts{
			"x": {USD: 2}, "y": {USD: 4}, "z": {USD: 0},
		}, nil,
	).Once()

	svc := util.NewWithFetcher(ff)
	current := map[string]float64{"x": 2, "y": 3}
	res, err := svc.UpdatePricesIfChangedBatch(context.Background(), []string{"x", "y", "y", "z", ""}, current)
	require.NoError(t, err)
	require.Len(t, res, 3) // x, y, z
	require.Equal(t, 2.0, res["x"].Price)
	require.Equal(t, 4.0, res["y"].Price)
	require.Equal(t, 0.0, res["z"].Price)
}

func TestGetCurrentPriceUSDAndETB_Success(t *testing.T) {
	ff := testmocks.NewPriceFetcher(t)
	ff.On("FetchPrices", mock.Anything, []string{"p1"}).Return(
		map[string]util.PriceAmounts{"p1": {USD: 10, ETB: 600}}, nil,
	).Once()

	svc := util.NewWithFetcher(ff)
	usd, etb, err := svc.GetCurrentPriceUSDAndETB(context.Background(), "p1")
	require.NoError(t, err)
	require.Equal(t, 10.0, usd)
	require.Equal(t, 600.0, etb)
}

func TestGetCurrentPriceUSDAndETB_ProductNotFound(t *testing.T) {
	ff := testmocks.NewPriceFetcher(t)
	ff.On("FetchPrices", mock.Anything, []string{"missing"}).Return(
		map[string]util.PriceAmounts{}, nil,
	).Once()

	svc := util.NewWithFetcher(ff)
	_, _, err := svc.GetCurrentPriceUSDAndETB(context.Background(), "missing")
	require.Error(t, err)
}

func TestGetCurrentPricesUSDAndETBBatch_Success(t *testing.T) {
	ff := testmocks.NewPriceFetcher(t)
	// Input: ["a", "b", "c", "a", ""] → After dedupe: ["a", "b", "c"]
	ff.On("FetchPrices", mock.Anything, []string{"a", "b", "c"}).Return(
		map[string]util.PriceAmounts{
			"a": {USD: 3, ETB: 150},
			"b": {USD: 4, ETB: 200},
			"c": {USD: 0, ETB: 0},
		}, nil,
	).Once()

	svc := util.NewWithFetcher(ff)
	out, err := svc.GetCurrentPricesUSDAndETBBatch(context.Background(), []string{"a", "b", "c", "a", ""})
	require.NoError(t, err)
	require.Len(t, out, 3)
	require.Equal(t, 3.0, out["a"].USD)
	require.Equal(t, 150.0, out["a"].ETB)
	require.Equal(t, 4.0, out["b"].USD)
	require.Equal(t, 200.0, out["b"].ETB)
	require.Equal(t, 0.0, out["c"].USD)
	require.Equal(t, 0.0, out["c"].ETB)
}
