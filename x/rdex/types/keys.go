package types

const (
	// ModuleName defines the module name
	ModuleName = "rdex"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rdex"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	SwapPoolStoreKeyPrefix = []byte{0x01}
)

func SwapPoolStoreKey(denom string) []byte {
	return append(SwapPoolStoreKeyPrefix, []byte(denom)...)
}
