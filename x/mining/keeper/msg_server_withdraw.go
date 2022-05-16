package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	recipientAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	userStakeRecord, found := k.Keeper.GetUserStakeRecord(ctx, msg.Creator, msg.StakeToken.Denom, msg.StakeRecordIndex)
	if !found {
		return nil, types.ErrUserStakeRecordNotExist
	}
	stakePool, found := k.Keeper.GetStakePool(ctx, msg.StakeToken.Denom)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}
	curBlockTime := uint64(ctx.BlockTime().Unix())

	if userStakeRecord.EndTimestamp > curBlockTime {
		return nil, types.ErrStakeTokenStillLocked
	}

	updateStakePool(stakePool, curBlockTime)
	willClaimCoins := calRewardTokens(stakePool, userStakeRecord)

	willRmPower := msg.StakeToken.Amount.Mul(userStakeRecord.StakedPower).Quo(userStakeRecord.StakedAmount)
	stakePool.TotalStakedAmount = stakePool.TotalStakedAmount.Sub(msg.StakeToken.Amount)
	stakePool.TotalStakedPower = stakePool.TotalStakedPower.Sub(willRmPower)

	userStakeRecord.StakedAmount = userStakeRecord.StakedAmount.Sub(msg.StakeToken.Amount)
	userStakeRecord.StakedPower = userStakeRecord.StakedPower.Sub(willRmPower)

	willClaimCoins = willClaimCoins.Add(msg.StakeToken)
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipientAddr, willClaimCoins); err != nil {
		return nil, err
	}

	k.SetStakePool(ctx, stakePool)
	k.SetUserStakeRecord(ctx, userStakeRecord)

	return &types.MsgWithdrawResponse{}, nil
}
