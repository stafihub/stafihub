package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
	"github.com/stafiprotocol/stafihub/x/relayers/keeper"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRelayer(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Relayer {
	items := make([]types.Relayer, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetRelayer(ctx, items[i])
	}
	return items
}

func TestRelayerGet(t *testing.T) {
	keeper, ctx := keepertest.RelayersKeeper(t)
	items := createNRelayer(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRelayer(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestRelayerRemove(t *testing.T) {
	keeper, ctx := keepertest.RelayersKeeper(t)
	items := createNRelayer(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRelayer(ctx,
			item.Index,
		)
		_, found := keeper.GetRelayer(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestRelayerGetAll(t *testing.T) {
	keeper, ctx := keepertest.RelayersKeeper(t)
	items := createNRelayer(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllRelayer(ctx))
}
