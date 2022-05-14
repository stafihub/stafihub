package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

func (k msgServer) AddStakePool(goCtx context.Context, msg *types.MsgAddStakePool) (*types.MsgAddStakePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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

	return &types.MsgAddStakePoolResponse{}, nil
}
