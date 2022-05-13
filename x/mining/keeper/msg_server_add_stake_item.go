package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) AddStakeItem(goCtx context.Context, msg *types.MsgAddStakeItem) (*types.MsgAddStakeItemResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	stakeItem := types.StakeItem{
		Index:           msg.Index,
		LockSecond:      msg.LockSecond,
		PowerRewardRate: msg.PowerRewardRate,
	}
	k.Keeper.SetStakeItem(ctx, &stakeItem)

	return &types.MsgAddStakeItemResponse{}, nil
}
