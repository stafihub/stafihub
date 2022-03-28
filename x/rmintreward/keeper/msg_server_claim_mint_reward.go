package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rmintreward/types"
)

func (k msgServer) ClaimMintReward(goCtx context.Context, msg *types.MsgClaimMintReward) (*types.MsgClaimMintRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	now := ctx.BlockHeight()

	claimInfo, found := k.Keeper.GetUserClaimInfo(ctx, user, msg.Denom, msg.Cycle, msg.MintIndex)
	if !found {
		return nil, types.ErrUserClaimInfoNotExist
	}
	act, found := k.Keeper.GetMintRewardAct(ctx, msg.Denom, msg.Cycle)
	if !found {
		return nil, types.ErrMintRewardActNotExist
	}
	finalBlock := claimInfo.MintBLock + act.LockedBlocks

	for _, tokenClaimInfo := range claimInfo.TokenClaimInfos {
		leftClaimAmount := tokenClaimInfo.TotalRewardAmount.Sub(tokenClaimInfo.TotalClaimedAmount)
		if leftClaimAmount.LTE(sdk.ZeroInt()) {
			continue
		}

		shouldClaimAmount := leftClaimAmount
		if now < finalBlock {
			duBlocks := now - claimInfo.LatestClaimedBlock
			lockedDuBlocks := finalBlock - claimInfo.LatestClaimedBlock
			shouldClaimAmount = leftClaimAmount.Mul(sdk.NewInt(int64(duBlocks))).Quo(sdk.NewInt(int64(lockedDuBlocks)))
		}
		if shouldClaimAmount.GT(sdk.ZeroInt()) {
			if err := k.bankKeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user, sdk.NewCoins(sdk.NewCoin(tokenClaimInfo.Denom, shouldClaimAmount))); err != nil {
				return nil, err
			}
			tokenClaimInfo.TotalClaimedAmount = tokenClaimInfo.TotalClaimedAmount.Add(shouldClaimAmount)
		}
	}
	claimInfo.LatestClaimedBlock = now

	k.Keeper.SetUserClaimInfo(ctx, user, msg.Denom, msg.Cycle, msg.MintIndex, claimInfo)

	return &types.MsgClaimMintRewardResponse{}, nil
}
