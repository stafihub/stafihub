package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SudoKeeper interface {
	IsAdmin(ctx sdk.Context, address string) bool
}

type RelayerKeeper interface {
	HasRelayer(ctx sdk.Context, taipe, denom, address string) bool
	GetThreshold(ctx sdk.Context, taipe, denom string) (uint32, bool)
	SetLastVoter(ctx sdk.Context, denom, voter string)
}
