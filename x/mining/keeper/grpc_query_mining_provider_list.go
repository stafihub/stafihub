package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MiningProviderList(goCtx context.Context, req *types.QueryMiningProviderListRequest) (*types.QueryMiningProviderListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryMiningProviderListResponse{
		MiningProviderList: k.GetMiningProviderList(ctx),
	}, nil
}
