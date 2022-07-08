package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IcaAccountInfo(goCtx context.Context, req *types.QueryIcaAccountInfoRequest) (*types.QueryIcaAccountInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	ica, found := k.GetICAAccount(ctx, req.Owner)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryIcaAccountInfoResponse{IcaAccount: &ica}, nil
}
