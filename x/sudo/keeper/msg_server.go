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

	lastAdmin := k.GetAdmin(ctx).String()
	newAdmin, _ := sdk.AccAddressFromBech32(msg.Address)
	isAdmin := k.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, types.ErrCreatorNotAdmin
	}

	k.SetAdmin(ctx, newAdmin)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAdminUpdated,
			sdk.NewAttribute(types.AttributeKeyLastAdmin, lastAdmin),
			sdk.NewAttribute(types.AttributeKeyCurrentAdmin, msg.Address),
		),
	)

	return &types.MsgUpdateAdminResponse{}, nil
}

func (k msgServer) AddDenom(goCtx context.Context, msg *types.MsgAddDenom) (*types.MsgAddDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, types.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Metadata.Base)
	if ok {
		return nil, types.ErrDenomAlreadyExist
	}

	k.bankKeeper.SetDenomMetaData(ctx, msg.Metadata)
	return &types.MsgAddDenomResponse{}, nil
}
