package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StakeRecord(goCtx context.Context, req *types.QueryStakeRecordRequest) (*types.QueryStakeRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	stakeRecord, found := k.GetUserStakeRecord(ctx, req.UserAddress, req.StakeTokenDenom, req.StakeRecordIndex)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}
	return &types.QueryStakeRecordResponse{
		StakeRecord: stakeRecord,
	}, nil
}
