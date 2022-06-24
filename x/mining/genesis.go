package mining

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/keeper"
	"github.com/stafihub/stafihub/x/mining/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.SetMaxRewardPoolNumber(ctx, genState.MaxRewardPoolNumber)
	k.SetMaxStakeItemNumber(ctx, genState.MaxStakeItemNumber)
	for _, miningProvider := range genState.MiningProviderList {
		miningProviderAddr, err := sdk.AccAddressFromBech32(miningProvider)
		if err != nil {
			panic(err)
		}
		k.AddMiningProvider(ctx, miningProviderAddr)
	}
	k.SetMiningProviderSwitch(ctx, genState.MiningProviderSwitch)

	for _, rewardToken := range genState.RewardTokenList {
		k.AddRewardToken(ctx, rewardToken)
	}

	k.SetStakeItemLimit(ctx, genState.StakeItemLimit)

	for _, stakeItem := range genState.StakeItemList {
		nextIndex := k.GetStakeItemNextIndex(ctx, stakeItem.StakePoolIndex)
		k.SetStakeItem(ctx, stakeItem)
		k.SetStakeItemIndex(ctx, stakeItem.StakePoolIndex, nextIndex)
	}

	for _, stakePool := range genState.StakePoolList {
		nextIndex := k.GetStakePoolNextIndex(ctx)
		k.SetStakePool(ctx, stakePool)
		k.SetStakePoolIndex(ctx, nextIndex)
	}

	for _, stakeToken := range genState.StakeTokenList {
		k.AddStakeToken(ctx, stakeToken)
	}
	for _, userStakeRecord := range genState.UserStakeRecordList {
		nextIndex := k.GetUserStakeRecordNextIndex(ctx, userStakeRecord.UserAddress, userStakeRecord.StakePoolIndex)
		k.SetUserStakeRecord(ctx, userStakeRecord)
		k.SetUserStakeRecordIndex(ctx, userStakeRecord.UserAddress, userStakeRecord.StakePoolIndex, nextIndex)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.MaxRewardPoolNumber = k.GetMaxRewardPoolNumber(ctx)
	genesis.MaxStakeItemNumber = k.GetMaxStakeItemNumber(ctx)
	genesis.MiningProviderList = k.GetMiningProviderList(ctx)
	genesis.MiningProviderSwitch = k.GetMiningProviderSwitch(ctx)
	genesis.RewardTokenList = k.GetRewardTokenList(ctx)
	genesis.StakeItemLimit = k.GetStakeItemLimit(ctx)
	genesis.StakeItemList = k.GetAllStakeItemList(ctx)
	genesis.StakePoolList = k.GetStakePoolList(ctx)
	genesis.StakeTokenList = k.GetStakeTokenList(ctx)
	genesis.UserStakeRecordList = k.GetAllUserStakeRecordList(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
