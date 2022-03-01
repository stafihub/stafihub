package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/x/bridge/types"
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

func (k Keeper) AddRelayer(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.RelayStoreKey(address), []byte{})
}

func (k Keeper) HasRelayer(ctx sdk.Context, address sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.RelayStoreKey(address))
}

func (k Keeper) GetRelayers(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RelayerStoreKeyPrefix)
	defer iterator.Close()

	relayerList := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		if len(iterator.Key()) < 1 {
			continue
		}
		relayerList = append(relayerList, sdk.AccAddress(iterator.Key()[1:]).String())
	}
	return relayerList
}

func (k Keeper) AddChainId(ctx sdk.Context, chainId uint8) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ChainIdStoreKey(chainId), []byte{})
}

func (k Keeper) HasChainId(ctx sdk.Context, chainId uint8) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ChainIdStoreKey(chainId))
}

func (k Keeper) GetAllChainId(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ChainIdStoreKeyPrefix)
	defer iterator.Close()

	chainIdList := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		if len(iterator.Key()) != 2 {
			continue
		}
		chainIdList = append(chainIdList, fmt.Sprintf("%d", iterator.Key()[1]))
	}
	return chainIdList
}

func (k Keeper) SetThreshold(ctx sdk.Context, threshold uint8) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ThresholdStoreKey, []byte{threshold})
}

func (k Keeper) GetThreshold(ctx sdk.Context) (uint8, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ThresholdStoreKey)
	if len(bts) == 0 {
		return 0, false
	}
	return bts[0], true
}

func (k Keeper) SetResourceIdToDenom(ctx sdk.Context, resourceId [32]byte, denom string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ResourceIdToDenomStoreKey(resourceId), []byte(denom))
}

func (k Keeper) GetDenomByResourceId(ctx sdk.Context, resourceId [32]byte) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ResourceIdToDenomStoreKey(resourceId))
	if len(bts) == 0 {
		return "", false
	}
	return string(bts), true
}

func (k Keeper) GetAllResourceIdToDenom(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ResourceIdToDenomStoreKeyPrefix)
	defer iterator.Close()

	chainIdList := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		if len(iterator.Key()) < 1 {
			continue
		}

		chainIdList = append(chainIdList, hex.EncodeToString(iterator.Key()[1:33])+":"+string(iterator.Value()))
	}
	return chainIdList
}

func (k Keeper) SetDepositCounts(ctx sdk.Context, chainId uint8, count uint64) {
	store := ctx.KVStore(k.storeKey)
	bts := sdk.Uint64ToBigEndian(count)
	store.Set(types.DepositCountsStoreKey(chainId), bts)
}

func (k Keeper) GetDepositCounts(ctx sdk.Context, chainId uint8) uint64 {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.DepositCountsStoreKey(chainId))
	if len(bts) == 0 {
		return 0
	}
	return sdk.BigEndianToUint64(bts)
}

func (k Keeper) SetProposal(ctx sdk.Context, chainId uint8, depositNonce uint64, resourceId [32]byte, prop *types.Proposal) {
	store := ctx.KVStore(k.storeKey)

	contentBts, err := prop.Content.Marshal()
	if err != nil {
		panic(err)
	}
	hashBts := make([]byte, 0)
	hashBts = append(hashBts, resourceId[:]...)
	hashBts = append(hashBts, contentBts...)
	hash := sha256.Sum256(hashBts)

	propBts, err := prop.Marshal()
	if err != nil {
		panic(err)
	}
	store.Set(types.ProposalStoreKey(chainId, depositNonce, hash), propBts)
}

func (k Keeper) GetProposal(ctx sdk.Context, chainId uint8, depositNonce uint64, resourceId [32]byte, content types.ProposalContent) (*types.Proposal, bool) {
	store := ctx.KVStore(k.storeKey)

	contentBts, err := content.Marshal()
	if err != nil {
		panic(err)
	}
	hashBts := make([]byte, 0)
	hashBts = append(hashBts, resourceId[:]...)
	hashBts = append(hashBts, contentBts...)
	hash := sha256.Sum256(hashBts)
	bts := store.Get(types.ProposalStoreKey(chainId, depositNonce, hash))
	if bts == nil {
		return nil, false
	}

	proposal := new(types.Proposal)
	err = proposal.Unmarshal(bts)
	if err != nil {
		panic(err)
	}
	return proposal, true
}

func (k Keeper) SetResourceIdType(ctx sdk.Context, resourceId [32]byte, idType types.ResourceIdType) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ResourceIdTypeStoreKey(resourceId), idType[:])
}

func (k Keeper) GetResourceIdType(ctx sdk.Context, resourceId [32]byte) types.ResourceIdType {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ResourceIdTypeStoreKey(resourceId))
	if len(bts) == 0 {
		return types.ResourceIdTypeForeign
	}
	var idType types.ResourceIdType
	copy(idType[:], bts)
	return idType
}

func (k Keeper) GetAllResourceTypes(ctx sdk.Context) []string {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ResourceIdTypeStoreKeyPrefix)
	defer iterator.Close()

	chainIdList := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		if len(iterator.Key()) < 1 {
			continue
		}
		value := iterator.Value()
		if len(value) == 0 {
			value = types.ResourceIdTypeForeign[:]
		}

		chainIdList = append(chainIdList, hex.EncodeToString(iterator.Key()[1:])+":"+fmt.Sprintf("%d", value[0]))
	}
	return chainIdList
}
