package keeper

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/claim/types"
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

func (k Keeper) SetMerkleRoot(ctx sdk.Context, round uint64, root NodeHash) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.MerkleRootStoreKey(round), root[:])
}

func (k Keeper) GetMerkleRoot(ctx sdk.Context, round uint64) (NodeHash, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.MerkleRootStoreKey(round))
	if bts == nil {
		return nil, false
	}

	return bts, true
}

func (k Keeper) GetMerkleRootList(ctx sdk.Context) []*types.MerkleRoot {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.MerkleRootStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.MerkleRoot, 0)
	for ; iterator.Valid(); iterator.Next() {

		key := iterator.Key()
		round := sdk.BigEndianToUint64(key[1:])
		merkleRoot := types.MerkleRoot{
			Round:    round,
			RootHash: hex.EncodeToString(iterator.Value()),
		}

		list = append(list, &merkleRoot)
	}
	return list
}

func (k Keeper) SetClaimBitMap(ctx sdk.Context, claimRound, wordIndex, bits uint64) {
	store := ctx.KVStore(k.storeKey)

	bts := sdk.Uint64ToBigEndian(bits)

	store.Set(types.ClaimBitMapStoreKey(claimRound, wordIndex), bts)
}

func (k Keeper) GetClaimBitMap(ctx sdk.Context, claimRound, wordIndex uint64) uint64 {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ClaimBitMapStoreKey(claimRound, wordIndex))
	if bts == nil {
		return 0
	}

	return sdk.BigEndianToUint64(bts)
}

func (k Keeper) GetClaimBitMapList(ctx sdk.Context) []*types.ClaimBitMap {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ClaimBitMapStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.ClaimBitMap, 0)
	for ; iterator.Valid(); iterator.Next() {

		key := iterator.Key()
		round := sdk.BigEndianToUint64(key[1:9])
		wordIndex := sdk.BigEndianToUint64(key[9:])
		bits := sdk.BigEndianToUint64(iterator.Value())

		claimBitMap := types.ClaimBitMap{
			Round:     round,
			WordIndex: wordIndex,
			Bits:      bits,
		}
		list = append(list, &claimBitMap)
	}
	return list
}

func (k Keeper) IsIndexClaimed(ctx sdk.Context, claimRound, index uint64) bool {
	claimedWordIndex := index / 64
	claimedBitIndex := index % 64

	mask := uint64(1 << claimedBitIndex)
	bits := k.GetClaimBitMap(ctx, claimRound, claimedWordIndex)

	return (bits & mask) == mask
}

func (k Keeper) SetIndexClaimed(ctx sdk.Context, claimRound, index uint64) {
	claimedWordIndex := index / 64
	claimedBitIndex := index % 64

	bits := k.GetClaimBitMap(ctx, claimRound, claimedWordIndex)
	newBits := bits | (1 << claimedBitIndex)

	k.SetClaimBitMap(ctx, claimRound, claimedWordIndex, newBits)
}

func (k Keeper) ToggleClaimSwitch(ctx sdk.Context, round uint64) {
	k.SetClaimSwitch(ctx, round, !k.GetClaimSwitch(ctx, round))
}

func (k Keeper) SetClaimSwitch(ctx sdk.Context, round uint64, isOpen bool) {
	store := ctx.KVStore(k.storeKey)
	state := types.SwitchStateClose
	if isOpen {
		state = types.SwitchStateOpen
	}
	store.Set(types.ClaimSwitchStoreKey(round), state)
}

func (k Keeper) GetClaimSwitch(ctx sdk.Context, round uint64) bool {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ClaimSwitchStoreKey(round))
	if bts == nil {
		return true
	}
	return bytes.Equal(bts, types.SwitchStateOpen)
}

func (k Keeper) GetClaimSwitchList(ctx sdk.Context) []*types.ClaimSwitch {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ClaimSwitchStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.ClaimSwitch, 0)
	for ; iterator.Valid(); iterator.Next() {

		key := iterator.Key()
		round := sdk.BigEndianToUint64(key[1:9])

		isOpen := bytes.Equal(types.SwitchStateOpen, iterator.Value())

		claimSwitch := types.ClaimSwitch{
			Round:  round,
			IsOpen: isOpen,
		}
		list = append(list, &claimSwitch)
	}
	return list
}
