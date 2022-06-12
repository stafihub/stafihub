package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LatestVotedCycle(goCtx context.Context, req *types.QueryLatestVotedCycleRequest) (*types.QueryLatestVotedCycleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryLatestVotedCycleResponse{
		LatestVotedCycle: k.GetLatestVotedCycle(ctx, req.Denom),
	}, nil
}
