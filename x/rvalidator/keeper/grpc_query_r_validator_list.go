package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RValidatorList(goCtx context.Context, req *types.QueryRValidatorListRequest) (*types.QueryRValidatorListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	list := make([]string, 0)
	for _, val := range k.GetSelectedRValidatorListByDenomPoolAddress(ctx, req.Denom, req.PoolAddress) {
		list = append(list, val.ValAddress)
	}

	return &types.QueryRValidatorListResponse{
		RValidatorList: list,
	}, nil
}
