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

var (
	RValidatorPrefix          = []byte{0x01}
	RValidatorIndicatorPrefix = []byte{0x02}
	RValidatorSetPrefix       = []byte{0x03}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
