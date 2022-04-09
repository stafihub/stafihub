package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rbank/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AddressPrefix(goCtx context.Context, req *types.QueryAddressPrefixRequest) (*types.QueryAddressPrefixResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	prefix, found := k.GetAddressPrefix(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryAddressPrefixResponse{
		AddressPrefix: prefix,
	}, nil
}
