package keeper

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/rdex/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper types.BankKeeper
		sudoKeeper types.SudoKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper,
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
		bankKeeper: bankKeeper,
		sudoKeeper: sudoKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetSwapPool(ctx sdk.Context, denom string, swapPool *types.SwapPool) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SwapPoolStoreKey(denom), k.cdc.MustMarshal(swapPool))
}

func (k Keeper) GetSwapPool(ctx sdk.Context, denom string) (*types.SwapPool, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.SwapPoolStoreKey(denom))
	if bts == nil {
		return nil, false
	}

	swapPool := types.SwapPool{}
	k.cdc.MustUnmarshal(bts, &swapPool)
	return &swapPool, true
}

func (k Keeper) GetSwapPoolList(ctx sdk.Context) []*types.SwapPool {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SwapPoolStoreKeyPrefix)
	defer iterator.Close()

	swapPoolList := make([]*types.SwapPool, 0)
	for ; iterator.Valid(); iterator.Next() {
		swapPool := types.SwapPool{}
		k.cdc.MustUnmarshal(iterator.Value(), &swapPool)
		swapPoolList = append(swapPoolList, &swapPool)
	}
	return swapPoolList
}

func (k Keeper) IsRDexLpToken(ctx sdk.Context, denom string) bool {
	_, found := k.GetSwapPool(ctx, denom)
	return found
}

func (k Keeper) AddProvider(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ProviderStoreKey(addr), []byte{})
}

func (k Keeper) RemoveProvider(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ProviderStoreKey(addr))
}

func (k Keeper) HasProvider(ctx sdk.Context, addr sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ProviderStoreKey(addr))
}

func (k Keeper) GetProviderList(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ProviderStoreKeyPrefix)
	defer iterator.Close()

	list := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		if len(key) <= 1 {
			continue
		}

		list = append(list, sdk.AccAddress(key[1:]).String())
	}
	return list
}

func (k Keeper) ToggleProviderSwitch(ctx sdk.Context) {
	k.SetProviderSwitch(ctx, !k.GetProviderSwitch(ctx))
}

func (k Keeper) SetProviderSwitch(ctx sdk.Context, isOpen bool) {
	store := ctx.KVStore(k.storeKey)
	state := types.SwitchStateClose
	if isOpen {
		state = types.SwitchStateOpen
	}
	store.Set(types.ProviderSwitchStoreKey, state)
}

func (k Keeper) GetProviderSwitch(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ProviderSwitchStoreKey)
	if bts == nil {
		return true
	}
	return bytes.Equal(bts, types.SwitchStateOpen)
}

func (k Keeper) AddPoolCreator(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PoolCreatorStoreKey(addr), []byte{})
}

func (k Keeper) RemovePoolCreator(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.PoolCreatorStoreKey(addr))
}

func (k Keeper) HasPoolCreator(ctx sdk.Context, addr sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.PoolCreatorStoreKey(addr))
}

func (k Keeper) GetPoolCreatorList(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PoolCreatorStoreKeyPrefix)
	defer iterator.Close()

	list := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		if len(key) <= 1 {
			continue
		}

		list = append(list, sdk.AccAddress(key[1:]).String())
	}
	return list
}

func (k Keeper) GetSwapPoolNextIndex(ctx sdk.Context) uint32 {
	store := ctx.KVStore(k.storeKey)

	seqBts := store.Get(types.SwapPoolIndexStoreKey)
	if seqBts == nil {
		return 0
	}

	seq := binary.LittleEndian.Uint32(seqBts)
	return seq + 1
}

func (k Keeper) SetSwapPoolIndex(ctx sdk.Context, index uint32) {
	store := ctx.KVStore(k.storeKey)

	seqBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(seqBts, index)
	store.Set(types.SwapPoolIndexStoreKey, seqBts)
}
