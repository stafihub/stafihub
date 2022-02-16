package stakextra_test

import (
	"testing"

	keepertest "github.com/stafihub/stafihub/testutil/keeper"
	"github.com/stafihub/stafihub/testutil/nullify"
	"github.com/stafihub/stafihub/x/stakextra"
	"github.com/stafihub/stafihub/x/stakextra/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StakextraKeeper(t)
	stakextra.InitGenesis(ctx, *k, genesisState)
	got := stakextra.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
