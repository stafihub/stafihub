package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetRelayFee(goCtx context.Context, msg *types.MsgSetRelayFee) (*types.MsgSetRelayFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetRelayFee(ctx, uint8(msg.ChainId), msg.Value)

	return &types.MsgSetRelayFeeResponse{}, nil
}
