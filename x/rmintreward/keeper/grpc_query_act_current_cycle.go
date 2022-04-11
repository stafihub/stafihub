package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rmintreward/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ActCurrentCycle(goCtx context.Context, req *types.QueryActCurrentCycleRequest) (*types.QueryActCurrentCycleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	currentCycle, found := k.GetActCurrentCycle(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryActCurrentCycleResponse{
		CurrentCycle: currentCycle,
	}, nil
}
