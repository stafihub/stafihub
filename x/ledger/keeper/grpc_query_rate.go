package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q Querier) GetExchangeRate(goCtx context.Context, req *types.QueryGetExchangeRateRequest) (*types.QueryGetExchangeRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := q.Keeper.GetExchangeRate(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryGetExchangeRateResponse{ExchangeRate: val}, nil
}

func (q Querier) ExchangeRateAll(goCtx context.Context, req *types.QueryExchangeRateAllRequest) (*types.QueryExchangeRateAllResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	exchangeRates := q.GetAllExchangeRate(ctx)

	return &types.QueryExchangeRateAllResponse{ExchangeRates: exchangeRates}, nil
}

func (q Querier) GetEraExchangeRate(goCtx context.Context, req *types.QueryGetEraExchangeRateRequest) (*types.QueryGetEraExchangeRateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := q.Keeper.GetEraExchangeRate(ctx, req.Denom, req.Era)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryGetEraExchangeRateResponse{EraExchangeRate: val}, nil
}

func (q Querier) EraExchangeRatesByDenom(goCtx context.Context, req *types.QueryEraExchangeRatesByDenomRequest) (*types.QueryEraExchangeRatesByDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	eers := q.GetEraExchangeRateByDenom(ctx, req.Denom)

	return &types.QueryEraExchangeRatesByDenomResponse{EraExchangeRates: eers}, nil
}
