package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) Stake(goCtx context.Context, msg *types.MsgStake) (*types.MsgStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakeToken.Denom)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}

	stakeItem, found := k.Keeper.GetStakeItem(ctx, msg.StakeItemIndex)
	if !found {
		return nil, types.ErrStakeItemNotExist
	}

	curBlockTime := uint64(ctx.BlockTime().Unix())
	// update pools
	updateStakePool(stakePool, curBlockTime)

	stakePool.TotalStakedAmount = stakePool.TotalStakedAmount.Add(msg.StakeToken.Amount)
	userStakePower := stakeItem.PowerRewardRate.MulInt(msg.StakeToken.Amount).TruncateInt()
	stakePool.TotalStakedPower = stakePool.TotalStakedPower.Add(userStakePower)

	willUseIndex := k.Keeper.GetUserStakeRecordNextIndex(ctx, msg.Creator, msg.StakeToken.Denom)

	rewardInfos := make([]*types.RewardInfo, 0)
	for _, rewardPool := range stakePool.RewardPools {
		rewardInfos = append(rewardInfos, &types.RewardInfo{
			RewardPoolIndex:  rewardPool.Index,
			RewardTokenDenom: rewardPool.RewardTokenDenom,
			RewardDebt:       userStakePower.Mul(rewardPool.RewardPerPower).Quo(types.RewardFactor),
		})
	}

	userStakeRecord := types.UserStakeRecord{
		UserAddress:     msg.Creator,
		StakeTokenDenom: msg.StakeToken.Denom,
		Index:           willUseIndex,
		StakedAmount:    msg.StakeToken.Amount,
		StakedPower:     userStakePower,
		StartTimestamp:  curBlockTime,
		EndTimestamp:    curBlockTime + stakeItem.LockSecond,
		RewardInfos:     rewardInfos,
		StakeItemIndex:  msg.StakeItemIndex,
	}

	if err := k.Keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddr, types.ModuleName, sdk.NewCoins(msg.StakeToken)); err != nil {
		return nil, err
	}

	k.Keeper.SetUserStakeRecordIndex(ctx, msg.Creator, msg.StakeToken.Denom, willUseIndex)
	k.Keeper.SetUserStakeRecord(ctx, &userStakeRecord)
	k.Keeper.SetStakePool(ctx, stakePool)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeStake,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyStakeTokenDenom, msg.StakeToken.Denom),
			sdk.NewAttribute(types.AttributeKeyStakeRecordIndex, fmt.Sprintf("%d", willUseIndex)),
			sdk.NewAttribute(types.AttributeKeyStakeTokenAmount, msg.StakeToken.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyStakePower, userStakePower.String()),
			sdk.NewAttribute(types.AttributeKeyStartTimestamp, fmt.Sprintf("%d", userStakeRecord.StartTimestamp)),
			sdk.NewAttribute(types.AttributeKeyEndTimestamp, fmt.Sprintf("%d", userStakeRecord.EndTimestamp)),
			sdk.NewAttribute(types.AttributeKeyStakeItemIndex, fmt.Sprintf("%d", msg.StakeItemIndex)),
		),
	)

	return &types.MsgStakeResponse{}, nil
}
