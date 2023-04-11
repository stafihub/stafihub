package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/rmintreward/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		sudoKeeper types.SudoKeeper
		bankKeper  types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
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
		bankKeper:  bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetActLatestCycle(ctx sdk.Context, denom string, cycle uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ActLatestCycleStoreKey(denom), sdk.Uint64ToBigEndian(cycle))
}

func (k Keeper) GetActLatestCycle(ctx sdk.Context, denom string) (uint64, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ActLatestCycleStoreKey(denom))
	if bts == nil {
		return 0, false
	}
	return sdk.BigEndianToUint64(bts), true
}

func (k Keeper) GetActLatestCycleList(ctx sdk.Context) []*types.ActLatestCycle {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ActLatestCycleStoreKeyPrefix)
	defer iterator.Close()

	latestCycleList := make([]*types.ActLatestCycle, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		denom := string(key[1:])
		cycle := sdk.BigEndianToUint64(iterator.Value())

		latestCycleList = append(latestCycleList, &types.ActLatestCycle{
			Denom: denom,
			Cycle: cycle,
		})
	}
	return latestCycleList
}

func (k Keeper) SetActCurrentCycle(ctx sdk.Context, denom string, cycle uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ActCurrentCycleStoreKey(denom), sdk.Uint64ToBigEndian(cycle))
}

func (k Keeper) GetActCurrentCycle(ctx sdk.Context, denom string) (uint64, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ActCurrentCycleStoreKey(denom))
	if bts == nil {
		return 0, false
	}
	return sdk.BigEndianToUint64(bts), true
}

func (k Keeper) GetActCurrentCycleList(ctx sdk.Context) []*types.ActCurrentCycle {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ActCurrentCycleStoreKeyPrefix)
	defer iterator.Close()

	currentCycleList := make([]*types.ActCurrentCycle, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		denom := string(key[1:])
		cycle := sdk.BigEndianToUint64(iterator.Value())

		currentCycleList = append(currentCycleList, &types.ActCurrentCycle{
			Denom: denom,
			Cycle: cycle,
		})
	}
	return currentCycleList
}

func (k Keeper) SetMintRewardAct(ctx sdk.Context, denom string, cycle uint64, act *types.MintRewardAct) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.MintRewardActStoreKey(denom, cycle), k.cdc.MustMarshal(act))
}

func (k Keeper) GetMintRewardAct(ctx sdk.Context, denom string, cycle uint64) (*types.MintRewardAct, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.MintRewardActStoreKey(denom, cycle))
	if bts == nil {
		return nil, false
	}
	act := types.MintRewardAct{}
	k.cdc.MustUnmarshal(bts, &act)
	return &act, true
}

func (k Keeper) GetMintRewardActList(ctx sdk.Context) []*types.GenesisMintRewardAct {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.MintRewardActStoreKeyPrefix)
	defer iterator.Close()

	mintRewardActList := make([]*types.GenesisMintRewardAct, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		denom := string(key[1 : len(key)-8])
		cycle := sdk.BigEndianToUint64(key[len(key)-8:])

		mintRewardAct := types.MintRewardAct{}
		k.cdc.MustUnmarshal(iterator.Value(), &mintRewardAct)

		mintRewardActList = append(mintRewardActList, &types.GenesisMintRewardAct{
			Denom:         denom,
			Cycle:         cycle,
			MintRewardAct: &mintRewardAct,
		})
	}
	return mintRewardActList
}

func (k Keeper) SetUserClaimInfo(ctx sdk.Context, account sdk.AccAddress, denom string, cycle uint64, mintIndex uint64, claimInfo *types.UserClaimInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.UserClaimInfoStoreKey(account, denom, cycle, mintIndex), k.cdc.MustMarshal(claimInfo))
}

func (k Keeper) GetUserClaimInfo(ctx sdk.Context, account sdk.AccAddress, denom string, cycle uint64, mintIndex uint64) (*types.UserClaimInfo, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.UserClaimInfoStoreKey(account, denom, cycle, mintIndex))
	if bts == nil {
		return nil, false
	}
	claimInfo := types.UserClaimInfo{}
	k.cdc.MustUnmarshal(bts, &claimInfo)
	return &claimInfo, true
}

func (k Keeper) GetUserClaimInfoList(ctx sdk.Context) []*types.GenesisUserClaimInfo {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UserClaimInfoStoreKeyPrefix)
	defer iterator.Close()

	userClaimInfoList := make([]*types.GenesisUserClaimInfo, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		accountLen := int(key[1])
		account := sdk.AccAddress(key[2 : 2+accountLen])
		denomLen := int(key[2+accountLen])
		denom := string(key[2+accountLen+1 : 2+accountLen+1+denomLen])
		cycle := sdk.BigEndianToUint64(key[2+accountLen+1+denomLen : 2+accountLen+1+denomLen+8])
		mintIndex := sdk.BigEndianToUint64(key[2+accountLen+1+denomLen+8:])

		userClaimInfo := types.UserClaimInfo{}
		k.cdc.MustUnmarshal(iterator.Value(), &userClaimInfo)

		userClaimInfoList = append(userClaimInfoList, &types.GenesisUserClaimInfo{
			Account:       account.String(),
			Denom:         denom,
			Cycle:         cycle,
			MintIndex:     mintIndex,
			UserClaimInfo: &userClaimInfo,
		})
	}
	return userClaimInfoList
}

func (k Keeper) SetUserActs(ctx sdk.Context, account sdk.AccAddress, denom string, actCycles *types.Acts) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.UserActsStoreKey(account, denom), k.cdc.MustMarshal(actCycles))
}

func (k Keeper) GetUserActs(ctx sdk.Context, account sdk.AccAddress, denom string) (*types.Acts, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.UserActsStoreKey(account, denom))
	if bts == nil {
		return nil, false
	}
	actCycles := types.Acts{}
	k.cdc.MustUnmarshal(bts, &actCycles)
	return &actCycles, true
}

func (k Keeper) GetUserActsList(ctx sdk.Context) []*types.GenesisUserAct {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UserActsStoreKeyPrefix)
	defer iterator.Close()

	userActList := make([]*types.GenesisUserAct, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		accountLen := int(key[1])
		account := sdk.AccAddress(key[2 : 2+accountLen])
		denom := string(key[2+accountLen:])

		acts := types.Acts{}
		k.cdc.MustUnmarshal(iterator.Value(), &acts)

		userActList = append(userActList, &types.GenesisUserAct{
			Account: account.String(),
			Denom:   denom,
			Acts:    &acts,
		})
	}
	return userActList
}

func (k Keeper) SetUserMintCount(ctx sdk.Context, account sdk.AccAddress, denom string, cycle, count uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.UserMintCountStoreKey(account, denom, cycle), sdk.Uint64ToBigEndian(count))
}

func (k Keeper) GetUserMintCount(ctx sdk.Context, account sdk.AccAddress, denom string, cycle uint64) (uint64, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.UserMintCountStoreKey(account, denom, cycle))
	if bts == nil {
		return 0, false
	}
	return sdk.BigEndianToUint64(bts), true
}

func (k Keeper) GetUserMintCountList(ctx sdk.Context) []*types.UserMintCount {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.UserMintCountStoreKeyPrefix)
	defer iterator.Close()

	userMintCountList := make([]*types.UserMintCount, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		accountLen := int(key[1])
		account := sdk.AccAddress(key[2 : 2+accountLen])
		denomLen := int(key[2+accountLen])
		denom := string(key[2+accountLen+1 : 2+accountLen+1+denomLen])
		cycle := sdk.BigEndianToUint64(key[2+accountLen+1+denomLen:])
		count := sdk.BigEndianToUint64(iterator.Value())

		userMintCountList = append(userMintCountList, &types.UserMintCount{
			Account: account.String(),
			Denom:   denom,
			Cycle:   cycle,
			Count:   count,
		})
	}
	return userMintCountList
}
