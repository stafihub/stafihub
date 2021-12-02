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
		RelayerList: []types.Relayer{
	{
		Index: "0",
},
	{
		Index: "1",
},
},
ThresholdList: []types.Threshold{
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

	require.Len(t, got.RelayerList, len(genesisState.RelayerList))
require.Subset(t, genesisState.RelayerList, got.RelayerList)
require.Len(t, got.ThresholdList, len(genesisState.ThresholdList))
require.Subset(t, genesisState.ThresholdList, got.ThresholdList)
// this line is used by starport scaffolding # genesis/test/assert
}
