package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) SubmitProposal(ctx sdk.Context,  content *types.ProposalContent, proposer string, inFavour bool) (*types.Proposal, error) {
	threshold, ok := k.GetThreshold(ctx, content.Denom)
	if !ok {
		return nil, types.ErrThresholdNotFound
	}

	curBlock := ctx.BlockHeight()
	prop, ok := k.GetProposal(ctx, content.PropId)
	if !ok {
		prop = &types.Proposal{
			PropId: content.PropId,
			Content: content,
			Status: types.StatusInitiated,
			StartBlock: curBlock,
			VotesFor: make([]string, 0),
			VotesAgainst: make([]string, 0),
		}
		prop.ExpireBlock = prop.StartBlock + k.ProposalLife(ctx)
	}

	if inFavour {
		prop.VotesFor = append(prop.VotesFor, proposer)
	} else {
		prop.VotesAgainst = append(prop.VotesAgainst, proposer)
	}

	if prop.IsExpired(curBlock) {
		prop.Status = types.StatusExpired
	} else {
		total := uint32(k.RelayerCount(ctx, content.Denom))
		if threshold.Value > total || uint32(len(prop.VotesAgainst)) + threshold.Value > total {
			prop.Status = types.StatusRejected
		} else if uint32(len(prop.VotesFor)) > threshold.Value {
			prop.Status = types.StatusApproved
		}
	}

	k.SetProposal(ctx, prop)
	return prop, nil
}

// SetProposal set a specific proposal in the store from its index
func (k Keeper) SetProposal(ctx sdk.Context, proposal *types.Proposal) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ProposalPrefix)
	b := k.cdc.MustMarshal(proposal)
	store.Set(proposal.PropId, b)
}

// GetProposal returns a proposal from its index
func (k Keeper) GetProposal(
    ctx sdk.Context,
	id []byte,
) (val *types.Proposal, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ProposalPrefix)

	b := store.Get(id)
    if b == nil {
        return val, false
    }

	k.cdc.MustUnmarshal(b, val)
	return val, true
}

// GetAllProposal returns all proposal
func (k Keeper) GetAllProposal(ctx sdk.Context) (list []types.Proposal) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ProposalPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Proposal
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}
