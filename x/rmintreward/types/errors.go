package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rmintreward module sentinel errors
var (
	ErrActBeginBlockTooSmall            = sdkerrors.Register(ModuleName, 1100, "act begin block must greater than zero")
	ErrActEndBlockTooSamll              = sdkerrors.Register(ModuleName, 1101, "act end block must greater than begin block")
	ErrActTotalRewardTooSmall           = sdkerrors.Register(ModuleName, 1102, "act total reward must greater than zero")
	ErrActTotalRewardLessThanUserLimit  = sdkerrors.Register(ModuleName, 1103, "act total reward less than userLimit")
	ErrActLockedBlocksTooSmall          = sdkerrors.Register(ModuleName, 1104, "act locked blocks must greater than zero")
	ErrActRewardRateTooSmall            = sdkerrors.Register(ModuleName, 1105, "act reward rate must greater than zero")
	ErrActEndBlockLessThanCurrentBlock  = sdkerrors.Register(ModuleName, 1106, "act end block less than current block")
	ErrLatestMintRewardActNotExist      = sdkerrors.Register(ModuleName, 1107, "latest mint reward act not exist")
	ErrActBeginBlockLTLatestActEndBlock = sdkerrors.Register(ModuleName, 1108, "act begin block less than latest act end block")
	ErrActNotExist                      = sdkerrors.Register(ModuleName, 1109, "act not exist")
)
