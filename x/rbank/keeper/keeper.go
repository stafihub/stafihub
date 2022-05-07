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
	bankKeeper types.BankKeeper,
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
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetAccAddressPrefix(ctx sdk.Context, denom, addrPrefix string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccAddressPrefix)
	store.Set([]byte(denom), []byte(addrPrefix))
}

func (k Keeper) GetAccAddressPrefix(ctx sdk.Context, denom string) (val string, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.AccAddressPrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}
	return string(b), true
}

func (k Keeper) GetAddressPrefixList(ctx sdk.Context) []*types.AddressPrefix {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AccAddressPrefix)
	defer iterator.Close()

	list := make([]*types.AddressPrefix, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		denom := string(key[1:])
		accPrefix := string(iterator.Value())
		valPrefix, found := k.GetValAddressPrefix(ctx, denom)
		if !found {
			continue
		}

		list = append(list, &types.AddressPrefix{
			Denom:            denom,
			AccAddressPrefix: accPrefix,
			ValAddressPrefix: valPrefix,
		})
	}

	return list
}

func (k Keeper) CheckAccAddress(ctx sdk.Context, denom, address string) error {
	prefix, found := k.GetAccAddressPrefix(ctx, denom)
	if !found {
		return types.ErrAccAddrPrefixNotExist
	}
	_, err := sdk.GetFromBech32(address, prefix)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) SetValAddressPrefix(ctx sdk.Context, denom, addrPrefix string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValAddressPrefix)
	store.Set([]byte(denom), []byte(addrPrefix))
}

func (k Keeper) GetValAddressPrefix(ctx sdk.Context, denom string) (val string, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ValAddressPrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}
	return string(b), true
}

func (k Keeper) CheckValAddress(ctx sdk.Context, denom, address string) error {
	prefix, found := k.GetValAddressPrefix(ctx, denom)
	if !found {
		return types.ErrValAddrPrefixNotExist
	}
	_, err := sdk.GetFromBech32(address, prefix)
	if err != nil {
		return err
	}
	return nil
}
