package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RelayFeeReceiver(goCtx context.Context, req *types.QueryRelayFeeReceiverRequest) (*types.QueryRelayFeeReceiverResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	receiver, found := k.GetRelayFeeReceiver(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}
	return &types.QueryRelayFeeReceiverResponse{
		Receiver: receiver.String(),
	}, nil
}
