package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	"github.com/stafihub/stafihub/utils"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		sudoKeeper    types.SudoKeeper
		bankKeeper    types.BankKeeper
		relayerKeeper types.RelayerKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	sudoKeeper types.SudoKeeper,
	bankKeeper types.BankKeeper,
	relayerKeeper types.RelayerKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		sudoKeeper:    sudoKeeper,
		bankKeeper:    bankKeeper,
		relayerKeeper: relayerKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetExchangeRate(ctx sdk.Context, denom string, total, rtotal sdk.Int) {
	dec := utils.OneDec()
	if total.Int64() != 0 && rtotal.Int64() != 0 {
		dec = dec.MulInt(total).QuoInt(rtotal)
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ExchangeRateKeyPrefix)
	e := types.ExchangeRate{
		Denom: denom,
		Value: dec,
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

func (k Keeper) SetEraExchangeRate(ctx sdk.Context, denom string, era uint32, rate utils.Dec) {
	pre := append(types.EraExchangeRateKeyPrefix, types.KeyPrefix(denom)...)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pre)
	e := types.EraExchangeRate{
		Denom: denom,
		Era:   era,
		Value: rate,
	}
	b := k.cdc.MustMarshal(&e)

	bera := make([]byte, 4)
	binary.LittleEndian.PutUint32(bera, era)
	store.Set(bera, b)
}

func (k Keeper) GetEraExchangeRate(ctx sdk.Context, denom string, era uint32) (val types.EraExchangeRate, found bool) {
	pre := append(types.EraExchangeRateKeyPrefix, types.KeyPrefix(denom)...)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pre)
	bera := make([]byte, 4)
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pre)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EraExchangeRate
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// token to rtoken
func (k Keeper) TokenToRtoken(ctx sdk.Context, denom string, balance sdk.Int) sdk.Int {
	er, ok := k.GetExchangeRate(ctx, denom)
	if !ok {
		return balance
	}

	return utils.OneDec().MulInt(balance).Quo(er.Value).TruncateInt()
}

// rtoken to token
func (k Keeper) RtokenToToken(ctx sdk.Context, denom string, rbalance sdk.Int) sdk.Int {
	er, ok := k.GetExchangeRate(ctx, denom)
	if !ok {
		return rbalance
	}

	return er.Value.MulInt(rbalance).TruncateInt()
}

func (k Keeper) IncreaseTotalProtocolFee(ctx sdk.Context, denom string, increase sdk.Int) {
	total, found := k.GetTotalProtocolFee(ctx, denom)
	if !found {
		total = types.TotalProtocolFee{
			Denom: denom,
			Value: sdk.ZeroInt(),
		}
	}
	total.Value = total.Value.Add(increase)
	k.SetTotalProtocolFee(ctx, denom, total.Value)
}

func (k Keeper) SetTotalProtocolFee(ctx sdk.Context, denom string, total sdk.Int) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TotalProtocolFeePrefix)
	e := types.TotalProtocolFee{
		Denom: denom,
		Value: total,
	}
	b := k.cdc.MustMarshal(&e)
	store.Set([]byte(denom), b)
}

func (k Keeper) GetTotalProtocolFee(ctx sdk.Context, denom string) (val types.TotalProtocolFee, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TotalProtocolFeePrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetAllTotalProtocolFee(ctx sdk.Context) (list []types.TotalProtocolFee) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.TotalProtocolFeePrefix)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TotalProtocolFee
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}
	return
}
