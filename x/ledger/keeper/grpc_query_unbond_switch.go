package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UnbondSwitch(goCtx context.Context, req *types.QueryUnbondSwitchRequest) (*types.QueryUnbondSwitchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	isOpen := k.GetUnbondSwitch(ctx, req.Denom)

	return &types.QueryUnbondSwitchResponse{
		IsOpen: isOpen,
	}, nil
}
