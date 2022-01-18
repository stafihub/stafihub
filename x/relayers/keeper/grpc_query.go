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

var _ types.QueryServer = Keeper{}

func (k Keeper) RelayerAll(c context.Context, req *types.QueryAllRelayerRequest) (*types.QueryAllRelayerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var relayers []types.Relayer
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	relayerStore := prefix.NewStore(store, types.RelayerPrefix)

	pageRes, err := query.Paginate(relayerStore, req.Pagination, func(key []byte, value []byte) error {
		var relayer types.Relayer
		if err := k.cdc.Unmarshal(value, &relayer); err != nil {
			return err
		}

		relayers = append(relayers, relayer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRelayerResponse{Relayers: relayers, Pagination: pageRes}, nil
}

func (k Keeper) RelayersByDenom(c context.Context, req *types.QueryRelayersByDenomRequest) (*types.QueryRelayersByDenomResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var relayers []types.Relayer
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	relayerStore := prefix.NewStore(store, types.RelayerPrefix)

	pageRes, err := query.Paginate(relayerStore, req.Pagination, func(key []byte, value []byte) error {
		var relayer types.Relayer
		if err := k.cdc.Unmarshal(value, &relayer); err != nil {
			return err
		}

		if relayer.Denom == req.Denom {
			relayers = append(relayers, relayer)
		}

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryRelayersByDenomResponse{Relayers: relayers, Pagination: pageRes}, nil
}

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
