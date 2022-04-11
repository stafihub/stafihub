package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rmintreward/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) UpdateMintRewardAct(goCtx context.Context, msg *types.MsgUpdateMintRewardAct) (*types.MsgUpdateMintRewardActResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	actOnchain, found := k.Keeper.GetMintRewardAct(ctx, msg.Denom, msg.Cycle)
	if !found {
		return nil, types.ErrActNotExist
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
	if msg.Act.LockedBlocks <= 0 {
		return nil, types.ErrActLockedBlocksTooSmall
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

		willUseLeftAmount := sdk.ZeroInt()
		for _, onchainRewardInfo := range actOnchain.TokenRewardInfos {
			if onchainRewardInfo.Denom == rewardInfo.Denom {
				willUseLeftAmount = onchainRewardInfo.LeftAmount

				if rewardInfo.TotalRewardAmount.GT(onchainRewardInfo.TotalRewardAmount) {
					willUseLeftAmount = onchainRewardInfo.LeftAmount.Add(rewardInfo.TotalRewardAmount).Sub(onchainRewardInfo.TotalRewardAmount)
				} else {
					willUseLeftAmount = onchainRewardInfo.LeftAmount.Sub(onchainRewardInfo.TotalRewardAmount.Sub(rewardInfo.TotalRewardAmount))
					if willUseLeftAmount.LT(sdk.ZeroInt()) {
						willUseLeftAmount = sdk.ZeroInt()
					}
				}
			}
		}

		willUseTokenRewardInfos = append(willUseTokenRewardInfos, &types.TokenRewardInfo{
			Denom:             rewardInfo.Denom,
			RewardRate:        rewardInfo.RewardRate,
			TotalRewardAmount: rewardInfo.TotalRewardAmount,
			LeftAmount:        willUseLeftAmount,
			UserLimit:         rewardInfo.UserLimit,
		})
	}

	willUseAct := types.MintRewardAct{
		Begin:                  msg.Act.Begin,
		End:                    msg.Act.End,
		LockedBlocks:           msg.Act.LockedBlocks,
		TotalRTokenAmount:      actOnchain.TotalRTokenAmount,
		TotalNativeTokenAmount: actOnchain.TotalNativeTokenAmount,
		TokenRewardInfos:       willUseTokenRewardInfos,
	}

	k.Keeper.SetMintRewardAct(ctx, msg.Denom, msg.Cycle, &willUseAct)
	return &types.MsgUpdateMintRewardActResponse{}, nil
}
