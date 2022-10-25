package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetDepositCount(goCtx context.Context, req *types.QueryGetDepositCountRequest) (*types.QueryGetDepositCountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	count := k.GetDepositCountById(ctx, uint8(req.ChainId))

	return &types.QueryGetDepositCountResponse{
		Count: uint32(count),
	}, nil
}
