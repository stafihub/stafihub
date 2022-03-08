package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RelayFee(goCtx context.Context, req *types.QueryRelayFeeRequest) (*types.QueryRelayFeeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryRelayFeeResponse{
		Value: k.GetRelayFee(ctx, uint8(req.ChainId)),
	}, nil
}
