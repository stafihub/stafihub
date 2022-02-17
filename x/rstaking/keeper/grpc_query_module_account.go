package keeper

import (
	"context"

	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stafihub/stafihub/x/rstaking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ModuleAccount(goCtx context.Context, req *types.QueryModuleAccountRequest) (*types.QueryModuleAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	return &types.QueryModuleAccountResponse{
		ModuleAccount: authTypes.NewModuleAddress(types.ModuleName).String(),
	}, nil
}
