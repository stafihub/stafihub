package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) AddStakePool(goCtx context.Context, msg *types.MsgAddStakePool) (*types.MsgAddStakePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	if k.GetMiningProviderSwitch(ctx) && !k.HasMiningProvider(ctx, user) {
		return nil, types.ErrUserNotAdminOrMiningProvider
	}

	rewardToken, found := k.Keeper.GetRewardToken(ctx, msg.RewardTokenDenom)
	if !found {
		return nil, types.ErrRewardTokenNotSupport
	}
	if msg.TotalRewardAmount.LT(rewardToken.MinTotalRewardAmount) {
		return nil, types.ErrTotalRewardAmountLessThanLimit
	}

	curBlockTime := uint64(ctx.BlockTime().Unix())

	willUseStakePoolIndex := k.Keeper.GetStakePoolNextIndex(ctx)

	willUseLastRewardTimestamp := msg.StartTimestamp
	if msg.StartTimestamp < curBlockTime {
		willUseLastRewardTimestamp = curBlockTime
	}

	rewardTokens := sdk.NewCoins(sdk.NewCoin(msg.RewardTokenDenom, msg.TotalRewardAmount))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, rewardTokens); err != nil {
		return nil, err
	}

	stakePool := types.StakePool{
		Index:           willUseStakePoolIndex,
		StakeTokenDenom: msg.StakeTokenDenom,
		RewardPools: []*types.RewardPool{
			{
				Index:               0,
				RewardTokenDenom:    msg.RewardTokenDenom,
				TotalRewardAmount:   msg.TotalRewardAmount,
				LeftRewardAmount:    msg.TotalRewardAmount,
				RewardPerSecond:     msg.RewardPerSecond,
				StartTimestamp:      msg.StartTimestamp,
				RewardPerPower:      sdk.ZeroInt(),
				LastRewardTimestamp: willUseLastRewardTimestamp,
			}},
		TotalStakedAmount: sdk.ZeroInt(),
		TotalStakedPower:  sdk.ZeroInt(),
	}

	k.SetStakePool(ctx, &stakePool)
	k.Keeper.SetRewardPoolIndex(ctx, willUseStakePoolIndex, 0)
	k.Keeper.SetStakePoolIndex(ctx, willUseStakePoolIndex)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddStakePool,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyStakeTokenDenom, msg.StakeTokenDenom),
			sdk.NewAttribute(types.AttributeKeyStakePoolIndex, fmt.Sprintf("%d", willUseStakePoolIndex)),
		),
	)

	return &types.MsgAddStakePoolResponse{}, nil
}
