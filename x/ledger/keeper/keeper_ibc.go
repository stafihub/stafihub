package keeper

import (
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
