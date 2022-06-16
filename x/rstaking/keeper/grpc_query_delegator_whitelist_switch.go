package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rstaking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DelegatorWhitelistSwitch(goCtx context.Context, req *types.QueryDelegatorWhitelistSwitchRequest) (*types.QueryDelegatorWhitelistSwitchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryDelegatorWhitelistSwitchResponse{
		IsOpen: k.GetDelegatorWhitelistSwitch(ctx),
	}, nil
}
