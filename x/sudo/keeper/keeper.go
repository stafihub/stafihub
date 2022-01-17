package keeper

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/sudo/types"
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

func (k Keeper) AddDenom(ctx sdk.Context, denom string) {
	store := ctx.KVStore(k.storeKey)
	sym := k.GetSymbol(ctx)
	sym.Denoms[denom] = true
	b := k.cdc.MustMarshal(&sym)
	store.Set(types.SymbolPrefix, b)
}

func (k Keeper) GetSymbol(ctx sdk.Context) (val types.Symbol) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.SymbolPrefix)
	if b == nil {
		return types.Symbol{Denoms: map[string]bool{}}
	}
	k.cdc.MustUnmarshal(b, &val)
	return
}

func (k Keeper) GetAllDenoms(ctx sdk.Context) []string {
	sym := k.GetSymbol(ctx)
	denoms := make([]string, 0)
	for denom, ok := range sym.Denoms {
		if ok {
			denoms = append(denoms, denom)
		}
	}

	return denoms
}

func (k Keeper) IsDenomValid(ctx sdk.Context, denom string) bool {
	sym := k.GetSymbol(ctx)
	_, ok := sym.Denoms[denom]
	return ok
}

