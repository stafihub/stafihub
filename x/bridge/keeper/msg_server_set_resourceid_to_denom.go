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
	if err != nil || len(resourceIdBts) != 32 {
		return nil, types.ErrResourceIdFormatNotRight
	}

	if msg.DenomType != types.External && msg.DenomType != types.Native {
		return nil, types.ErrDenomTypeUnmatch
	}

	rs := types.ResourceIdToDenom{
		ResourceId: msg.ResourceId,
		Denom:      msg.Denom,
		DenomType:  msg.DenomType,
	}

	k.Keeper.SetResourceIdToDenom(ctx, &rs)
	return &types.MsgSetResourceidToDenomResponse{}, nil
}
