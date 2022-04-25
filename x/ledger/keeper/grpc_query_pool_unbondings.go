package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PoolUnbondings(goCtx context.Context, req *types.QueryPoolUnbondingsRequest) (*types.QueryPoolUnbondingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	total := k.GetPoolUnbondNextSequence(ctx, req.Denom, req.Pool, req.UnlockEra)
	unbondings := make([]*types.Unbonding, total)
	for i := uint32(0); i < total; i++ {
		unbonding, _ := k.GetPoolUnbonding(ctx, req.Denom, req.Pool, req.UnlockEra, i)
		unbondings[i] = unbonding
	}

	return &types.QueryPoolUnbondingsResponse{
		Unbondings: unbondings,
	}, nil
}
