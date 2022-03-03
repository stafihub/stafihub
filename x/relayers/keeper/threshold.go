package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/relayers/types"
)

// SetThreshold set a specific threshold in the store from its denom
func (k Keeper) SetThreshold(ctx sdk.Context, taipe, denom string, value uint32) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ThresholdPrefix)
	bth := make([]byte, 4)
	binary.LittleEndian.PutUint32(bth, value)
	store.Set([]byte(taipe+denom), bth)
}

// GetThreshold returns a threshold from its index
func (k Keeper) GetThreshold(ctx sdk.Context, taipe, denom string) (uint32, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.ThresholdPrefix)

	b := store.Get([]byte(taipe+denom))
	if b == nil {
		return 0, false
	}

	return binary.LittleEndian.Uint32(b), true
}
