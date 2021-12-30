package keeper

import (
	"context"

    "github.com/stafiprotocol/stafihub/x/rvote/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) SubmitProposal(goCtx context.Context,  msg *types.MsgSubmitProposal) (*types.MsgSubmitProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgSubmitProposalResponse{}, nil
}
