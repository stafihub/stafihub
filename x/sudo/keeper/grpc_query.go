package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/sudo/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Admin(goCtx context.Context, req *types.QueryAdminRequest) (*types.QueryAdminResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	admin := k.GetAdmin(ctx)
	if admin == nil {
		return &types.QueryAdminResponse{}, nil
	}

	return &types.QueryAdminResponse{Address: admin.String()}, nil
}
