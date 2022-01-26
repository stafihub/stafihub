package keeper

import (
	"context"
	"encoding/hex"

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
	prop, ok := q.Keeper.GetProposal(ctx, req.PropId)
	if !ok {
		return nil, status.Errorf(codes.NotFound, "proposal %s not found", hex.EncodeToString(req.PropId))
	}

	return &types.QueryGetProposalResponse{Proposal: prop.String()}, nil
}
