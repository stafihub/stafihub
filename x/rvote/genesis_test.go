package rvote_test

import (
	"testing"

	keepertest "github.com/stafihub/stafihub/testutil/keeper"
	"github.com/stafihub/stafihub/x/rvote"
	"github.com/stafihub/stafihub/x/rvote/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RvoteKeeper(t)
	rvote.InitGenesis(ctx, *k, genesisState)
	got := rvote.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
