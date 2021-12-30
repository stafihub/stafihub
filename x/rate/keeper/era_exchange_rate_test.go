package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
	"github.com/stafiprotocol/stafihub/x/rate/keeper"
	"github.com/stafiprotocol/stafihub/x/rate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNEraExchangeRate(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.EraExchangeRate {
	items := make([]types.EraExchangeRate, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetEraExchangeRate(ctx, items[i])
	}
	return items
}

func TestEraExchangeRateRemove(t *testing.T) {
	keeper, ctx := keepertest.RateKeeper(t)
	items := createNEraExchangeRate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEraExchangeRate(ctx,
		    item.Index,

		)
		_, found := keeper.GetEraExchangeRate(ctx,
		    item.Index,

		)
		require.False(t, found)
	}
}

func TestEraExchangeRateGetAll(t *testing.T) {
	keeper, ctx := keepertest.RateKeeper(t)
	items := createNEraExchangeRate(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllEraExchangeRate(ctx))
}
