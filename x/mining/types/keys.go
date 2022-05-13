package types

import (
	"encoding/binary"
)

const (
	// ModuleName defines the module name
	ModuleName = "mining"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mining"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	StakePoolStoreKeyPrefix = []byte{0x01}
	StakeItemStoreKeyPrefix = []byte{0x02}
)

func StakePoolStoreKey(denom string) []byte {
	return append(StakePoolStoreKeyPrefix, []byte(denom)...)
}

func StakeItemStoreKey(index uint32) []byte {
	indexBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(indexBts, index)
	return append(StakeItemStoreKeyPrefix, indexBts...)
}
