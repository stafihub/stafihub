package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rstaking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) InflationBase(goCtx context.Context, req *types.QueryInflationBaseRequest) (*types.QueryInflationBaseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	inflationBase := k.GetInflationBase(ctx)

	return &types.QueryInflationBaseResponse{
		InflationBase: inflationBase,
	}, nil
}
