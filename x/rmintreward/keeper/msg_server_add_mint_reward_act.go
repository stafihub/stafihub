package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rmintreward/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) AddMintRewardAct(goCtx context.Context, msg *types.MsgAddMintRewardAct) (*types.MsgAddMintRewardActResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	if msg.Act.End <= ctx.BlockHeight() {
		return nil, types.ErrActEndBlockLessThanCurrentBlock
	}

	if msg.Act.Begin <= 0 {
		return nil, types.ErrActBeginBlockTooSmall
	}
	if msg.Act.Begin >= msg.Act.End {
		return nil, types.ErrActEndBlockTooSamll
	}
	
	willUseTokenRewardInfos := make([]*types.TokenRewardInfo, 0)
	for _, rewardInfo := range msg.Act.TokenRewardInfos {
		if rewardInfo.TotalRewardAmount.LTE(sdk.ZeroInt()) {
			return nil, types.ErrActTotalRewardTooSmall
		}
		if rewardInfo.TotalRewardAmount.LTE(rewardInfo.UserLimit) {
			return nil, types.ErrActTotalRewardLessThanUserLimit
		}
		if rewardInfo.RewardRate.LTE(sdk.ZeroDec()) {
			return nil, types.ErrActRewardRateTooSmall
		}

		willUseTokenRewardInfos = append(willUseTokenRewardInfos, &types.TokenRewardInfo{
			Denom:             rewardInfo.Denom,
			RewardRate:        rewardInfo.RewardRate,
			TotalRewardAmount: rewardInfo.TotalRewardAmount,
			LeftAmount:        rewardInfo.TotalRewardAmount,
			UserLimit:         rewardInfo.UserLimit,
		})

	}
	if msg.Act.LockedBlocks <= 0 {
		return nil, types.ErrActLockedBlocksTooSmall
	}

	willUseCycle := uint64(0)
	latestCycle, found := k.Keeper.GetActLatestCycle(ctx, msg.Denom)
	if found {
		latestMintReward, found := k.Keeper.GetMintRewardAct(ctx, msg.Denom, latestCycle)
		if !found {
			return nil, types.ErrLatestMintRewardActNotExist
		}

		if msg.Act.Begin <= latestMintReward.End {
			return nil, types.ErrActEndBlockLessThanCurrentBlock
		}

		willUseCycle = latestCycle + 1
	}

	willUseAct := types.MintRewardAct{
		Begin:                  msg.Act.Begin,
		End:                    msg.Act.End,
		LockedBlocks:           msg.Act.LockedBlocks,
		TotalRTokenAmount:      sdk.ZeroInt(),
		TotalNativeTokenAmount: sdk.ZeroInt(),
		TokenRewardInfos:       willUseTokenRewardInfos,
	}

	k.Keeper.SetMintRewardAct(ctx, msg.Denom, willUseCycle, &willUseAct)
	k.Keeper.SetActLatestCycle(ctx, msg.Denom, willUseCycle)

	return &types.MsgAddMintRewardActResponse{}, nil
}
