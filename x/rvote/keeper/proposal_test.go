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

func createNProposal(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Proposal {
	items := make([]types.Proposal, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetProposal(ctx, items[i])
	}
	return items
}

func TestProposalGet(t *testing.T) {
	keeper, ctx := keepertest.RelayersKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetProposal(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestProposalRemove(t *testing.T) {
	keeper, ctx := keepertest.RelayersKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProposal(ctx,
			item.Index,
		)
		_, found := keeper.GetProposal(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestProposalGetAll(t *testing.T) {
	keeper, ctx := keepertest.RelayersKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllProposal(ctx))
}
