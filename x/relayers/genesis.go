package relayers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/relayers/keeper"
	"github.com/stafihub/stafihub/x/relayers/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the relayer
	for _, rel := range genState.Relayers {
		for _, addr := range rel.Addrs {
			k.AddRelayer(ctx, rel.Arena, rel.Denom, addr)
		}
	}
	// Set all the threshold
	for _, elem := range genState.Thresholds {
		k.SetThreshold(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Relayers = k.GetAllRelayer(ctx)
	genesis.Thresholds = k.GetAllThreshold(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
