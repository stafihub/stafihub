package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) ToggleEmergencySwitch(goCtx context.Context, msg *types.MsgToggleEmergencySwitch) (*types.MsgToggleEmergencySwitchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	stakePool, found := k.Keeper.GetStakePool(ctx, msg.GetStakePoolIndex())
	if !found {
		return nil, types.ErrStakePoolNotExist
	}

	stakePool.EmergencySwitch = !stakePool.EmergencySwitch

	k.Keeper.SetStakePool(ctx, stakePool)

	return &types.MsgToggleEmergencySwitchResponse{}, nil
}
