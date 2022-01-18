package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	relayerstypes "github.com/stafiprotocol/stafihub/x/relayers/types"
)

type SudoKeeper interface {
	IsAdmin(ctx sdk.Context, address string) bool
}

type RelayerKeeper interface {
	IsRelayer(ctx sdk.Context, denom, address string) bool
	GetThreshold(ctx sdk.Context, denom string) (val relayerstypes.Threshold, found bool)
	RelayerCount(ctx sdk.Context, denom string) int32
	SetLastVoter(ctx sdk.Context, denom, voter string)
	LastVoter(ctx sdk.Context, denom string) (val relayerstypes.LastVoter, found bool)
}

