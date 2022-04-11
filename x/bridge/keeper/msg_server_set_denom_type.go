package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetDenomType(goCtx context.Context, msg *types.MsgSetDenomType) (*types.MsgSetDenomTypeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	resourceId, found := k.Keeper.GetResourceIdByDenom(ctx, msg.Denom)
	if !found {
		return nil, types.ErrResourceIdNotFound
	}

	switch msg.IdType {
	case "0":
		k.Keeper.SetResourceIdType(ctx, resourceId, types.ResourceIdTypeForeign)
	case "1":
		k.Keeper.SetResourceIdType(ctx, resourceId, types.ResourceIdTypeNative)
	default:
		return nil, types.ErrUnKnownResourceIdType
	}

	return &types.MsgSetDenomTypeResponse{}, nil
}
