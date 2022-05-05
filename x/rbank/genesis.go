package rbank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rbank/keeper"
	"github.com/stafihub/stafihub/x/rbank/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	for _, prefix := range genState.AddressPrefix {
		if len(prefix.Denom) == 0 || len(prefix.AccAddressPrefix) == 0 || len(prefix.ValAddressPrefix) == 0 {
			panic("not valid address prefix")
		}
		k.SetAccAddressPrefix(ctx, prefix.Denom, prefix.AccAddressPrefix)
		k.SetValAddressPrefix(ctx, prefix.Denom, prefix.ValAddressPrefix)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	list := k.GetAddressPrefixList(ctx)
	genesis.AddressPrefix = list

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
