package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShuffleSeconds(goCtx context.Context, req *types.QueryShuffleSecondsRequest) (*types.QueryShuffleSecondsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryShuffleSecondsResponse{
		ShuffleSeconds: k.GetShuffleSeconds(ctx, req.Denom),
	}, nil
}
