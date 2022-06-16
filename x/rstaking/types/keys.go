package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "rstaking"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rstaking"
)

var (
	ValAddressStoreKeyPrefix       = []byte{0x00}
	InflationBaseKey               = []byte{0x01}
	WhitelistSwitchKey             = []byte{0x02}
	DelegatorAddressStoreKeyPrefix = []byte{0x03}
	DelegatorWhitelistSwitchKey    = []byte{0x04}
)

var (
	SwitchStateClose = []byte{0x00}
	SwitchStateOpen  = []byte{0x01}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func ValAddressStoreKey(addr sdk.ValAddress) []byte {
	return append(ValAddressStoreKeyPrefix, addr.Bytes()...)
}

func DelegatorAddressStoreKey(addr sdk.AccAddress) []byte {
	return append(DelegatorAddressStoreKeyPrefix, addr.Bytes()...)
}
