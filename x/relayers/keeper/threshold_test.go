package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
	"github.com/stafiprotocol/stafihub/x/relayers/keeper"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNThreshold(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Threshold {
	items := make([]types.Threshold, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
        
		keeper.SetThreshold(ctx, items[i])
	}
	return items
}

func TestThresholdGet(t *testing.T) {
	keeper, ctx := keepertest.RelayersKeeper(t)
	items := createNThreshold(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetThreshold(ctx,
		    item.Index,
            
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestThresholdRemove(t *testing.T) {
	keeper, ctx := keepertest.RelayersKeeper(t)
	items := createNThreshold(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveThreshold(ctx,
		    item.Index,
            
		)
		_, found := keeper.GetThreshold(ctx,
		    item.Index,
            
		)
		require.False(t, found)
	}
}

func TestThresholdGetAll(t *testing.T) {
	keeper, ctx := keepertest.RelayersKeeper(t)
	items := createNThreshold(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllThreshold(ctx))
}
