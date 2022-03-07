package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
)

func (k msgServer) VoteStatusProposal(goCtx context.Context, msg *types.MsgVoteStatusProposal) (*types.MsgVoteStatusProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgVoteStatusProposalResponse{}, nil
}
