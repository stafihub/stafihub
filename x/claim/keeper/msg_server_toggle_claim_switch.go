package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/claim/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) ToggleClaimSwitch(goCtx context.Context, msg *types.MsgToggleClaimSwitch) (*types.MsgToggleClaimSwitchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}
	k.Keeper.ToggleClaimSwitch(ctx, msg.Round)

	return &types.MsgToggleClaimSwitchResponse{}, nil
}
