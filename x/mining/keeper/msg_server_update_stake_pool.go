package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) UpdateStakePool(goCtx context.Context, msg *types.MsgUpdateStakePool) (*types.MsgUpdateStakePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	stakePool, found := k.GetStakePool(ctx, msg.StakeTokenDenom)
	if !found {
		return nil, types.ErrStakePoolNotExist
	}

	stakePool.MaxRewardPools = msg.MaxRewardPools
	stakePool.MinTotalRewardAmount = msg.MinTotalRewardAmount

	k.Keeper.SetStakePool(ctx, stakePool)

	return &types.MsgUpdateStakePoolResponse{}, nil
}
