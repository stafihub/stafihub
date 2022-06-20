package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		sudoKeeper  types.SudoKeeper
		rBankKeeper types.RBankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	sudoKeeper types.SudoKeeper,
	rBankKeeper types.RBankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:         cdc,
		storeKey:    storeKey,
		memKey:      memKey,
		paramstore:  ps,
		sudoKeeper:  sudoKeeper,
		rBankKeeper: rBankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) AddSelectedRValidator(ctx sdk.Context, rValidator *types.RValidator) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SelectedRValdidatorStoreKey(rValidator.Denom, rValidator.PoolAddress, rValidator.ValAddress), []byte{})
}

func (k Keeper) RemoveSelectedRValidator(ctx sdk.Context, rValidator *types.RValidator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.SelectedRValdidatorStoreKey(rValidator.Denom, rValidator.PoolAddress, rValidator.ValAddress))
}

func (k Keeper) HasSelectedRValidator(ctx sdk.Context, rValidator *types.RValidator) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SelectedRValdidatorStoreKey(rValidator.Denom, rValidator.PoolAddress, rValidator.ValAddress))
}

func (k Keeper) GetSelectedRValidatorListByDenomPoolAddress(ctx sdk.Context, denom, poolAddress string) []*types.RValidator {
	store := ctx.KVStore(k.storeKey)
	denomLen := len([]byte(denom))
	poolAddressLen := len([]byte(poolAddress))

	key := make([]byte, 1+1+denomLen+1+poolAddressLen)
	copy(key[0:], types.SelectedRValidatorStoreKeyPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	key[2+denomLen] = byte(poolAddressLen)
	copy(key[2+denomLen+1:], []byte(poolAddress))

	iterator := sdk.KVStorePrefixIterator(store, key)
	defer iterator.Close()

	list := make([]*types.RValidator, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		address := string(key[2+denomLen+1+poolAddressLen+1:])

		rValidator := types.RValidator{
			Denom:       denom,
			PoolAddress: poolAddress,
			ValAddress:  address,
		}

		list = append(list, &rValidator)
	}
	return list
}

func (k Keeper) GetSelectedRValidatorList(ctx sdk.Context) []*types.RValidator {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SelectedRValidatorStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.RValidator, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		denomLen := int(key[1])
		denom := string(key[2 : 2+denomLen])
		poolAddressLen := int(key[2+denomLen])
		poolAddress := string(key[2+denomLen+1 : 2+denomLen+1+poolAddressLen])
		valAddress := string(key[2+denomLen+1+poolAddressLen+1:])

		rValidator := types.RValidator{
			Denom:       denom,
			PoolAddress: poolAddress,
			ValAddress:  valAddress,
		}

		list = append(list, &rValidator)
	}
	return list
}

func (k Keeper) SetLatestVotedCycle(ctx sdk.Context, cycle *types.Cycle) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LatestVotedCycleStoreKey(cycle.Denom, cycle.PoolAddress), k.cdc.MustMarshal(cycle))
}

func (k Keeper) GetLatestVotedCycle(ctx sdk.Context, denom, poolAddress string) *types.Cycle {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.LatestVotedCycleStoreKey(denom, poolAddress))
	if bts == nil {
		return &types.Cycle{
			Denom:   denom,
			Version: 0,
			Number:  0,
		}
	}
	cycle := types.Cycle{}
	k.cdc.MustUnmarshal(bts, &cycle)

	return &cycle
}

func (k Keeper) SetCycleSeconds(ctx sdk.Context, cycleSeconds *types.CycleSeconds) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.CycleSecondsStoreKey(cycleSeconds.Denom), k.cdc.MustMarshal(cycleSeconds))
}

func (k Keeper) GetCycleSeconds(ctx sdk.Context, denom string) *types.CycleSeconds {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.CycleSecondsStoreKey(denom))
	if bts == nil {
		return &types.CycleSeconds{
			Denom:   denom,
			Version: 0,
			Seconds: types.DefaultCycleSeconds,
		}
	}

	cycleSeconds := types.CycleSeconds{}
	k.cdc.MustUnmarshal(bts, &cycleSeconds)
	return &cycleSeconds
}

func (k Keeper) SetShuffleSeconds(ctx sdk.Context, shuffleSeconds *types.ShuffleSeconds) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ShuffleSecondsStoreKey(shuffleSeconds.Denom), k.cdc.MustMarshal(shuffleSeconds))
}

func (k Keeper) GetShuffleSeconds(ctx sdk.Context, denom string) *types.ShuffleSeconds {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ShuffleSecondsStoreKey(denom))
	if bts == nil {
		return &types.ShuffleSeconds{
			Denom:   denom,
			Version: 0,
			Seconds: types.DefaultShuffleSeconds,
		}
	}
	shuffleSeconds := types.ShuffleSeconds{}

	k.cdc.MustUnmarshal(bts, &shuffleSeconds)

	return &shuffleSeconds
}
