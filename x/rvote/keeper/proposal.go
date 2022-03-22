package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvote/types"
)

func (k Keeper) SetProposal(ctx sdk.Context, proposal *types.Proposal) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ProposalPrefix)
	b := k.cdc.MustMarshal(proposal)
	store.Set([]byte(proposal.PropId()), b)
}

func (k Keeper) GetProposal(ctx sdk.Context, propId string) (val *types.Proposal, found bool) {
	val = &types.Proposal{}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ProposalPrefix)

	b := store.Get([]byte(propId))
	if b == nil {
		return nil, false
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
