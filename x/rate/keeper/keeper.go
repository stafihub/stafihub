package keeper

import (
	"encoding/binary"
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/rate/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
	}
)

func NewKeeper(
    cdc codec.BinaryCodec,
    storeKey,
    memKey sdk.StoreKey,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Set the rate
func (k Keeper) SetRate(ctx sdk.Context, denom string, total, rtotal sdk.Int) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	rateStore := prefix.NewStore(store, types.RatePrefix)

	dec := sdk.OneDec()
	if total.Int64() != 0 && rtotal.Int64() != 0 {
		dec = dec.MulInt(rtotal).QuoInt(total)
	}

	bz, err := dec.Marshal()
	if err != nil {
		panic(fmt.Errorf("unable to marshal rate value %v", err))
	}

	rateStore.Set([]byte(denom), bz)
	return dec
}

// Load the rate
func (k Keeper) GetRate(ctx sdk.Context, denom string) *sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	rateStore := prefix.NewStore(store, types.RatePrefix)

	bz := rateStore.Get([]byte(denom))
	if bz == nil {
		return nil
	}

	var dec sdk.Dec
	err := dec.Unmarshal(bz)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal rate value %v", err))
	}

	return &dec
}

// Set the era rate.
func (k Keeper) SetEraRate(ctx sdk.Context, denom string, era uint32, dec sdk.Dec) {
	intBytes, err := dec.Marshal()
	if err != nil {
		panic(fmt.Errorf("unable to marshal era rate value %v", err))
	}

	store := ctx.KVStore(k.storeKey)
	rateStore := prefix.NewStore(store, types.RatePrefix)

	key := []byte(denom)
	binary.BigEndian.PutUint32(key, era)
	rateStore.Set(key, intBytes)
}

// Load the era rate.
func (k Keeper) GetEraRate(ctx sdk.Context, denom string, era uint32) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	rateStore := prefix.NewStore(store, types.EraRatePrefix)

	key := []byte(denom)
	binary.BigEndian.PutUint32(key, era)
	bz := rateStore.Get(key)
	if bz == nil {
		return sdk.OneDec()
	}

	var ratio sdk.Dec
	err := ratio.Unmarshal(bz)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal rate value %v", err))
	}

	return ratio
}

// token amount to rtoken amount
func (k Keeper) TokenToRtoken(ctx sdk.Context, denom string, balance sdk.Int) sdk.Int {
	rate := k.GetRate(ctx, denom)

	return rate.MulInt(balance).TruncateInt()
}

// token amount to rtoken amount
func (k Keeper) RtokenToToken(ctx sdk.Context, denom string, rbalance sdk.Int) sdk.Int {
	rate := k.GetRate(ctx, denom)

	return sdk.OneDec().MulInt(rbalance).Quo(rate).TruncateInt()
	//return sdk.NewInt(i.Int64())
}






