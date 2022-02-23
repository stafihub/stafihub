package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/sudo/types"
)

type (
	Keeper struct {
		cdc              codec.BinaryCodec
		storeKey         sdk.StoreKey
		memKey           sdk.StoreKey
		feeCollectorName string
		bankKeeper       types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	bankKeeper types.BankKeeper,
) *Keeper {
	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
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

func (k Keeper) SetAddressPrefix(ctx sdk.Context, denom, addrPrefix string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AddressPrefix)
	store.Set([]byte(denom), []byte(addrPrefix))
}

func (k Keeper) GetAddressPrefix(ctx sdk.Context, denom string) (val string, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AddressPrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	return string(b), true
}