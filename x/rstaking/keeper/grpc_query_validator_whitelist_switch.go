package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rstaking/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ValidatorWhitelistSwitch(goCtx context.Context, req *types.QueryValidatorWhitelistSwitchRequest) (*types.QueryValidatorWhitelistSwitchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	isOpen := k.GetValidatorWhitelistSwitch(ctx)

	return &types.QueryValidatorWhitelistSwitchResponse{
		IsOpen: isOpen,
	}, nil
}
