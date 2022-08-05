package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ResourceidToDenoms(goCtx context.Context, req *types.QueryResourceidToDenomsRequest) (*types.QueryResourceidToDenomsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryResourceidToDenomsResponse{
		ResourceidToDenoms: k.GetResourceIdToDenomList(ctx),
	}, nil
}
