package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetThreshold(goCtx context.Context, msg *types.MsgSetThreshold) (*types.MsgSetThresholdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetThreshold(ctx, uint8(msg.ChainId), uint8(msg.Threshold))

	return &types.MsgSetThresholdResponse{}, nil
}
