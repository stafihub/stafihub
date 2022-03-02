package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Threshold(goCtx context.Context, req *types.QueryThresholdRequest) (*types.QueryThresholdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	threshold, found := k.GetThreshold(ctx, uint8(req.ChainId))
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryThresholdResponse{
		Threshold: fmt.Sprintf("%d", threshold),
	}, nil
}
