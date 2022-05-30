package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetStakeItemLimit(goCtx context.Context, msg *types.MsgSetStakeItemLimit) (*types.MsgSetStakeItemLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	k.Keeper.SetStakeItemLimit(ctx, &types.StakeItemLimit{
		MaxPowerRewardRate: msg.MaxPowerRewardRate,
		MaxLockSecond:      msg.MaxLockSecond,
	})

	return &types.MsgSetStakeItemLimitResponse{}, nil
}
