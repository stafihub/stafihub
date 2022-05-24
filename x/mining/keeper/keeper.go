package keeper

import (
	"bytes"
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

func (k Keeper) SetStakePool(ctx sdk.Context, stakePool *types.StakePool) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.StakePoolStoreKey(stakePool.Index), k.cdc.MustMarshal(stakePool))
}

func (k Keeper) GetStakePool(ctx sdk.Context, stakePoolIndex uint32) (*types.StakePool, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.StakePoolStoreKey(stakePoolIndex))
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
	store.Set(types.StakeItemStoreKey(stakeItem.StakePoolIndex, stakeItem.Index), k.cdc.MustMarshal(stakeItem))
}

func (k Keeper) GetStakeItem(ctx sdk.Context, stakePoolIndex, index uint32) (*types.StakeItem, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.StakeItemStoreKey(stakePoolIndex, index))
	if bts == nil {
		return nil, false
	}

	stakeItem := types.StakeItem{}
	k.cdc.MustUnmarshal(bts, &stakeItem)
	return &stakeItem, true
}

func (k Keeper) GetStakeItemList(ctx sdk.Context, stakePoolIndex uint32) []*types.StakeItem {
	store := ctx.KVStore(k.storeKey)
	bts := make([]byte, 4)
	binary.LittleEndian.PutUint32(bts, stakePoolIndex)

	iterator := sdk.KVStorePrefixIterator(store, append(types.StakeItemStoreKeyPrefix, bts...))
	defer iterator.Close()

	stakeItemList := make([]*types.StakeItem, 0)
	for ; iterator.Valid(); iterator.Next() {
		stakeItem := types.StakeItem{}
		k.cdc.MustUnmarshal(iterator.Value(), &stakeItem)
		stakeItemList = append(stakeItemList, &stakeItem)
	}
	return stakeItemList
}

func (k Keeper) GetUserStakeRecordNextIndex(ctx sdk.Context, userAddress string, stakePoolIndex uint32) uint32 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UserStakeRecordIndexStoreKeyPrefix)

	indexBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(indexBts, stakePoolIndex)

	key := append([]byte(userAddress), indexBts...)

	seqBts := store.Get(key)
	if seqBts == nil {
		return 0
	}

	seq := binary.LittleEndian.Uint32(seqBts)
	return seq + 1
}

func (k Keeper) SetUserStakeRecordIndex(ctx sdk.Context, userAddress string, stakePoolIndex, index uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.UserStakeRecordIndexStoreKeyPrefix)

	indexBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(indexBts, stakePoolIndex)

	key := append([]byte(userAddress), indexBts...)

	seqBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(seqBts, index)
	store.Set(key, seqBts)
}

func (k Keeper) GetRewardPoolNextIndex(ctx sdk.Context, stakePoolIndex uint32) uint32 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RewardPoolIndexStoreKeyPrefix)
	indexBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(indexBts, stakePoolIndex)

	seqBts := store.Get(indexBts)
	if seqBts == nil {
		return 0
	}

	seq := binary.LittleEndian.Uint32(seqBts)
	return seq + 1
}

func (k Keeper) SetRewardPoolIndex(ctx sdk.Context, stakePoolIndex, index uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.RewardPoolIndexStoreKeyPrefix)
	indexBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(indexBts, stakePoolIndex)

	seqBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(seqBts, index)
	store.Set(indexBts, seqBts)
}

func (k Keeper) GetStakePoolNextIndex(ctx sdk.Context) uint32 {
	store := ctx.KVStore(k.storeKey)

	seqBts := store.Get(types.StakePoolIndexStoreKey)
	if seqBts == nil {
		return 0
	}

	seq := binary.LittleEndian.Uint32(seqBts)
	return seq + 1
}

func (k Keeper) SetStakePoolIndex(ctx sdk.Context, index uint32) {
	store := ctx.KVStore(k.storeKey)

	seqBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(seqBts, index)
	store.Set(types.StakePoolIndexStoreKey, seqBts)
}

func (k Keeper) GetStakeItemNextIndex(ctx sdk.Context, stakePoolIndex uint32) uint32 {
	store := ctx.KVStore(k.storeKey)
	seqBts := store.Get(types.StakeItemIndexStoreKey(stakePoolIndex))
	if seqBts == nil {
		return 0
	}

	seq := binary.LittleEndian.Uint32(seqBts)
	return seq + 1
}

func (k Keeper) SetStakeItemIndex(ctx sdk.Context, stakePoolIndex, index uint32) {
	store := ctx.KVStore(k.storeKey)

	seqBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(seqBts, index)

	store.Set(types.StakeItemIndexStoreKey(stakePoolIndex), seqBts)
}

func (k Keeper) SetUserStakeRecord(ctx sdk.Context, userStakeRecord *types.UserStakeRecord) {
	store := ctx.KVStore(k.storeKey)
	store.Set(
		types.UserStakeRecordStoreKey(userStakeRecord.UserAddress, userStakeRecord.StakePoolIndex, userStakeRecord.Index),
		k.cdc.MustMarshal(userStakeRecord))
}

func (k Keeper) GetUserStakeRecord(ctx sdk.Context, userAddress string, stakePoolIndex, index uint32) (*types.UserStakeRecord, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.UserStakeRecordStoreKey(userAddress, stakePoolIndex, index))
	if bts == nil {
		return nil, false
	}

	userStakeRecord := types.UserStakeRecord{}
	k.cdc.MustUnmarshal(bts, &userStakeRecord)
	return &userStakeRecord, true
}

// prefix + len(userAddress) + userAddress + stakePoolIndex + index
func (k Keeper) GetUserStakeRecordList(ctx sdk.Context, userAddress string, stakePoolIndex uint32) []*types.UserStakeRecord {
	userAddressLen := len(userAddress)
	key := make([]byte, 1+1+userAddressLen+4)

	key[0] = types.UserStakeRecordStoreKeyPrefix[0]
	key[1] = byte(len(userAddress))
	copy(key[2:2+userAddressLen], userAddress)
	binary.LittleEndian.PutUint32(key[2+userAddressLen:], stakePoolIndex)

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

func (k Keeper) AddMiningProvider(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.MiningProviderStoreKey(addr), []byte{})
}

func (k Keeper) RemoveMiningProvider(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.MiningProviderStoreKey(addr))
}

func (k Keeper) HasMiningProvider(ctx sdk.Context, addr sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.MiningProviderStoreKey(addr))
}

func (k Keeper) GetMiningProviderList(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.MiningProviderStoreKeyPrefix)
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

func (k Keeper) AddRewardToken(ctx sdk.Context, rewardToken *types.RewardToken) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.RewardTokenStoreKey(rewardToken.RewardTokenDenom), k.cdc.MustMarshal(rewardToken))
}

func (k Keeper) RemoveRewardToken(ctx sdk.Context, denom string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.RewardTokenStoreKey(denom))
}

func (k Keeper) GetRewardToken(ctx sdk.Context, denom string) (*types.RewardToken, bool) {
	store := ctx.KVStore(k.storeKey)

	bts := store.Get(types.RewardTokenStoreKey(denom))
	if bts == nil {
		return nil, false
	}

	rewardToken := types.RewardToken{}
	k.cdc.Unmarshal(bts, &rewardToken)

	return &rewardToken, true
}

func (k Keeper) HasRewardToken(ctx sdk.Context, denom string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.RewardTokenStoreKey(denom))
}

func (k Keeper) GetRewardTokenList(ctx sdk.Context) []*types.RewardToken {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RewardTokenStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.RewardToken, 0)
	for ; iterator.Valid(); iterator.Next() {
		rewardToken := types.RewardToken{}

		k.cdc.Unmarshal(iterator.Value(), &rewardToken)
		list = append(list, &rewardToken)
	}
	return list
}

func (k Keeper) SetMaxRewardPoolNumber(ctx sdk.Context, number uint32) {
	store := ctx.KVStore(k.storeKey)
	bts := make([]byte, 4)
	binary.LittleEndian.PutUint32(bts, number)

	store.Set(types.MaxRewardPoolNumberStoreKey, bts)
}

func (k Keeper) GetMaxRewardPoolNumber(ctx sdk.Context) uint32 {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.MaxRewardPoolNumberStoreKey)
	if b == nil {
		return 32
	}
	return binary.LittleEndian.Uint32(b)
}

func (k Keeper) ToggleMiningProviderSwitch(ctx sdk.Context) {
	k.SetMiningProviderSwitch(ctx, !k.GetMiningProviderSwitch(ctx))
}

func (k Keeper) SetMiningProviderSwitch(ctx sdk.Context, isOpen bool) {
	store := ctx.KVStore(k.storeKey)
	state := types.SwitchStateClose
	if isOpen {
		state = types.SwitchStateOpen
	}
	store.Set(types.MiningProviderSwitchStoreKey, state)
}

func (k Keeper) GetMiningProviderSwitch(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.MiningProviderSwitchStoreKey)
	if bts == nil {
		return true
	}
	return bytes.Equal(bts, types.SwitchStateOpen)
}
