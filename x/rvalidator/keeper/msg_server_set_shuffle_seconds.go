package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetShuffleSeconds(goCtx context.Context, msg *types.MsgSetShuffleSeconds) (*types.MsgSetShuffleSecondsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetShuffleSeconds(ctx, msg.Denom, msg.Seconds)

	return &types.MsgSetShuffleSecondsResponse{}, nil
}
