package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/sudo/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) InflationBase(goCtx context.Context, req *types.QueryInflationBaseRequest) (*types.QueryInflationBaseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	inflationBase, found := k.GetInflationBase(ctx)
	if !found {
		return nil, types.ErrInflationBaseNotFound
	}

	return &types.QueryInflationBaseResponse{
		InflationBase: inflationBase,
	}, nil
}
