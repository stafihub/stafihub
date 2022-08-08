package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rvalidator module sentinel errors
var (
	ErrRValidatorAlreadyExist              = sdkerrors.Register(ModuleName, 1101, "rValidator already exist")
	ErrRValidatorNotExist                  = sdkerrors.Register(ModuleName, 1102, "rValidator not exist")
	ErrCycleBehindLatestCycle              = sdkerrors.Register(ModuleName, 1103, "cycle behind latest voted cycle")
	ErrCycleVersionNotMatch                = sdkerrors.Register(ModuleName, 1104, "cycle version not match")
	ErrLatestVotedCycleNotDealed           = sdkerrors.Register(ModuleName, 1105, "latest voted cycle not dealed")
	ErrLedgerIsBusyWithEra                 = sdkerrors.Register(ModuleName, 1106, "ledger is busy with era")
	ErrReportCycleNotMatchLatestVotedCycle = sdkerrors.Register(ModuleName, 1107, "report cycle not match latest voted cycle")
	ErrLedgerChainEraNotExist              = sdkerrors.Register(ModuleName, 1108, "ledger chain era not exist")
	ErrDealingRvalidatorNotFound           = sdkerrors.Register(ModuleName, 1109, "dealing rvalidator not found")
	ErrOldEqualNewRValidator               = sdkerrors.Register(ModuleName, 1110, "old euqal new rValidator")
)
