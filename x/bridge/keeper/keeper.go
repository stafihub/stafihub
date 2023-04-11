package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/bridge/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper     types.BankKeeper
		sudoKeeper     types.SudoKeeper
		relayersKeeper types.RelayersKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	bankKeeper types.BankKeeper,
	sudoKeeper types.SudoKeeper,
	relayersKeeper types.RelayersKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:            cdc,
		storeKey:       storeKey,
		memKey:         memKey,
		paramstore:     ps,
		bankKeeper:     bankKeeper,
		sudoKeeper:     sudoKeeper,
		relayersKeeper: relayersKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) AddChainId(ctx sdk.Context, chainId uint8) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ChainIdStoreKey(chainId), []byte{})
}

func (k Keeper) HasChainId(ctx sdk.Context, chainId uint8) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ChainIdStoreKey(chainId))
}

func (k Keeper) RmChainId(ctx sdk.Context, chainId uint8) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ChainIdStoreKey(chainId))
}

func (k Keeper) GetChainIdList(ctx sdk.Context) []uint32 {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ChainIdStoreKeyPrefix)
	defer iterator.Close()

	chainIdList := make([]uint32, 0)
	for ; iterator.Valid(); iterator.Next() {
		chainIdList = append(chainIdList, uint32(iterator.Key()[1]))
	}
	return chainIdList
}

func (k Keeper) AddBannedDenom(ctx sdk.Context, chainId uint8, denom string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.BannedDenomStoreKey(chainId, denom), []byte{})
}

func (k Keeper) HasBannedDenom(ctx sdk.Context, chainId uint8, denom string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.BannedDenomStoreKey(chainId, denom))
}

func (k Keeper) RmBannedDenom(ctx sdk.Context, chainId uint8, denom string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.BannedDenomStoreKey(chainId, denom))
}

func (k Keeper) GetBannedDenomList(ctx sdk.Context) []*types.BannedDenom {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.BannedDenomStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.BannedDenom, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		chainId := uint32(key[1])
		denom := string(key[2:])

		list = append(list, &types.BannedDenom{
			ChainId: chainId,
			Denom:   denom,
		})
	}
	return list
}

func (k Keeper) SetRelayFeeReceiver(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.RelayFeeReceiverStoreKey, address)
}

func (k Keeper) GetRelayFeeReceiver(ctx sdk.Context) (sdk.AccAddress, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.RelayFeeReceiverStoreKey)
	if bts == nil {
		return nil, false
	}
	return bts, true
}

func (k Keeper) SetResourceIdToDenom(ctx sdk.Context, rs *types.ResourceIdToDenom) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ResourceIdToDenomStoreKey(rs.ResourceId), k.cdc.MustMarshal(rs))
}

func (k Keeper) GetResourceIdToDenomByResourceId(ctx sdk.Context, resourceId string) (*types.ResourceIdToDenom, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ResourceIdToDenomStoreKey(resourceId))
	if bts == nil {
		return nil, false
	}
	rs := types.ResourceIdToDenom{}
	k.cdc.MustUnmarshal(bts, &rs)

	return &rs, true
}

func (k Keeper) GetResourceIdToDenomByDenom(ctx sdk.Context, denom string) (*types.ResourceIdToDenom, bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ResourceIdToDenomStoreKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		value := iterator.Value()
		if value == nil {
			continue
		}

		rs := types.ResourceIdToDenom{}
		k.cdc.MustUnmarshal(value, &rs)

		if rs.Denom == denom {
			return &rs, true
		}
	}
	return nil, false
}

func (k Keeper) GetResourceIdToDenomList(ctx sdk.Context) []*types.ResourceIdToDenom {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ResourceIdToDenomStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.ResourceIdToDenom, 0)
	for ; iterator.Valid(); iterator.Next() {
		value := iterator.Value()
		if value == nil {
			continue
		}

		rs := types.ResourceIdToDenom{}
		k.cdc.MustUnmarshal(value, &rs)
		list = append(list, &rs)
	}
	return list
}

func (k Keeper) SetDepositCount(ctx sdk.Context, chainId uint8, count uint64) {
	store := ctx.KVStore(k.storeKey)
	bts := sdk.Uint64ToBigEndian(count)
	store.Set(types.DepositCountStoreKey(chainId), bts)
}

func (k Keeper) GetDepositCountById(ctx sdk.Context, chainId uint8) uint64 {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.DepositCountStoreKey(chainId))
	if bts == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bts)
}

func (k Keeper) GetDepositCountList(ctx sdk.Context) []*types.DepositCount {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DepositCountStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.DepositCount, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		chainId := uint32(key[1])
		count := sdk.BigEndianToUint64(iterator.Value())
		list = append(list, &types.DepositCount{
			ChainId: chainId,
			Count:   count,
		})
	}
	return list
}

func (k Keeper) SetProposal(ctx sdk.Context, chainId uint8, depositNonce uint64, resourceId [32]byte, prop *types.Proposal) {
	store := ctx.KVStore(k.storeKey)

	contentBts := k.cdc.MustMarshal(prop.Content)
	hashBts := make([]byte, 0)
	hashBts = append(hashBts, resourceId[:]...)
	hashBts = append(hashBts, contentBts...)
	hash := sha256.Sum256(hashBts)

	propBts := k.cdc.MustMarshal(prop)
	store.Set(types.ProposalStoreKey(chainId, depositNonce, resourceId, hash), propBts)
}

func (k Keeper) GetProposal(ctx sdk.Context, chainId uint8, depositNonce uint64, resourceId [32]byte, content types.ProposalContent) (*types.Proposal, bool) {
	store := ctx.KVStore(k.storeKey)

	contentBts := k.cdc.MustMarshal(&content)
	hashBts := make([]byte, 0)
	hashBts = append(hashBts, resourceId[:]...)
	hashBts = append(hashBts, contentBts...)
	hash := sha256.Sum256(hashBts)
	bts := store.Get(types.ProposalStoreKey(chainId, depositNonce, resourceId, hash))
	if bts == nil {
		return nil, false
	}

	proposal := new(types.Proposal)
	k.cdc.MustUnmarshal(bts, proposal)
	return proposal, true
}

func (k Keeper) GetProposalList(ctx sdk.Context) []*types.GenesisProposal {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ProposalStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.GenesisProposal, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		chainId := uint32(key[1])
		depositNonce := sdk.BigEndianToUint64(key[2:10])
		resourceId := hex.EncodeToString(key[10 : 10+32])

		proposal := new(types.Proposal)
		k.cdc.MustUnmarshal(iterator.Value(), proposal)

		list = append(list, &types.GenesisProposal{
			ChainId:      chainId,
			DepositNonce: depositNonce,
			ResourceId:   resourceId,
			Proposal:     proposal,
		})

	}
	return list
}

func (k Keeper) SetRelayFee(ctx sdk.Context, chainId uint8, value sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&value)
	store.Set(types.RelayFeeStoreKey(chainId), b)
}

func (k Keeper) GetRelayFee(ctx sdk.Context, chainId uint8) (value sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.RelayFeeStoreKey(chainId))
	if b == nil {
		return sdk.NewCoin(utils.FisDenom, sdk.ZeroInt())
	}
	k.cdc.MustUnmarshal(b, &value)
	return value
}

func (k Keeper) GetRelayFeeList(ctx sdk.Context) []*types.RelayFee {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RelayFeeStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.RelayFee, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		chainId := uint32(key[1])

		value := sdk.Coin{}
		k.cdc.MustUnmarshal(iterator.Value(), &value)

		list = append(list, &types.RelayFee{
			ChainId: chainId,
			Value:   value,
		})
	}
	return list
}
