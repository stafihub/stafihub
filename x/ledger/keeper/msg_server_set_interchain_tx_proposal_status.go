package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetInterchainTxProposalStatus(goCtx context.Context, msg *types.MsgSetInterchainTxProposalStatus) (*types.MsgSetInterchainTxProposalStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Status != types.InterchainTxStatusUnspecified && msg.Status != types.InterchainTxStatusInit &&
		msg.Status != types.InterchainTxStatusSuccess && msg.Status != types.InterchainTxStatusFailed {
		return nil, types.ErrInterchainTxStatusUnmatch
	}

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	_, found := k.Keeper.GetInterchainTxProposalStatus(ctx, msg.ProposalId)
	if !found {
		return nil, types.ErrInterchainTxPropIdNotFound
	}

	k.Keeper.SetInterchainTxProposalStatus(ctx, msg.ProposalId, msg.Status)

	return &types.MsgSetInterchainTxProposalStatusResponse{}, nil
}
