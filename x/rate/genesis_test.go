package rate_test

import (
	"testing"

	keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
	"github.com/stafiprotocol/stafihub/x/rate"
	"github.com/stafiprotocol/stafihub/x/rate/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		ExchangeRateList: []types.ExchangeRate{
	{
		Index: "0",
},
	{
		Index: "1",
},
},
EraExchangeRateList: []types.EraExchangeRate{
	{
		Index: "0",
},
	{
		Index: "1",
},
},
// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RateKeeper(t)
	rate.InitGenesis(ctx, *k, genesisState)
	got := rate.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.ExchangeRateList, len(genesisState.ExchangeRateList))
require.Subset(t, genesisState.ExchangeRateList, got.ExchangeRateList)
require.Len(t, got.EraExchangeRateList, len(genesisState.EraExchangeRateList))
require.Subset(t, genesisState.EraExchangeRateList, got.EraExchangeRateList)
// this line is used by starport scaffolding # genesis/test/assert
}
