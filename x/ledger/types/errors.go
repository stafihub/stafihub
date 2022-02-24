package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/ledger module sentinel errors
var (
	ErrPoolAlreadyAdded       = sdkerrors.Register(ModuleName, 1, "pool already added error")
	ErrPoolNotFound           = sdkerrors.Register(ModuleName, 2, "pool not found error")
	ErrPoolNotBonded          = sdkerrors.Register(ModuleName, 3, "pool not bonded error")
	ErrRepeatInitBond         = sdkerrors.Register(ModuleName, 4, "repeat init bond error")
	ErrActiveAlreadySet       = sdkerrors.Register(ModuleName, 5, "active already set error")
	ErrBondPipelineNotFound   = sdkerrors.Register(ModuleName, 6, "bond pipeline not found error")
	ErrEraNotContinuable      = sdkerrors.Register(ModuleName, 7, "era not continuable error")
	ErrLastVoterNobody        = sdkerrors.Register(ModuleName, 8, "last era nobody error")
	ErrEraSkipped             = sdkerrors.Register(ModuleName, 9, "era skipped error")
	ErrNoProtocolFeeReceiver  = sdkerrors.Register(ModuleName, 10, "no protocol receiver error")
	ErrSnapshotNotFound       = sdkerrors.Register(ModuleName, 11, "snapshot not found error")
	ErrStateNotBondReported   = sdkerrors.Register(ModuleName, 12, "state not bond reported error")
	ErrStateNotEraUpdated     = sdkerrors.Register(ModuleName, 13, "state not era updated error")
	ErrRateIsNone             = sdkerrors.Register(ModuleName, 14, "rate is none error")
	ErrStateNotActiveReported = sdkerrors.Register(ModuleName, 15, "state not active reported error")
	ErrStateNotTransferable   = sdkerrors.Register(ModuleName, 16, "state not transferable error")
	ErrTransferReported       = sdkerrors.Register(ModuleName, 17, "transfer reported error")
	ErrChainEraNotFound       = sdkerrors.Register(ModuleName, 18, "chain era not found error")
	ErrBondingDurationNotSet  = sdkerrors.Register(ModuleName, 19, "bonding duration not set error")
	ErrPoolLimitReached       = sdkerrors.Register(ModuleName, 20, "pool limit reached error")
	ErrBondRepeated           = sdkerrors.Register(ModuleName, 21, "bond repeated error")
	ErrInvalidBonder          = sdkerrors.Register(ModuleName, 22, "invalid bonder error")
	ErrCommissionTooBig       = sdkerrors.Register(ModuleName, 23, "commission too big error")
	ErrInvalidEra             = sdkerrors.Register(ModuleName, 24, "invalid era error")
	ErrSignatureRepeated      = sdkerrors.Register(ModuleName, 25, "signature repeated error")
	ErrPoolDetailNotFound     = sdkerrors.Register(ModuleName, 26, "pool detail not found error")
	ErrAddrPrefixNotMatched   = sdkerrors.Register(ModuleName, 27, "addr prefix not matched error")
	ErrLockedDenomNotMatch    = sdkerrors.Register(ModuleName, 28, "locked not matched error")
	ErrLockedNotEnough        = sdkerrors.Register(ModuleName, 29, "locked not enough error")
	ErrRValidatorAlreadyExist = sdkerrors.Register(ModuleName, 30, "rvalidator already exist error")
	ErrNoRelayFeeReceiver     = sdkerrors.Register(ModuleName, 31, "no relay receiver error")
)
