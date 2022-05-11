package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rdex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SwapPoolInfo(goCtx context.Context, req *types.QuerySwapPoolInfoRequest) (*types.QuerySwapPoolInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	swapPool, found := k.GetSwapPool(ctx, req.LpDenom)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QuerySwapPoolInfoResponse{
		SwapPool: swapPool,
	}, nil
}
