package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/relayers/types"
)

func (k Keeper) AddRelayer(ctx sdk.Context, arena, denom, addr string) {
	rel, _ := k.GetRelayer(ctx, arena, denom)
	rel.Addrs = append(rel.Addrs, addr)
	k.setRelayer(ctx, rel)
}

func (k Keeper) HasRelayer(ctx sdk.Context, arena, denom, addr string) bool {
	rel, ok := k.GetRelayer(ctx, arena, denom)
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

func (k Keeper) RemoveRelayer(ctx sdk.Context, arena, denom, addr string) {
	rel, ok := k.GetRelayer(ctx, arena, denom)
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
	store.Set([]byte(rel.Arena+rel.Denom), b)
}

func (k Keeper) GetRelayer(ctx sdk.Context, arena, denom string) (types.Relayer, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	val := types.Relayer{Arena: arena, Denom: denom, Addrs: []string{}}
	b := store.Get([]byte(arena + denom))

	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

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
