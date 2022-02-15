package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/sudo/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey

		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	bankKeeper types.BankKeeper,
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,

		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetAdmin(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.AdminPrefix, address)
}

func (k Keeper) GetAdmin(ctx sdk.Context) sdk.AccAddress {
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.AdminPrefix)
}

func (k Keeper) IsAdmin(ctx sdk.Context, address string) bool {
	admin := k.GetAdmin(ctx)
	return admin.String() == address
}

func (k Keeper) SetInflationBase(ctx sdk.Context, inflationBase sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	bts, err := inflationBase.Marshal()
	if err != nil {
		panic(fmt.Errorf("unable to marshal amount value %v", err))
	}
	store.Set(types.InflationBasePrefix, bts)
}

func (k Keeper) GetInflationBase(ctx sdk.Context) sdk.Int {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.InflationBasePrefix)
	var amount sdk.Int
	err := amount.Unmarshal(bts)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal supply value %v", err))
	}
	return amount
}
