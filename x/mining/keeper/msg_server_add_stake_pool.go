package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) AddStakePool(goCtx context.Context, msg *types.MsgAddStakePool) (*types.MsgAddStakePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	_, found := k.GetStakePool(ctx, msg.StakeTokenDenom)
	if found {
		return nil, types.ErrStakePoolAlreadyExist
	}

	stakePool := types.StakePool{
		StakeTokenDenom:   msg.StakeTokenDenom,
		MaxRewardPools:    msg.MaxRewardPools,
		RewardPools:       []*types.RewardPool{},
		TotalStakedAmount: sdk.ZeroInt(),
		TotalStakedPower:  sdk.ZeroInt(),
	}

	k.SetStakePool(ctx, &stakePool)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddStakePool,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyStakeTokenDenom, msg.StakeTokenDenom),
			sdk.NewAttribute(types.AttributeKeyMaxRewardPools, fmt.Sprintf("%d", msg.MaxRewardPools)),
		),
	)

	return &types.MsgAddStakePoolResponse{}, nil
}
