package types

const (
	// ModuleName defines the module name
	ModuleName = "rvalidator"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rvalidator"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	DefaultCycleSeconds   = uint64(60 * 10)
	DefaultShuffleSeconds = uint64(60 * 60 * 24 * 14)
)

var (
	SelectedRValidatorStoreKeyPrefix = []byte{0x01}
	LatestVotedCycleStoreKeyPrefix   = []byte{0x02}
	CycleSecondsStoreKeyPrefix       = []byte{0x03}
	ShuffleSecondsStoreKeyPrefix     = []byte{0x04}
	LatestDealedCycleStoreKeyPrefix  = []byte{0x05}
)

// prefix + denomLen + denom + poolAddressLen + poolAddress + rValidatorAddressLen + rValidatorAddress
func SelectedRValdidatorStoreKey(denom, poolAddress, rValidatorAddress string) []byte {
	denomLen := len([]byte(denom))
	poolAddressLen := len([]byte(poolAddress))
	rValidatorAddressLen := len([]byte(rValidatorAddress))

	key := make([]byte, 1+1+denomLen+1+poolAddressLen+1+rValidatorAddressLen)
	copy(key[0:], SelectedRValidatorStoreKeyPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	key[2+denomLen] = byte(poolAddressLen)
	copy(key[2+denomLen+1:], []byte(poolAddress))
	key[2+denomLen+1+poolAddressLen] = byte(rValidatorAddressLen)
	copy(key[2+denomLen+1+poolAddressLen+1:], []byte(rValidatorAddress))
	return key
}

// prefix + denomLen + denom + poolAddressLen + poolAddress
func LatestVotedCycleStoreKey(denom, poolAddress string) []byte {
	denomLen := len([]byte(denom))
	poolAddressLen := len([]byte(poolAddress))

	key := make([]byte, 1+1+denomLen+1+poolAddressLen)
	copy(key[0:], LatestVotedCycleStoreKeyPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	key[2+denomLen] = byte(poolAddressLen)
	copy(key[2+denomLen+1:], []byte(poolAddress))

	return key
}

// prefix + denomLen + denom + poolAddressLen + poolAddress
func LatestDealedCycleStoreKey(denom, poolAddress string) []byte {
	denomLen := len([]byte(denom))
	poolAddressLen := len([]byte(poolAddress))

	key := make([]byte, 1+1+denomLen+1+poolAddressLen)
	copy(key[0:], LatestDealedCycleStoreKeyPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	key[2+denomLen] = byte(poolAddressLen)
	copy(key[2+denomLen+1:], []byte(poolAddress))

	return key
}

func CycleSecondsStoreKey(denom string) []byte {
	return append(CycleSecondsStoreKeyPrefix, []byte(denom)...)
}

func ShuffleSecondsStoreKey(denom string) []byte {
	return append(ShuffleSecondsStoreKeyPrefix, []byte(denom)...)
}
