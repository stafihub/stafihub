package keeper

import (
	"context"

    "github.com/stafiprotocol/stafihub/x/sudo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Admin(goCtx context.Context,  req *types.QueryAdminRequest) (*types.QueryAdminResponse, error) {
	if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }

	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Process the query
    _ = ctx

	return &types.QueryAdminResponse{}, nil
}
