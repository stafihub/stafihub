package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
	ProposalStoreKeyPrefix          = []byte{0x03}
	ChainIdStoreKeyPrefix           = []byte{0x04}
	ResourceIdTypeStoreKeyPrefix    = []byte{0x05}
	RelayFeeReceiverStoreKey        = []byte{0x06}
	ThresholdStoreKeyPrefix         = []byte{0x07}
	RelayFeeStoreKeyPrefix          = []byte{0x08}
)

type ResourceIdType [1]byte

var (
	ResourceIdTypeForeign = ResourceIdType([1]byte{0x00})
	ResourceIdTypeNative  = ResourceIdType([1]byte{0x01})
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func RelayStoreKey(chainId uint8, addr sdk.AccAddress) []byte {
	return append(append(RelayerStoreKeyPrefix, chainId), addr.Bytes()...)
}

func ThresholdStoreKey(chaindId uint8) []byte {
	return append(ThresholdStoreKeyPrefix, chaindId)
}

func ChainIdStoreKey(chaindId uint8) []byte {
	return append(ChainIdStoreKeyPrefix, chaindId)
}

func ResourceIdToDenomStoreKey(resourceId [32]byte) []byte {
	return append(ResourceIdToDenomStoreKeyPrefix, resourceId[:]...)
}

func DepositCountsStoreKey(chainId uint8) []byte {
	return append(DepositCountsStoreKeyPrefix, chainId)
}

func ResourceIdTypeStoreKey(resourceId [32]byte) []byte {
	return append(ResourceIdTypeStoreKeyPrefix, resourceId[:]...)
}

func ProposalStoreKey(chainId uint8, depositNonce uint64, hash [32]byte) []byte {
	key := make([]byte, 41)
	key[0] = chainId
	copy(key[1:], sdk.Uint64ToBigEndian(depositNonce))
	copy(key[9:], hash[:])

	return append(ProposalStoreKeyPrefix, key...)
}

func RelayFeeStoreKey(chainId uint8) []byte {
	return append(RelayFeeStoreKeyPrefix, chainId)
}
