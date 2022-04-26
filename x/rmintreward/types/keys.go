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
	MintRewardActStoreKeyPrefix   = []byte{0x01}
	ActLatestCycleStoreKeyPrefix  = []byte{0x02}
	UserClaimInfoStoreKeyPrefix   = []byte{0x03}
	UserActsStoreKeyPrefix        = []byte{0x04}
	UserMintCountStoreKeyPrefix   = []byte{0x05}
	ActCurrentCycleStoreKeyPrefix = []byte{0x07}
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
func ActCurrentCycleStoreKey(denom string) []byte {
	return append(ActCurrentCycleStoreKeyPrefix, []byte(denom)...)
}

// prefix + accountLen + account + denomLen + denom + cycle + mintIndex
func UserClaimInfoStoreKey(account sdk.AccAddress, denom string, cycle uint64, mintIndex uint64) []byte {
	prefixLen := len(UserClaimInfoStoreKeyPrefix)
	accountLen := len(account)
	denomLen := len([]byte(denom))

	key := make([]byte, prefixLen+1+accountLen+1+denomLen+8+8)

	copy(key, UserClaimInfoStoreKeyPrefix)
	key[prefixLen] = byte(accountLen)
	copy(key[prefixLen+1:], account)
	key[prefixLen+1+accountLen] = byte(denomLen)
	copy(key[prefixLen+1+accountLen+1:], []byte(denom))
	copy(key[prefixLen+1+accountLen+1+denomLen:], sdk.Uint64ToBigEndian(cycle))
	copy(key[prefixLen+1+accountLen+1+denomLen+8:], sdk.Uint64ToBigEndian(mintIndex))

	return key
}

// prefix + accountLen + account + denom
func UserActsStoreKey(account sdk.AccAddress, denom string) []byte {
	prefixLen := len(UserActsStoreKeyPrefix)
	accountLen := len(account)
	denomLen := len([]byte(denom))

	key := make([]byte, prefixLen+1+accountLen+denomLen)

	copy(key, UserActsStoreKeyPrefix)
	key[prefixLen] = byte(accountLen)
	copy(key[prefixLen+1:], account)
	copy(key[prefixLen+1+accountLen:], denom)

	return key
}

// prefix + accountLen + account + denomLen + denom + cycle
func UserMintCountStoreKey(account sdk.AccAddress, denom string, cycle uint64) []byte {
	prefixLen := len(UserMintCountStoreKeyPrefix)
	accountLen := len(account)
	denomLen := len([]byte(denom))

	key := make([]byte, prefixLen+1+accountLen+1+denomLen+8)

	copy(key, UserMintCountStoreKeyPrefix)
	key[prefixLen] = byte(accountLen)
	copy(key[prefixLen+1:], account)
	key[prefixLen+1+accountLen] = byte(accountLen)
	copy(key[prefixLen+1+accountLen+1:], []byte(denom))
	copy(key[prefixLen+1+accountLen+1+denomLen:], sdk.Uint64ToBigEndian(cycle))

	return key
}
