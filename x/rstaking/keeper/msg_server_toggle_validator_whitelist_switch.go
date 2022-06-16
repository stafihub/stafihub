package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rstaking/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) ToggleValidatorWhitelistSwitch(goCtx context.Context, msg *types.MsgToggleValidatorWhitelistSwitch) (*types.MsgToggleValidatorWhitelistSwitchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	k.Keeper.ToggleValidatorWhitelistSwitch(ctx)

	return &types.MsgToggleValidatorWhitelistSwitchResponse{}, nil
}
