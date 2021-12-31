package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// SetThreshold set a specific threshold in the store from its denom
func (k Keeper) SetThreshold(ctx sdk.Context, threshold *types.Threshold) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ThresholdPrefix)
	b := k.cdc.MustMarshal(threshold)
	store.Set([]byte(threshold.Denom), b)
}

// GetThreshold returns a threshold from its index
func (k Keeper) GetThreshold(ctx sdk.Context, denom string) (val types.Threshold, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ThresholdPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllThreshold returns all threshold
func (k Keeper) GetAllThreshold(ctx sdk.Context) (list []*types.Threshold) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ThresholdPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Threshold
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, &val)
	}

    return
}
