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
