package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	relayerstypes "github.com/stafiprotocol/stafihub/x/relayers/types"
	"github.com/stafiprotocol/stafihub/x/rvote/types"
)

func (k Keeper) SubmitProposal(ctx sdk.Context, content types.Content, proposer string) (*types.Proposal, error) {
	threshold, ok := k.relayerKeeper.GetThreshold(ctx, content.GetDenom())
	if !ok {
		return nil, relayerstypes.ErrThresholdNotFound
	}

	curBlock := ctx.BlockHeight()
	prop, ok := k.GetProposal(ctx, content.GetPropId())
	if !ok {
		prop = &types.Proposal{
			Status:     types.StatusInitiated,
			StartBlock: curBlock,
			Voted:      []string{proposer},
		}
		prop.ExpireBlock = prop.StartBlock + k.ProposalLife(ctx)
		if err := prop.SetContent(content); err != nil {
			return nil, err
		}
	}

	if prop.HasVoted(proposer) {
		return nil, relayerstypes.ErrAlreadyVoted
	}

	if prop.IsExpired(curBlock) {
		prop.Status = types.StatusExpired
	} else {
		if uint32(len(prop.Voted)) > threshold.Value {
			prop.Status = types.StatusApproved
		}
	}

	k.SetProposal(ctx, prop)
	return prop, nil
}

func (k Keeper) SetProposal(ctx sdk.Context, proposal *types.Proposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ProposalPrefix)
	b := k.cdc.MustMarshal(proposal)
	store.Set(proposal.PropId(), b)
}

func (k Keeper) GetProposal(ctx sdk.Context, id []byte) (val *types.Proposal, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ProposalPrefix)

	b := store.Get(id)
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, val)
	return val, true
}

// GetAllProposal returns all proposal
func (k Keeper) GetAllProposal(ctx sdk.Context) (list []types.Proposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ProposalPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Proposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
