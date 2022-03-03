package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/relayers/types"
)


func (k Keeper) AddRelayer(ctx sdk.Context, taipe, denom, addr string) {
	rel, _ := k.GetRelayer(ctx, taipe, denom)
	rel.Addrs = append(rel.Addrs, addr)
	k.setRelayer(ctx, rel)
}

func (k Keeper) HasRelayer(ctx sdk.Context, taipe, denom, addr string) bool {
	rel, ok := k.GetRelayer(ctx, taipe, denom)
	if !ok {
		return false
	}

	for _, adr := range rel.Addrs {
		if adr == addr {
			return true
		}
	}

	return false
}

func (k Keeper) RemoveRelayer(ctx sdk.Context, taipe, denom, addr string) {
	rel, ok := k.GetRelayer(ctx, taipe, denom)
	if !ok {
		return
	}

	addrs := make([]string, 0)
	for _, adr := range rel.Addrs {
		if adr != addr {
			addrs = append(addrs, adr)
		}
	}
	rel.Addrs = addrs
	k.setRelayer(ctx, rel)
}

func (k Keeper) setRelayer(ctx sdk.Context, rel types.Relayer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	b := k.cdc.MustMarshal(&rel)
	store.Set([]byte(rel.Taipe+rel.Denom), b)
}

func (k Keeper) GetRelayer(ctx sdk.Context, taipe, denom string) (types.Relayer, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	val := types.Relayer{Taipe: taipe, Denom: denom, Addrs: []string{}}
	b := store.Get([]byte(taipe+denom))

	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetRelayersByTaipeAndDenom(ctx sdk.Context, taipe, denom string) (list []types.Relayer) {
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
