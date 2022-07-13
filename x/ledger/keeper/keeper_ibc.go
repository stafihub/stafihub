package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/stafihub/stafihub/x/ledger/types"
)

// ClaimCapability claims the channel capability passed via the OnOpenChanInit callback
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

func (k Keeper) SetICAAccount(ctx sdk.Context, ica types.IcaAccount) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&ica)

	store.Set(types.ICAStoreKey(ica.Owner, ica.CtrlConnectionId), b)
}

func (k Keeper) GetICAAccount(ctx sdk.Context, owner, ctrlConnectionId string) (val types.IcaAccount, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.ICAStoreKey(owner, ctrlConnectionId))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetIcaPoolNextIndex(ctx sdk.Context, denom string) uint32 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IcaPoolNextIndexPrefix)

	key := []byte(denom)

	bts := store.Get(key)
	if bts == nil {
		return 0
	}

	seq := binary.LittleEndian.Uint32(bts)
	return seq + 1
}

func (k Keeper) SetIcaPoolIndex(ctx sdk.Context, denom string, seq uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IcaPoolNextIndexPrefix)

	key := []byte(denom)

	seqBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(seqBts, seq)
	store.Set(key, seqBts)
}

func (k Keeper) SetIcaPoolDetail(ctx sdk.Context, ica *types.IcaPoolDetail) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(ica)

	store.Set(types.IcaPoolDetailStoreKey(ica.Denom, ica.Index), b)
}

func (k Keeper) GetIcaPoolDetail(ctx sdk.Context, denom string, index uint32) (val *types.IcaPoolDetail, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.IcaPoolDetailStoreKey(denom, index))
	if b == nil {
		return val, false
	}
	val = &types.IcaPoolDetail{}
	k.cdc.MustUnmarshal(b, val)
	return val, true
}

func (k Keeper) GetIcaPoolDetailList(ctx sdk.Context, denom string) []*types.IcaPoolDetail {
	denomLen := len(denom)
	key := make([]byte, 1+1+denomLen)
	copy(key[0:], types.IcaPoolDetailPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, key)
	defer iterator.Close()

	list := make([]*types.IcaPoolDetail, 0)
	for ; iterator.Valid(); iterator.Next() {
		detail := types.IcaPoolDetail{}
		k.cdc.MustUnmarshal(iterator.Value(), &detail)
		list = append(list, &detail)
	}
	return list
}

func (k Keeper) SetIcaPoolDelegationAddrIndex(ctx sdk.Context, ica *types.IcaPoolDetail) {
	store := ctx.KVStore(k.storeKey)
	denomLen := len(ica.Denom)
	value := make([]byte, 1+denomLen+4)

	//1+denomLen+4
	value[0] = byte(denomLen)
	copy(value[1:], []byte(ica.Denom))
	binary.LittleEndian.PutUint32(value[1+denomLen:], ica.Index)

	store.Set(types.IcaPoolDelegationAddrIndexStoreKey(ica.DelegationAccount.Address), value)
}

func (k Keeper) GetIcaPoolByDelegationAddr(ctx sdk.Context, delegationAddr string) (val *types.IcaPoolDetail, found bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.IcaPoolDelegationAddrIndexStoreKey(delegationAddr))
	if bts == nil {
		return nil, false
	}
	if len(bts) < 1 {
		return nil, false
	}
	denomLen := int(bts[0])
	if len(bts) != 1+denomLen+4 {
		return nil, false
	}
	denomBts := bts[1 : 1+denomLen]
	index := binary.LittleEndian.Uint32(bts[1+denomLen:])

	return k.GetIcaPoolDetail(ctx, string(denomBts), index)
}

func (k Keeper) SetInterchainTxProposalStatus(ctx sdk.Context, propId string, status uint) {
	store := ctx.KVStore(k.storeKey)

	store.Set(types.InterchainTxPropIdKey(propId), []byte{byte(status)})
}

func (k Keeper) GetInterchainTxProposalStatus(ctx sdk.Context, propId string) (status uint, found bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.InterchainTxPropIdKey(propId))
	if len(bts) == 0 {
		return 0, false
	}
	return uint(bts[0]), true
}

func (k Keeper) SetInterchainTxProposalSequenceIndex(ctx sdk.Context, ctrPortId, ctrChannelId string, sequence uint64, propId string) {
	store := ctx.KVStore(k.storeKey)

	store.Set(types.InterchainTxPropSeqIndexStoreKey(ctrPortId, ctrChannelId, sequence), []byte(propId))
}

func (k Keeper) GetInterchainTxPropIdBySeq(ctx sdk.Context, ctrPortId, ctrChannelId string, sequence uint64) (propId string, found bool) {
	store := ctx.KVStore(k.storeKey)

	bts := store.Get(types.InterchainTxPropSeqIndexStoreKey(ctrPortId, ctrChannelId, sequence))
	if len(bts) == 0 {
		return "", false
	}
	return string(bts), true
}
