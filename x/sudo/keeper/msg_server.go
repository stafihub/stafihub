package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/sudo/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) UpdateAdmin(goCtx context.Context, msg *types.MsgUpdateAdmin) (*types.MsgUpdateAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	lastAdmin := k.GetAdmin(ctx)
	newAdmin, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	if lastAdmin.Equals(newAdmin) {
		return nil, types.ErrLastAdminEqualNewAdmin
	}

	isAdmin := k.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, types.ErrCreatorNotAdmin
	}

	k.SetAdmin(ctx, newAdmin)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAdminUpdated,
			sdk.NewAttribute(types.AttributeKeyLastAdmin, lastAdmin.String()),
			sdk.NewAttribute(types.AttributeKeyCurrentAdmin, msg.Address),
		),
	)

	return &types.MsgUpdateAdminResponse{}, nil
}
