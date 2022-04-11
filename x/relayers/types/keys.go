package types

const (
	// ModuleName defines the module name
	ModuleName = "relayers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_relayers"
)

var (
	RelayerPrefix      = []byte{0x00}
	RelayerCountPrefix = []byte{0x01}
	ThresholdPrefix    = []byte{0x02}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// [prefix]+[len(arena)]+[arena]+[len(denom)]+[denom]+[len(addr)]+[addr]
func RelayerStoreKey(arena, denom, addr string) []byte {
	key := make([]byte, 1+3+len(arena)+len(denom)+len(addr))
	key[0] = RelayerPrefix[0]
	key[1] = byte(len(arena))
	copy(key[2:], arena)
	key[2+len(arena)] = byte(len(denom))
	copy(key[3+len(arena):], denom)
	key[3+len(arena)+len(denom)] = byte(len(addr))
	copy(key[4+len(arena)+len(denom):], addr)
	return key
}
