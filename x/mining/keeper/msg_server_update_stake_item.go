package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) UpdateStakeItem(goCtx context.Context, msg *types.MsgUpdateStakeItem) (*types.MsgUpdateStakeItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	stakeItem, found := k.GetStakeItem(ctx, msg.StakePoolIndex, msg.Index)
	if !found {
		return nil, types.ErrStakeItemNotExist
	}

	stakeItem.LockSecond = msg.LockSecond
	stakeItem.PowerRewardRate = msg.PowerRewardRate
	stakeItem.Enable = msg.Enable

	k.SetStakeItem(ctx, stakeItem)

	return &types.MsgUpdateStakeItemResponse{}, nil
}
