package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StakePoolInfo(goCtx context.Context, req *types.QueryStakePoolInfoRequest) (*types.QueryStakePoolInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	stakePool, found := k.GetStakePool(ctx, req.StakePoolIndex)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryStakePoolInfoResponse{
		StakePool: stakePool,
	}, nil
}
