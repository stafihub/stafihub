package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LatestDealedCycle(goCtx context.Context, req *types.QueryLatestDealedCycleRequest) (*types.QueryLatestDealedCycleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryLatestDealedCycleResponse{
		LatestDealedCycle: k.GetLatestDealedCycle(ctx, req.Denom, req.PoolAddress),
	}, nil
}
