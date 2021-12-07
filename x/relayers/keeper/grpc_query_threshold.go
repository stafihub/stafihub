package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ThresholdAll(c context.Context, req *types.QueryAllThresholdRequest) (*types.QueryAllThresholdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var thresholds []types.Threshold
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	thresholdStore := prefix.NewStore(store, types.ThresholdPrefix)

	pageRes, err := query.Paginate(thresholdStore, req.Pagination, func(key []byte, value []byte) error {
		var threshold types.Threshold
		if err := k.cdc.Unmarshal(value, &threshold); err != nil {
			return err
		}

		thresholds = append(thresholds, threshold)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllThresholdResponse{Threshold: thresholds, Pagination: pageRes}, nil
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

	return &types.QueryGetThresholdResponse{Threshold: val}, nil
}