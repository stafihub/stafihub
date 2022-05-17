package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
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

func (k Keeper) SetStakePool(ctx sdk.Context, stakePool *types.StakePool) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.StakePoolStoreKey(stakePool.StakeTokenDenom), k.cdc.MustMarshal(stakePool))
}

func (k Keeper) GetStakePool(ctx sdk.Context, denom string) (*types.StakePool, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.StakePoolStoreKey(denom))
	if bts == nil {
		return nil, false
	}

	stakePool := types.StakePool{}
	k.cdc.MustUnmarshal(bts, &stakePool)
	return &stakePool, true
}

func (k Keeper) GetStakePoolList(ctx sdk.Context) []*types.StakePool {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.StakePoolStoreKeyPrefix)
	defer iterator.Close()

	stakePoolList := make([]*types.StakePool, 0)
	for ; iterator.Valid(); iterator.Next() {
		stakePool := types.StakePool{}
		k.cdc.MustUnmarshal(iterator.Value(), &stakePool)
		stakePoolList = append(stakePoolList, &stakePool)
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

func (k Keeper) GetUserStakeRecordNextIndex(ctx sdk.Context, userAddress, stakeTokenDenom string) uint32 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UserStakeRecordIndexStoreKeyPrefix)

	key := []byte(userAddress + stakeTokenDenom)

	seqBts := store.Get(key)
	if seqBts == nil {
		return 0
	}

	seq := binary.LittleEndian.Uint32(seqBts)
	return seq + 1
}

func (k Keeper) SetUserStakeRecordIndex(ctx sdk.Context, userAddress, stakeTokenDenom string, index uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UserStakeRecordIndexStoreKeyPrefix)

	key := []byte(userAddress + stakeTokenDenom)

	seqBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(seqBts, index)
	store.Set(key, seqBts)
}

func (k Keeper) GetRewardPoolNextIndex(ctx sdk.Context, stakeTokenDenom string) uint32 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RewardPoolIndexStoreKeyPrefix)

	key := []byte(stakeTokenDenom)

	seqBts := store.Get(key)
	if seqBts == nil {
		return 0
	}

	seq := binary.LittleEndian.Uint32(seqBts)
	return seq + 1
}

func (k Keeper) SetRewardPoolIndex(ctx sdk.Context, stakeTokenDenom string, index uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RewardPoolIndexStoreKeyPrefix)

	key := []byte(stakeTokenDenom)

	seqBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(seqBts, index)
	store.Set(key, seqBts)
}

func (k Keeper) GetStakeItemNextIndex(ctx sdk.Context) uint32 {
	store := ctx.KVStore(k.storeKey)
	seqBts := store.Get(types.StakeItemIndexStoreKey)
	if seqBts == nil {
		return 0
	}

	seq := binary.LittleEndian.Uint32(seqBts)
	return seq + 1
}

func (k Keeper) SetStakeItemIndex(ctx sdk.Context, index uint32) {
	store := ctx.KVStore(k.storeKey)

	seqBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(seqBts, index)
	store.Set(types.StakeItemIndexStoreKey, seqBts)
}

func (k Keeper) SetUserStakeRecord(ctx sdk.Context, userStakeRecord *types.UserStakeRecord) {
	store := ctx.KVStore(k.storeKey)
	store.Set(
		types.UserStakeRecordStoreKey(userStakeRecord.UserAddress, userStakeRecord.StakeTokenDenom, userStakeRecord.Index),
		k.cdc.MustMarshal(userStakeRecord))
}

func (k Keeper) GetUserStakeRecord(ctx sdk.Context, userAddress, stakeTokenDenom string, index uint32) (*types.UserStakeRecord, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.UserStakeRecordStoreKey(userAddress, stakeTokenDenom, index))
	if bts == nil {
		return nil, false
	}

	userStakeRecord := types.UserStakeRecord{}
	k.cdc.MustUnmarshal(bts, &userStakeRecord)
	return &userStakeRecord, true
}

func (k Keeper) GetUserStakeRecordList(ctx sdk.Context, userAddress, stakeTokenDenom string) []*types.UserStakeRecord {
	userAddressLen := len(userAddress)
	stakeTokenDenomLen := len(stakeTokenDenom)

	key := make([]byte, 1+1+userAddressLen+1+stakeTokenDenomLen)
	key[0] = types.UserStakeRecordStoreKeyPrefix[0]
	key[1] = byte(len(userAddress))
	copy(key[2:2+userAddressLen], userAddress)
	key[2+userAddressLen] = byte(stakeTokenDenomLen)
	copy(key[2+userAddressLen+1:2+userAddressLen+1+stakeTokenDenomLen], stakeTokenDenom)

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, key)
	defer iterator.Close()

	userStakeRecordList := make([]*types.UserStakeRecord, 0)
	for ; iterator.Valid(); iterator.Next() {
		userStakeRecordPool := types.UserStakeRecord{}
		k.cdc.MustUnmarshal(iterator.Value(), &userStakeRecordPool)
		userStakeRecordList = append(userStakeRecordList, &userStakeRecordPool)
	}
	return userStakeRecordList
}

func (k Keeper) AddRewarder(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.RewarderStoreKey(addr), []byte{})
}

func (k Keeper) RemoveRewarder(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.RewarderStoreKey(addr))
}

func (k Keeper) HasRewarder(ctx sdk.Context, addr sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.RewarderStoreKey(addr))
}

func (k Keeper) GetRewarderList(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RewarderStoreKeyPrefix)
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
