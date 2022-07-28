package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/claim/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IsClaimed(goCtx context.Context, req *types.QueryIsClaimedRequest) (*types.QueryIsClaimedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryIsClaimedResponse{
		IsClaimed: k.IsIndexClaimed(ctx, req.Round, req.Index),
	}, nil
}
