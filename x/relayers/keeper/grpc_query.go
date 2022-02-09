package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/relayers/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) RelayersByDenom(c context.Context, req *types.QueryRelayersByDenomRequest) (*types.QueryRelayersByDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	rel, _ := k.GetRelayerByDenom(ctx, req.Denom)

	return &types.QueryRelayersByDenomResponse{Relayers: rel.Addrs}, nil
}

func (k Keeper) Threshold(c context.Context, req *types.QueryGetThresholdRequest) (*types.QueryGetThresholdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	val, found := k.GetThreshold(
		ctx,
		req.Denom,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetThresholdResponse{Threshold: val.Value}, nil
}
