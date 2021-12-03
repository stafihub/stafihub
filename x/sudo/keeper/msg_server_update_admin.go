package keeper

import (
	"context"
	"fmt"

	"github.com/stafiprotocol/stafihub/x/sudo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) UpdateAdmin(goCtx context.Context,  msg *types.MsgUpdateAdmin) (*types.MsgUpdateAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	newAdmin, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	isAdmin := k.IsAdmin(ctx, creator)
    if !isAdmin {
		return nil, fmt.Errorf("creator is not admin")
	}

	k.SetAdmin(ctx, newAdmin)

	return &types.MsgUpdateAdminResponse{}, nil
}
