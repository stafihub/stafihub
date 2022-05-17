package types

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

var RewardFactor = sdk.NewInt(1e12)

var (
	StakePoolStoreKeyPrefix            = []byte{0x01}
	StakeItemStoreKeyPrefix            = []byte{0x02}
	UserStakeRecordStoreKeyPrefix      = []byte{0x03}
	UserStakeRecordIndexStoreKeyPrefix = []byte{0x04}
	RewardPoolIndexStoreKeyPrefix      = []byte{0x05}
	StakeItemIndexStoreKey             = []byte{0x06}
)

func StakePoolStoreKey(denom string) []byte {
	return append(StakePoolStoreKeyPrefix, []byte(denom)...)
}

func StakeItemStoreKey(index uint32) []byte {
	indexBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(indexBts, index)
	return append(StakeItemStoreKeyPrefix, indexBts...)
}

// prefix + len(userAddress) + userAddress + len(stakeTokenDenom) + stakeTokenDenom + index
func UserStakeRecordStoreKey(userAddress, stakeTokenDenom string, index uint32) []byte {
	userAddressLen := len(userAddress)
	stakeTokenDenomLen := len(stakeTokenDenom)

	key := make([]byte, 1+1+userAddressLen+1+stakeTokenDenomLen+4)
	key[0] = UserStakeRecordStoreKeyPrefix[0]
	key[1] = byte(len(userAddress))
	copy(key[2:2+userAddressLen], userAddress)
	key[2+userAddressLen] = byte(stakeTokenDenomLen)
	copy(key[2+userAddressLen+1:2+userAddressLen+1+stakeTokenDenomLen], stakeTokenDenom)

	binary.LittleEndian.PutUint32(key[2+userAddressLen+1+stakeTokenDenomLen:], index)
	return key
}
