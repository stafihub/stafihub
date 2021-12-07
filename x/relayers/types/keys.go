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
	RelayerPrefix    = []byte{0x00}
	ThresholdPrefix = []byte{0x01}
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}
