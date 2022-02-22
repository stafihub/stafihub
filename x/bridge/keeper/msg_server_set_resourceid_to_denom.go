package keeper

import (
	"context"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetResourceidToDenom(goCtx context.Context, msg *types.MsgSetResourceidToDenom) (*types.MsgSetResourceidToDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}
	resourceIdBts, err := hex.DecodeString(msg.ResourceId)
	if err != nil {
		return nil, err
	}
	var resourceId [32]byte
	copy(resourceId[:], resourceIdBts)

	k.Keeper.SetResourceIdToDenom(ctx, resourceId, msg.Denom)
	return &types.MsgSetResourceidToDenomResponse{}, nil
}
