package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	relayertypes "github.com/stafihub/stafihub/x/relayers/types"
)

type SudoKeeper interface {
	IsAdmin(ctx sdk.Context, address string) bool
}

type RelayerKeeper interface {
	HasRelayer(ctx sdk.Context, arena, denom, address string) bool
	GetThreshold(ctx sdk.Context, arena, denom string) (relayertypes.Threshold, bool)
}
