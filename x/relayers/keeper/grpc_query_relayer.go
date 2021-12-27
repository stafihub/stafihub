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

func (k Keeper) RelayersByDenom(c context.Context,  req *types.QueryRelayersByDenomRequest) (*types.QueryRelayersByDenomResponse, error) {
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