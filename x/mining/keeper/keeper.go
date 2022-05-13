package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/mining/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

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
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetStakePool(ctx sdk.Context, denom string, swapPool *types.StakePool) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.StakePoolStoreKey(denom), k.cdc.MustMarshal(swapPool))
}

func (k Keeper) GetStakePool(ctx sdk.Context, denom string) (*types.StakePool, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.StakePoolStoreKey(denom))
	if bts == nil {
		return nil, false
	}

	swapPool := types.StakePool{}
	k.cdc.MustUnmarshal(bts, &swapPool)
	return &swapPool, true
}

func (k Keeper) GetStakePoolList(ctx sdk.Context) []*types.StakePool {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.StakePoolStoreKeyPrefix)
	defer iterator.Close()

	stakePoolList := make([]*types.StakePool, 0)
	for ; iterator.Valid(); iterator.Next() {
		swapPool := types.StakePool{}
		k.cdc.MustUnmarshal(iterator.Value(), &swapPool)
		stakePoolList = append(stakePoolList, &swapPool)
	}
	return stakePoolList
}

func (k Keeper) SetStakeItem(ctx sdk.Context, stakeItem *types.StakeItem) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.StakeItemStoreKey(stakeItem.Index), k.cdc.MustMarshal(stakeItem))
}

func (k Keeper) GetStakeItem(ctx sdk.Context, index uint32) (*types.StakeItem, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.StakeItemStoreKey(index))
	if bts == nil {
		return nil, false
	}

	stakeItem := types.StakeItem{}
	k.cdc.MustUnmarshal(bts, &stakeItem)
	return &stakeItem, true
}

func (k Keeper) GetStakeItemList(ctx sdk.Context) []*types.StakeItem {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.StakeItemStoreKeyPrefix)
	defer iterator.Close()

	stakeItemList := make([]*types.StakeItem, 0)
	for ; iterator.Valid(); iterator.Next() {
		stakeItem := types.StakeItem{}
		k.cdc.MustUnmarshal(iterator.Value(), &stakeItem)
		stakeItemList = append(stakeItemList, &stakeItem)
	}
	return stakeItemList
}
