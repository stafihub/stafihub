package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		sudoKeeper    types.SudoKeeper
		bankKeeper    types.BankKeeper
		relayerKeeper types.RelayersKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

	sudoKeeper types.SudoKeeper,
	bankKeeper types.BankKeeper,
	relayerKeeper types.RelayersKeeper,

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

func (k Keeper) SetRValidatorIndicator(ctx sdk.Context, rvi types.RValidatorIndicator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RValidatorIndicatorPrefix)

	b := k.cdc.MustMarshal(&rvi)
	store.Set([]byte(rvi.Denom), b)
}

func (k Keeper) GetRValidatorIndicator(ctx sdk.Context, denom string) (val types.RValidatorIndicator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RValidatorIndicatorPrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) AddRValidator(ctx sdk.Context, rv types.RValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RValidatorPrefix)

	b := k.cdc.MustMarshal(&rv)
	store.Set([]byte(rv.Denom+rv.Address), b)
}

func (k Keeper) GetRValidator(ctx sdk.Context, denom, address string) (val types.RValidator, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RValidatorPrefix)
	b := store.Get([]byte(denom + address))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) AddRValidatorToSet(ctx sdk.Context, rv types.RValidator) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RValidatorSetPrefix)
	rset, _ := k.GetRValidatorSet(ctx, rv.Denom)
	rvStr := rv.Denom + rv.Address
	switch rv.Status {
	case types.Onboard:
		rset.Onboards = append(rset.Onboards, rvStr)
	case types.Pickable:
		rset.Pickables = append(rset.Onboards, rvStr)
	case types.Picked:
		rset.Onboards = append(rset.Onboards, rvStr)
	case types.Unpickable:
		rset.Pickables = append(rset.Onboards, rvStr)
	case types.Offboard:
		rset.Pickables = append(rset.Onboards, rvStr)
	}

	b := k.cdc.MustMarshal(&rset)
	store.Set([]byte(rset.Denom), b)
}

func (k Keeper) GetRValidatorSet(ctx sdk.Context, denom string) (types.RValidatorSet, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RValidatorSetPrefix)

	val := types.RValidatorSet{Denom: denom, Onboards: []string{}, Pickables: []string{},
		Pickeds: []string{}, Unpickables: []string{}, Offboards: []string{}}
	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) UpdateRValidatorSet(ctx sdk.Context, rset types.RValidatorSet) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RValidatorSetPrefix)

	b := k.cdc.MustMarshal(&rset)
	store.Set([]byte(rset.Denom), b)
}
