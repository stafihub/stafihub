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

func (k Keeper) GetIcapPoolNextSequence(ctx sdk.Context, denom string) uint32 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IcaPoolNextSequencePrefix)

	key := []byte(denom)

	seqBts := store.Get(key)
	if seqBts == nil {
		return 0
	}

	seq := binary.LittleEndian.Uint32(seqBts)
	return seq + 1
}

func (k Keeper) SetIcaPoolSequence(ctx sdk.Context, denom string, seq uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.IcaPoolNextSequencePrefix)

	key := []byte(denom)

	seqBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(seqBts, seq)
	store.Set(key, seqBts)
}

func (k Keeper) SetIcaPoolDetail(ctx sdk.Context, ica *types.IcaPoolDetail) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(ica)

	store.Set(types.IcaPoolDetailStoreKey(ica.Denom, ica.Sequence), b)
}

func (k Keeper) GetIcaPoolDetail(ctx sdk.Context, denom, sequence string) (val *types.IcaPoolDetail, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.IcaPoolDetailStoreKey(denom, sequence))
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

func (k Keeper) SetIcaPoolIndex(ctx sdk.Context, ica *types.IcaPoolDetail) {
	store := ctx.KVStore(k.storeKey)
	denomLen := len(ica.Denom)
	value := make([]byte, 1+denomLen+len(ica.Sequence))

	//1+denomLen+sequenceLen
	value[0] = byte(denomLen)
	copy(value[1:], []byte(ica.Denom))
	copy(value[1+denomLen:], []byte(ica.Sequence))

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
	if len(bts) < 1+denomLen {
		return nil, false
	}
	denom := bts[1 : 1+denomLen]
	sequence := bts[1+denomLen:]

	return k.GetIcaPoolDetail(ctx, string(denom), string(sequence))
}
