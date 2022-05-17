package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) RmRewarder(goCtx context.Context, msg *types.MsgRmRewarder) (*types.MsgRmRewarderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	userAddr, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		return nil, err
	}

	if !k.Keeper.HasRewarder(ctx, userAddr) {
		return nil, types.ErrUserNotRewarder
	}

	k.Keeper.RemoveRewarder(ctx, userAddr)
	return &types.MsgRmRewarderResponse{}, nil
}
