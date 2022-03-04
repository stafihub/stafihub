package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/relayers/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Relayers(c context.Context, req *types.QueryRelayersRequest) (*types.QueryRelayersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(c)
	rel, found := k.GetRelayer(ctx, req.Arena, req.Denom)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryRelayersResponse{Relayers: rel.Addrs}, nil
}

func (k Keeper) Threshold(c context.Context, req *types.QueryGetThresholdRequest) (*types.QueryGetThresholdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}
	ctx := sdk.UnwrapSDKContext(c)
	val, found := k.GetThreshold(
		ctx,
		req.Arena,
		req.Denom,
	)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryGetThresholdResponse{Threshold: val.Value}, nil
}
