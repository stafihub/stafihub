package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/rvote/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Querier struct {
	Keeper
}

var _ types.QueryServer = Querier{}

func (q Querier) GetProposal(goCtx context.Context, req *types.QueryGetProposalRequest) (*types.QueryGetProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	prop, found := q.Keeper.GetProposal(ctx, req.PropId)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetProposalResponse{Proposal: prop.String()}, nil
}

func (k Keeper) GetProposalLife(goCtx context.Context, req *types.QueryGetProposalLifeRequest) (*types.QueryGetProposalLifeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	pl := k.ProposalLife(ctx)

	return &types.QueryGetProposalLifeResponse{ProposalLife: pl}, nil
}
