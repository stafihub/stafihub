package keeper

import (
	"fmt"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
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

// SetAdmin set the admin account in the store
func (k Keeper) SetAdmin(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.AdminPrefix, address)
}

// GetRelayer returns a relayer from its index
func (k Keeper) GetAdmin(ctx sdk.Context) sdk.AccAddress {
	store := ctx.KVStore(k.storeKey)
	return store.Get(types.AdminPrefix)
}