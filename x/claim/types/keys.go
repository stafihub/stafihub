package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "claim"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_claim"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	MerkleRootStoreKeyPrefix = []byte{0x01}
	ClaimRoundStoreKey       = []byte{0x02}
)

func MerkleRootStoreKey(round uint64) []byte {
	key := make([]byte, 9)

	key[0] = MerkleRootStoreKeyPrefix[0]
	copy(key[1:], sdk.Uint64ToBigEndian(round))

	return key
}
