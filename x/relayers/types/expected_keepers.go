package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SudoKeeper interface {
	IsAdmin(ctx sdk.Context, address string) bool
	IsDenomValid(ctx sdk.Context, denom string) bool
}
