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
	finalBlock := claimInfo.MintBlock + act.LockedBlocks

	shouldClaimCoins := sdk.Coins{}
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
			shouldClaimCoins = shouldClaimCoins.Add(sdk.NewCoin(tokenClaimInfo.Denom, shouldClaimAmount))
			tokenClaimInfo.TotalClaimedAmount = tokenClaimInfo.TotalClaimedAmount.Add(shouldClaimAmount)
		}
	}
	if shouldClaimCoins.Empty() {
		return nil, types.ErrNoRewardToClaim
	}

	claimInfo.LatestClaimedBlock = now
	if err := k.bankKeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user, shouldClaimCoins); err != nil {
		return nil, err
	}
	k.Keeper.SetUserClaimInfo(ctx, user, msg.Denom, msg.Cycle, msg.MintIndex, claimInfo)

	return &types.MsgClaimMintRewardResponse{}, nil
}
