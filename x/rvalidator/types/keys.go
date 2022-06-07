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
)

// prefix + denomLen + denom + rValidator
func SelectedRValdidatorStoreKey(denom, rValidator string) []byte {
	denomLen := len([]byte(denom))
	rValidatorLen := len([]byte(rValidator))

	key := make([]byte, 1+1+denomLen+rValidatorLen)
	copy(key[0:], SelectedRValidatorStoreKeyPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	copy(key[2+denomLen:], []byte(rValidator))
	return key
}

func LatestVotedCycleStoreKey(denom string) []byte {
	return append(LatestVotedCycleStoreKeyPrefix, []byte(denom)...)
}

func CycleSecondsStoreKey(denom string) []byte {
	return append(CycleSecondsStoreKeyPrefix, []byte(denom)...)
}

func ShuffleSecondsStoreKey(denom string) []byte {
	return append(ShuffleSecondsStoreKeyPrefix, []byte(denom)...)
}
