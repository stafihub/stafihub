package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/rate/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) SetExchangeRate(ctx sdk.Context, denom string, rate sdk.Dec) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ExchangeRateKeyPrefix)
	e := types.ExchangeRate{
		Denom: denom,
		Value: rate,
	}
	b := k.cdc.MustMarshal(&e)
	store.Set([]byte(denom), b)
}

func (k Keeper) GetExchangeRate(ctx sdk.Context, denom string) (val types.ExchangeRate, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ExchangeRateKeyPrefix)

	b := store.Get([]byte(denom))
    if b == nil {
        return val, false
    }

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetAllExchangeRate returns all exchangeRate
func (k Keeper) GetAllExchangeRate(ctx sdk.Context) (list []types.ExchangeRate) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ExchangeRateKeyPrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ExchangeRate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}
