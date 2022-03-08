package keeper

import (
	"context"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetResourceidType(goCtx context.Context, msg *types.MsgSetResourceidType) (*types.MsgSetResourceidTypeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}
	resourceIdSlice, err := hex.DecodeString(msg.ResourceId)
	if err != nil {
		return nil, types.ErrResourceIdFormatNotRight
	}
	var resourceId [32]byte
	copy(resourceId[:], resourceIdSlice)

	switch msg.IdType {
	case "0":
		k.Keeper.SetResourceIdType(ctx, resourceId, types.ResourceIdTypeForeign)
	case "1":
		k.Keeper.SetResourceIdType(ctx, resourceId, types.ResourceIdTypeNative)
	default:
		return nil, types.ErrUnKnownResourceIdType
	}

	return &types.MsgSetResourceidTypeResponse{}, nil
}
