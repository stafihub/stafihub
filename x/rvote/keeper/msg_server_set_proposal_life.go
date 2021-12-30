package keeper

import (
	"context"

    "github.com/stafiprotocol/stafihub/x/rvote/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) SetProposalLife(goCtx context.Context,  msg *types.MsgSetProposalLife) (*types.MsgSetProposalLifeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgSetProposalLifeResponse{}, nil
}
