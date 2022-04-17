package types

const (
	// ModuleName defines the module name
	ModuleName = "rbank"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rbank"
)

var AccAddressPrefix = []byte{0x01}
var ValAddressPrefix = []byte{0x02}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
