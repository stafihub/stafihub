package exchangerate_test

import (
	"testing"

	keepertest "github.com/stafiprotocol/stafihub/testutil/keeper"
	"github.com/stafiprotocol/stafihub/x/exchangerate"
	"github.com/stafiprotocol/stafihub/x/exchangerate/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ExchangerateKeeper(t)
	exchangerate.InitGenesis(ctx, *k, genesisState)
	got := exchangerate.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
