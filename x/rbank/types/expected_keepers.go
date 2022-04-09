package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool
}

type SudoKeeper interface {
	IsAdmin(ctx sdk.Context, address string) bool
}
