package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/stafiprotocol/stafihub/x/rate/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EraRate(goCtx context.Context,  req *types.QueryEraRateRequest) (*types.QueryEraRateResponse, error) {
	if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

	if req.Denom == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid denom")
	}

	era, err := strconv.ParseUint(req.Era, 10, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("req.era %s not a valid uint, please input a valid era", req.Era))
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	ratio := k.GetEraRate(ctx, req.Denom, uint32(era))
	return &types.QueryEraRateResponse{Ratio: &ratio}, nil
}
