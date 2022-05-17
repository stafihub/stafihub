package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) AddStakeItem(goCtx context.Context, msg *types.MsgAddStakeItem) (*types.MsgAddStakeItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	willUseIndex := k.GetStakeItemNextIndex(ctx)

	stakeItem := types.StakeItem{
		Enable:          msg.Enable,
		Index:           willUseIndex,
		LockSecond:      msg.LockSecond,
		PowerRewardRate: msg.PowerRewardRate,
	}

	k.Keeper.SetStakeItemIndex(ctx, willUseIndex)
	k.Keeper.SetStakeItem(ctx, &stakeItem)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddStakeItem,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyStakeItemIndex, fmt.Sprintf("%d", willUseIndex)),
			sdk.NewAttribute(types.AttributeKeyLockSecond, fmt.Sprintf("%d", msg.LockSecond)),
			sdk.NewAttribute(types.AttributeKeyPowerRewardRate, msg.PowerRewardRate.String()),
		),
	)
	return &types.MsgAddStakeItemResponse{}, nil
}
