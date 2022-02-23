package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TotalFeeList(goCtx context.Context, req *types.QueryTotalFeeListRequest) (*types.QueryTotalFeeListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryTotalFeeListResponse{
		TotalFeeList: k.GetAllTotalFee(ctx),
	}, nil
}
