package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	ledgerTypes "github.com/stafihub/stafihub/x/ledger/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

type SudoKeeper interface {
	IsAdmin(ctx sdk.Context, address string) bool
}

type RBankKeeper interface {
	CheckValAddress(ctx sdk.Context, denom, address string) error
}

type LedgerKeeper interface {
	CurrentEraSnapshots(ctx sdk.Context, denom string) ledgerTypes.EraSnapshot
	GetChainEra(ctx sdk.Context, denom string) (val ledgerTypes.ChainEra, found bool)
}
