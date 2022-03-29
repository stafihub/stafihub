package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/rmintreward/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		sudoKeeper types.SudoKeeper
		bankKeper  types.BankKeeper
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

func (k Keeper) SetActCurrentCycle(ctx sdk.Context, denom string, cycle uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ActCurrentCycleStoreKey(denom), sdk.Uint64ToBigEndian(cycle))
}

func (k Keeper) GetActCurrenttCycle(ctx sdk.Context, denom string) (uint64, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ActCurrentCycleStoreKey(denom))
	if bts == nil {
		return 0, false
	}
	return sdk.BigEndianToUint64(bts), true
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

func (k Keeper) SetUserClaimInfo(ctx sdk.Context, account sdk.AccAddress, denom string, cycle uint64, mintIndex uint64, claimInfo *types.UserClaimInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.UserClaimInforStoreKey(account, denom, cycle, mintIndex), k.cdc.MustMarshal(claimInfo))
}

func (k Keeper) GetUserClaimInfo(ctx sdk.Context, account sdk.AccAddress, denom string, cycle uint64, mintIndex uint64) (*types.UserClaimInfo, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.UserClaimInforStoreKey(account, denom, cycle, mintIndex))
	if bts == nil {
		return nil, false
	}
	claimInfo := types.UserClaimInfo{}
	k.cdc.MustUnmarshal(bts, &claimInfo)
	return &claimInfo, true
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

func (k Keeper) AddActDenom(ctx sdk.Context, denom string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ActDenomsStoreKey(denom), []byte{})
}

func (k Keeper) GetActDenoms(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ActDenomsStoreKeyPrefix)
	defer iterator.Close()

	denoms := []string{}
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		denoms = append(denoms, string(key[len(types.ActDenomsStoreKeyPrefix):]))
	}
	return denoms
}
