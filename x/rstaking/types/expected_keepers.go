package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

type MintKeeper interface {
	GetMinter(ctx sdk.Context) (minter minttypes.Minter)
	GetParams(ctx sdk.Context) (params minttypes.Params)
}

type SudoKeeper interface {
	IsAdmin(ctx sdk.Context, address string) bool
	GetAdmin(ctx sdk.Context) sdk.AccAddress
}
