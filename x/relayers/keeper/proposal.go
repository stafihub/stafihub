package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// SetProposal set a specific proposal in the store from its index
func (k Keeper) SetProposal(ctx sdk.Context, proposal *types.Proposal) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ProposalPrefix)
	b := k.cdc.MustMarshal(proposal)
	store.Set(proposal.ProposalId, b)
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
