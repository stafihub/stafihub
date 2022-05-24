package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StakeRecordList(goCtx context.Context, req *types.QueryStakeRecordListRequest) (*types.QueryStakeRecordListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryStakeRecordListResponse{
		StakeRecordList: k.GetUserStakeRecordList(ctx, req.UserAddress, req.StakePoolIndex),
	}, nil
}
