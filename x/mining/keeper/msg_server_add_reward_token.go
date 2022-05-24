package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) AddRewardToken(goCtx context.Context, msg *types.MsgAddRewardToken) (*types.MsgAddRewardTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	rewardToken := types.RewardToken{
		RewardTokenDenom:     msg.Denom,
		MinTotalRewardAmount: msg.MinTotalRewardAmount,
	}

	k.Keeper.AddRewardToken(ctx, &rewardToken)

	return &types.MsgAddRewardTokenResponse{}, nil
}
