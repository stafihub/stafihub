package keeper

import (
	"context"
	"github.com/tendermint/tendermint/crypto"

	"github.com/stafiprotocol/stafihub/x/relayers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) SubmitProposal(goCtx context.Context,  msg *types.MsgSubmitProposal) (*types.MsgSubmitProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.CheckIsRelayer(ctx, msg.Denom, msg.Creator) {
		return nil, types.ErrCreatorNotRelayer
	}

	content := msg.Content()
	cbz := k.cdc.MustMarshal(content)
	id := crypto.Sha256(cbz)
	threshold, ok := k.GetThreshold(ctx, content.Denom)
	if !ok {
		return nil, types.ErrThresholdNotFound
	}
	curBlock := ctx.BlockHeight()
	prop, ok := k.GetProposal(ctx, id)
	if !ok {
		prop = &types.Proposal{
			ProposalId: id,
			Content: content,
			Status: types.StatusInitiated,
			StartBlock: curBlock,
			VotesFor: make([]string, 0),
			VotesAgainst: make([]string, 0),
		}
		prop.ExpireBlock = prop.StartBlock + k.ProposalLife(ctx)
	}

	if msg.InFavour {
		prop.VotesFor = append(prop.VotesFor, msg.Creator)
	} else {
		prop.VotesAgainst = append(prop.VotesAgainst, msg.Creator)
	}

	if prop.IsExpired(curBlock) {
		prop.Status = types.StatusExpired
	} else {
		total := uint32(k.RelayerCount(ctx, msg.Denom))
		if threshold.Value > total || uint32(len(prop.VotesAgainst)) + threshold.Value > total {
			prop.Status = types.StatusRejected
		} else if uint32(len(prop.VotesFor)) > threshold.Value {
			prop.Status = types.StatusApproved
		}
	}

	k.SetProposal(ctx, prop)

	if prop.Status == types.StatusApproved {
	}


	return &types.MsgSubmitProposalResponse{}, nil
}
