package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	relayerTypes "github.com/stafihub/stafihub/x/relayers/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

type BankKeeper interface {
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

type SudoKeeper interface {
	IsAdmin(ctx sdk.Context, address string) bool
	GetAdmin(ctx sdk.Context) sdk.AccAddress
}

type RelayersKeeper interface {
	HasRelayer(ctx sdk.Context, arena, denom, addr string) bool
	GetThreshold(ctx sdk.Context, arena, denom string) (relayerTypes.Threshold, bool)
}
