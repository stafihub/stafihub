package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)


type SudoKeeper interface {
	// Methods imported from sudo should be defined here
	IsAdmin(ctx sdk.Context, address sdk.AccAddress) bool
}
