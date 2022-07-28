package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/claim/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ClaimSwitch(goCtx context.Context, req *types.QueryClaimSwitchRequest) (*types.QueryClaimSwitchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryClaimSwitchResponse{
		IsOpen: k.GetClaimSwitch(ctx, req.Round),
	}, nil
}
