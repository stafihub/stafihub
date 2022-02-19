package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvote/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (q Querier) GetProposal(goCtx context.Context, req *types.QueryGetProposalRequest) (*types.QueryGetProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	prop, found := q.Keeper.GetProposal(ctx, req.PropId)
	if !found {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}

	return &types.QueryGetProposalResponse{Proposal: prop.String()}, nil
}

func (k Keeper) GetProposalLife(goCtx context.Context, req *types.QueryGetProposalLifeRequest) (*types.QueryGetProposalLifeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	pl := k.ProposalLife(ctx)

	return &types.QueryGetProposalLifeResponse{ProposalLife: pl}, nil
}
