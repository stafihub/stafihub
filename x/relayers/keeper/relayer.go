package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/relayers/types"
)

func (k Keeper) AddRelayer(ctx sdk.Context, denom, addr string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	rel, _ := k.GetRelayerByDenom(ctx, denom)
	rel.Addrs = append(rel.Addrs, addr)
	b := k.cdc.MustMarshal(&rel)
	store.Set([]byte(denom), b)
}

func (k Keeper) IsRelayer(ctx sdk.Context, denom, addr string) bool {
	rel, ok := k.GetRelayerByDenom(ctx, denom)
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

func (k Keeper) RemoveRelayer(ctx sdk.Context, denom, addr string) {
	rel, ok := k.GetRelayerByDenom(ctx, denom)
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	b := k.cdc.MustMarshal(&rel)
	store.Set([]byte(denom), b)
}

func (k Keeper) GetRelayerByDenom(ctx sdk.Context, denom string) (types.Relayer, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	val := types.Relayer{Denom: denom, Addrs: []string{}}
	b := store.Get([]byte(denom))

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
