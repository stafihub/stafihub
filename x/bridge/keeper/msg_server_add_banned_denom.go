package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) AddBannedDenom(goCtx context.Context, msg *types.MsgAddBannedDenom) (*types.MsgAddBannedDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.AddBannedDenom(ctx, uint8(msg.ChainId), msg.Denom)

	return &types.MsgAddBannedDenomResponse{}, nil
}
