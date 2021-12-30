package keeper

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/rate/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) SetEraExchangeRate(ctx sdk.Context, denom string, era uint32, rate sdk.Dec) {
	pre := append(types.EraExchangeRateKeyPrefix, types.KeyPrefix(denom)...)
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), pre)
	e := types.EraExchangeRate{
		Denom: denom,
		Era: era,
		Value: rate,
	}
	b := k.cdc.MustMarshal(&e)

	bera:= make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	store.Set(bera, b)
}

func (k Keeper) GetEraExchangeRate(ctx sdk.Context, denom string, era uint32) (val types.EraExchangeRate, found bool) {
	pre := append(types.EraExchangeRateKeyPrefix, types.KeyPrefix(denom)...)
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), pre)
	bera:= make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	b := store.Get(bera)
    if b == nil {
        return val, false
    }

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllEraExchangeRate returns all eraExchangeRate
func (k Keeper) GetEraExchangeRateByDenom(ctx sdk.Context, denom string) (list []types.EraExchangeRate) {
	pre := append(types.EraExchangeRateKeyPrefix, types.KeyPrefix(denom)...)
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), pre)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EraExchangeRate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}
