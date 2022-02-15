package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetInflationBase(goCtx context.Context, msg *types.MsgSetInflationBase) (*types.MsgSetInflationBaseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	isAdmin := k.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, types.ErrCreatorNotAdmin
	}

	k.Keeper.SetInflationBase(ctx, msg.InflationBase)

	return &types.MsgSetInflationBaseResponse{}, nil
}
