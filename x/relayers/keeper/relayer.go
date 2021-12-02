package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// SetRelayer set a specific relayer in the store from its index
func (k Keeper) SetRelayer(ctx sdk.Context, relayer types.Relayer) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	b := k.cdc.MustMarshal(&relayer)
	store.Set([]byte(relayer.Denom+relayer.Address), b)
}

// GetRelayer returns a relayer from its index
func (k Keeper) CheckIsRelayer(ctx sdk.Context, denom, address string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)

	b := store.Get([]byte(denom+address))
    return b != nil
}

// RemoveRelayer removes a relayer from the store
func (k Keeper) RemoveRelayer(ctx sdk.Context, denom, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	store.Delete([]byte(denom+address))
}

// GetAllRelayer returns all relayer
func (k Keeper) GetAllRelayer(ctx sdk.Context) (list []types.Relayer) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Relayer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}

// GetAllRelayer returns all relayer
func (k Keeper) GetRelayersByDenom(ctx sdk.Context, denom string) (list []types.Relayer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Relayer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Denom == denom {
			list = append(list, val)
		}
	}

	return
}