package keeper

import (
	"context"

    "github.com/stafiprotocol/stafihub/x/rate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Rate(goCtx context.Context,  req *types.QueryRateRequest) (*types.QueryRateResponse, error) {
	if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

	if req.Denom == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid denom")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
    ratio := k.GetRate(ctx, req.Denom)

	return &types.QueryRateResponse{Ratio: &ratio}, nil
}
