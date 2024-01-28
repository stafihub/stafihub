package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LatestLsmBondProposalId(goCtx context.Context, req *types.QueryLatestLsmBondProposalIdRequest) (*types.QueryLatestLsmBondProposalIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	id, found := k.GetLatestLsmBondProposalId(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryLatestLsmBondProposalIdResponse{ProposalId: id}, nil
}
