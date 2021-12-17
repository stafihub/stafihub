package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) IsAdminOrRelayer(ctx sdk.Context, denom, address string) bool {
	return k.relayerKeeper.CheckIsRelayer(ctx, denom, address) ||
		k.sudoKeeper.IsAdmin(ctx, address)
}
