package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "bridge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bridge"
)

var (
	RelayerStoreKeyPrefix           = []byte{0x00}
	ResourceIdToDenomStoreKeyPrefix = []byte{0x01}
	DepositCountsStoreKeyPrefix     = []byte{0x02}
)

var (
	ThresholdStoreKey = []byte("thresholdStoreKey")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func RelayStoreKey(addr sdk.AccAddress) []byte {
	return append(RelayerStoreKeyPrefix, addr.Bytes()...)
}

func ResourceIdToDenomStoreKey(resourceId [32]byte) []byte {
	return append(ResourceIdToDenomStoreKeyPrefix, resourceId[:]...)
}

func DepositCountsStoreKey(chainId uint8) []byte {
	return append(DepositCountsStoreKeyPrefix, chainId)
}
