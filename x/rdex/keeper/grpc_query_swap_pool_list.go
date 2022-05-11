package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rdex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SwapPoolList(goCtx context.Context, req *types.QuerySwapPoolListRequest) (*types.QuerySwapPoolListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QuerySwapPoolListResponse{
		SwapPoolList: k.GetSwapPoolList(ctx),
	}, nil
}
