package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/rate/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) EraExchangeRate(c context.Context, req *types.QueryGetEraExchangeRateRequest) (*types.QueryGetEraExchangeRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEraExchangeRate(ctx, req.Denom, req.Era)
	if !found {
	    return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetEraExchangeRateResponse{EraExchangeRate: val}, nil
}

func (k Keeper) EraExchangeRateByDenom(goCtx context.Context,  req *types.QueryEraExchangeRateByDenomRequest) (*types.QueryEraExchangeRateByDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	eers := k.GetEraExchangeRateByDenom(ctx, req.Denom)

	return &types.QueryEraExchangeRateByDenomResponse{EraExchangeRates: eers}, nil
}