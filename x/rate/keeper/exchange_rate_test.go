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

func createNExchangeRate(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ExchangeRate {
	items := make([]types.ExchangeRate, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
        
		keeper.SetExchangeRate(ctx, items[i])
	}
	return items
}

func TestExchangeRateGet(t *testing.T) {
	keeper, ctx := keepertest.RateKeeper(t)
	items := createNExchangeRate(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetExchangeRate(ctx,
		    item.Index,
            
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestExchangeRateRemove(t *testing.T) {
	keeper, ctx := keepertest.RateKeeper(t)
	items := createNExchangeRate(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveExchangeRate(ctx,
		    item.Index,
            
		)
		_, found := keeper.GetExchangeRate(ctx,
		    item.Index,
            
		)
		require.False(t, found)
	}
}

func TestExchangeRateGetAll(t *testing.T) {
	keeper, ctx := keepertest.RateKeeper(t)
	items := createNExchangeRate(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllExchangeRate(ctx))
}
