package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
)

func (k Keeper) SetRelayer(ctx sdk.Context, denom, addr string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	rel, ok := k.GetRelayerByDenom(ctx, denom)
	if !ok {
		rel = types.NewRelayer(denom, addr)
	} else {
		rel.Addrs[addr] = true
	}

	b := k.cdc.MustMarshal(&rel)
	store.Set([]byte(denom), b)
}

func (k Keeper) IsRelayer(ctx sdk.Context, denom, addr string) bool {
	rel, ok := k.GetRelayerByDenom(ctx, denom)
	if !ok {
		return false
	}

	return rel.Addrs != nil && rel.Addrs[addr]
}

func (k Keeper) RemoveRelayer(ctx sdk.Context, denom, addr string) {
	rel, ok := k.GetRelayerByDenom(ctx, denom)
	if !ok || rel.Addrs == nil || !rel.Addrs[addr] {
		return
	}

	delete(rel.Addrs, addr)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	b := k.cdc.MustMarshal(&rel)
	store.Set([]byte(denom), b)
}

func (k Keeper) GetRelayerByDenom(ctx sdk.Context, denom string) (val types.Relayer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)

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


