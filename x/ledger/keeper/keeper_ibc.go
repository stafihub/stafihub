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

func (k Keeper) SetIcaPoolDetail(ctx sdk.Context, ica types.IcaPoolDetail) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshal(&ica)

	store.Set(types.IcaPoolDetailStoreKey(ica.Denom, ica.Sequence), b)
}

func (k Keeper) GetIcaPoolDetail(ctx sdk.Context, denom, sequence string) (val types.IcaPoolDetail, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.IcaPoolDetailStoreKey(denom, sequence))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetIcaPoolIndex(ctx sdk.Context, ica types.IcaPoolDetail) {
	store := ctx.KVStore(k.storeKey)

	store.Set(types.IcaPoolIndexStoreKey(ica.Denom, ica.DelegationAccount.Address), []byte(ica.Sequence))
}

func (k Keeper) GetIcaPoolIndex(ctx sdk.Context, denom, delegationAddr string) (val string, found bool) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.IcaPoolIndexStoreKey(denom, delegationAddr))
	if b == nil {
		return "", false
	}

	return string(b), true
}
