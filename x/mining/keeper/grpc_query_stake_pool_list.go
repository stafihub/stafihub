package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StakePoolList(goCtx context.Context, req *types.QueryStakePoolListRequest) (*types.QueryStakePoolListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryStakePoolListResponse{
		StakePoolList: k.GetStakePoolList(ctx),
	}, nil
}
