package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/sudo/types"
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

func (k msgServer) UpdateAdmin(goCtx context.Context,  msg *types.MsgUpdateAdmin) (*types.MsgUpdateAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	newAdmin, _ := sdk.AccAddressFromBech32(msg.Address)

	isAdmin := k.IsAdmin(ctx, creator)
	if !isAdmin {
		return nil, types.ErrCreatorNotAdmin
	}

	k.SetAdmin(ctx, newAdmin)

	return &types.MsgUpdateAdminResponse{}, nil
}

func (k msgServer) AddDenom(goCtx context.Context,  msg *types.MsgAddDenom) (*types.MsgAddDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Keeper.AddDenom(ctx, msg.Denom)
	return &types.MsgAddDenomResponse{}, nil
}