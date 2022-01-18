package relayers_test

import (
	"testing"

	keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
	"github.com/stafiprotocol/stafihub/x/relayers"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Relayers: []types.Relayer{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		Thresholds: []types.Threshold{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		ProposalList: []types.Proposal{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RelayersKeeper(t)
	relayers.InitGenesis(ctx, *k, genesisState)
	got := relayers.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.Relayers, len(genesisState.Relayers))
	require.Subset(t, genesisState.Relayers, got.Relayers)
	require.Len(t, got.Thresholds, len(genesisState.Thresholds))
	require.Subset(t, genesisState.Thresholds, got.Thresholds)
	require.Len(t, got.ProposalList, len(genesisState.ProposalList))
	require.Subset(t, genesisState.ProposalList, got.ProposalList)
	// this line is used by starport scaffolding # genesis/test/assert
}
