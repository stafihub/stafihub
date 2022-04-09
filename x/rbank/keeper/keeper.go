package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/rbank/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
		sudoKeeper types.SudoKeeper
		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	sudoKeeper types.SudoKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		sudoKeeper: sudoKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
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

func (k Keeper) CheckAddress(ctx sdk.Context, denom, address string) error {
	prefix, found := k.GetAddressPrefix(ctx, denom)
	if !found {
		return types.ErrAddrPrefixNotExist
	}
	_, err := sdk.GetFromBech32(address, prefix)
	if err != nil {
		return err
	}
	return nil
}
