package rdex

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rdex/keeper"
	"github.com/stafihub/stafihub/x/rdex/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	for _, poolCreator := range genState.PoolCreatorList {
		poolCreatorAddr, err := sdk.AccAddressFromBech32(poolCreator)
		if err != nil {
			panic(err)
		}
		k.AddPoolCreator(ctx, poolCreatorAddr)
	}

	for _, provider := range genState.ProviderList {
		providerAddr, err := sdk.AccAddressFromBech32(provider)
		if err != nil {
			panic(err)
		}
		k.AddProvider(ctx, providerAddr)
	}

	k.SetProviderSwitch(ctx, genState.ProviderSwitch)
	for _, swapPool := range genState.SwapPoolList {
		lpDenom := types.GetLpTokenDenom(swapPool.Index)
		k.SetSwapPool(ctx, lpDenom, swapPool)
	}
	if len(genState.SwapPoolList) > 0 {
		k.SetSwapPoolIndex(ctx, uint32(len(genState.SwapPoolList)-1))
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.PoolCreatorList = k.GetPoolCreatorList(ctx)
	genesis.ProviderList = k.GetProviderList(ctx)
	genesis.ProviderSwitch = k.GetProviderSwitch(ctx)
	genesis.SwapPoolList = k.GetSwapPoolList(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
