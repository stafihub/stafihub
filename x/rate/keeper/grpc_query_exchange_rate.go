package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/rate/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ExchangeRateAll(c context.Context, req *types.QueryAllExchangeRateRequest) (*types.QueryAllExchangeRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	exchangeRates := k.GetAllExchangeRate(ctx)

	return &types.QueryAllExchangeRateResponse{ExchangeRate: exchangeRates}, nil
}

func (k Keeper) ExchangeRate(c context.Context, req *types.QueryGetExchangeRateRequest) (*types.QueryGetExchangeRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetExchangeRate(ctx, req.Denom)
	if !found {
	    return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetExchangeRateResponse{ExchangeRate: val}, nil
}