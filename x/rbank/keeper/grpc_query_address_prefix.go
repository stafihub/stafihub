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

	accPrefix, found := k.GetAccAddressPrefix(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}
	valPrefix, found := k.GetValAddressPrefix(ctx, req.Denom)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryAddressPrefixResponse{
		AccAddressPrefix: accPrefix,
		ValAddressPrefix: valPrefix,
	}, nil
}
