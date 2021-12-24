package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/sudo/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Admin(goCtx context.Context,  req *types.QueryAdminRequest) (*types.QueryAdminResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	admin := k.GetAdmin(ctx)

	return &types.QueryAdminResponse{Address: admin.String()}, nil
}

func (k Keeper) AllDenoms(goCtx context.Context,  req *types.QueryAllDenomsRequest) (*types.QueryAllDenomsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	return &types.QueryAllDenomsResponse{Denoms: k.GetAllDenoms(ctx)}, nil
}