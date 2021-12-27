package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	gogotypes "github.com/gogo/protobuf/types"
)

func (k Keeper) SetRelayer(ctx sdk.Context, relayer *types.Relayer) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	b := k.cdc.MustMarshal(relayer)
	store.Set([]byte(relayer.Denom+relayer.Address), b)
	k.IncRelayerCount(ctx, relayer.Denom)
}

func (k Keeper) IsRelayer(ctx sdk.Context, denom, address string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)

	b := store.Get([]byte(denom+address))
	var r types.Relayer
	k.cdc.MustUnmarshal(b, &r)
	return b != nil
}

func (k Keeper) IncRelayerCount(ctx sdk.Context, denom string) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerCountPrefix)
	old := k.RelayerCount(ctx, denom)
	b := k.cdc.MustMarshal(&gogotypes.Int32Value{Value: old + 1})
	store.Set([]byte(denom), b)
}

func (k Keeper) RelayerCount(ctx sdk.Context, denom string) int32 {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerCountPrefix)
	b := store.Get([]byte(denom))
	intV := gogotypes.Int32Value{}
	err := k.cdc.Unmarshal(b, &intV)
	if err != nil {
		return 0
	}
	return intV.GetValue()
}

func (k Keeper) RemoveRelayer(ctx sdk.Context, denom, address string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	store.Delete([]byte(denom+address))
}

func (k Keeper) GetAllRelayer(ctx sdk.Context) (list []*types.Relayer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RelayerPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Relayer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, &val)
	}

    return
}
