package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "rmintreward"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rmintreward"
)

var (
	MintRewardActStoreKeyPrefix  = []byte{0x01}
	ActLatestCycleStoreKeyPrefix = []byte{0x02}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func MintRewardActStoreKey(denom string, cycle uint64) []byte {
	denomBts := []byte(denom)
	prefixLen := len(MintRewardActStoreKeyPrefix)
	denomLen := len(denomBts)
	key := make([]byte, prefixLen+8+denomLen)
	copy(key, MintRewardActStoreKeyPrefix)
	copy(key[prefixLen:], denomBts)
	copy(key[prefixLen+denomLen:], sdk.Uint64ToBigEndian(cycle))
	return key
}

func ActLatestCycleStoreKey(denom string) []byte {
	return append(ActLatestCycleStoreKeyPrefix, []byte(denom)...)
}
