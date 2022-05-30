package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rdex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ProviderSwitch(goCtx context.Context, req *types.QueryProviderSwitchRequest) (*types.QueryProviderSwitchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryProviderSwitchResponse{
		ProviderSwitch: k.GetProviderSwitch(ctx),
	}, nil
}
