package keeper

import (
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
