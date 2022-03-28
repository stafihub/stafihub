package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rmintreward/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserMintCount(goCtx context.Context, req *types.QueryUserMintCountRequest) (*types.QueryUserMintCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	count, found := k.GetUserMintCount(ctx, address, req.Denom, req.Cycle)
	if !found {
		count = 0
	}

	return &types.QueryUserMintCountResponse{
		Count: count,
	}, nil
}
