package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/rvote/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	relayerstypes "github.com/stafiprotocol/stafihub/x/relayers/types"
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
			Status: types.StatusInitiated,
			StartBlock: curBlock,
			VotesFor: make([]string, 0),
			VotesAgainst: make([]string, 0),
		}
		prop.ExpireBlock = prop.StartBlock + k.ProposalLife(ctx)
		if err := prop.SetContent(content); err != nil {
			return nil, err
		}
	}

	if prop.HasVoted(proposer) {
		return nil, relayerstypes.ErrAlreadyVoted
	}

	if prop.InFavour() {
		prop.VotesFor = append(prop.VotesFor, proposer)
	} else {
		prop.VotesAgainst = append(prop.VotesAgainst, proposer)
	}

	if prop.IsExpired(curBlock) {
		prop.Status = types.StatusExpired
	} else {
		total := uint32(k.relayerKeeper.RelayerCount(ctx, content.GetDenom()))
		if threshold.Value > total || uint32(len(prop.VotesAgainst)) + threshold.Value > total {
			prop.Status = types.StatusRejected
		} else if uint32(len(prop.VotesFor)) > threshold.Value {
			prop.Status = types.StatusApproved
		}
	}

	k.SetProposal(ctx, prop)
	return prop, nil
}

func (k Keeper) SetProposal(ctx sdk.Context, proposal *types.Proposal) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ProposalPrefix)
	b := k.cdc.MustMarshal(proposal)
	store.Set(proposal.PropId(), b)
}


func (k Keeper) GetProposal(ctx sdk.Context, id []byte) (val *types.Proposal, found bool) {
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


