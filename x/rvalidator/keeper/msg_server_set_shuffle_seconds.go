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

	shuffleSeconds := k.Keeper.GetShuffleSeconds(ctx, msg.Denom)

	k.Keeper.SetShuffleSeconds(ctx, &types.ShuffleSeconds{
		Denom:   msg.Denom,
		Version: shuffleSeconds.Version,
		Seconds: shuffleSeconds.Seconds,
	})

	return &types.MsgSetShuffleSecondsResponse{}, nil
}
