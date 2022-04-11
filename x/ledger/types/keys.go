package types

const (
	// ModuleName defines the module name
	ModuleName = "ledger"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ledger"
)

var (
	BondedPoolPrefix              = []byte{0x01}
	EraUnbondLimitPrefix          = []byte{0x02}
	ChainBondingDurationPrefix    = []byte{0x03}
	PoolDetailPrefix              = []byte{0x04}
	SnapshotPrefix                = []byte{0x06}
	CurrentEraSnapshotPrefix      = []byte{0x07}
	BondPipelinePrefix            = []byte{0x08}
	ChainEraPrefix                = []byte{0x09}
	EraSnapshotPrefix             = []byte{0x0a}
	StakingRewardCommissionPrefix = []byte{0x0b}
	ProtocolFeeReceiverPrefix     = []byte{0x0c}
	TotalExpectedActivePrefix     = []byte{0x0d}
	PoolUnbondPrefix              = []byte{0x0e}
	ExchangeRateKeyPrefix         = []byte{0x0f}
	EraExchangeRateKeyPrefix      = []byte{0x10}
	UnbondFeePrefix               = []byte{0x11}
	UnbondCommissionPrefix        = []byte{0x12}
	AccountUnbondPrefix           = []byte{0x13}
	BondRecordPrefix              = []byte{0x14}
	SignaturePrefix               = []byte{0x15}
	RParamsPrefix                 = []byte{0x16}
	RValidatorIndicatorPrefix     = []byte{0x17}
	RValidatorPrefix              = []byte{0x18}
	TotalProtocolFeePrefix        = []byte{0x19}
	RelayFeeReceiverPrefix        = []byte{0x1a}
	UnbondSwitchPrefix            = []byte{0x1b}
)

const (
	AccountMaxUnbondChunks = 32
	AccountMinUnbondChunks = 16
)

var (
	SwitchStateClose = []byte{0x00}
	SwitchStateOpen  = []byte{0x01}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func StakingRewardCommissionStoreKey(denom string) []byte {
	return append(StakingRewardCommissionPrefix, []byte(denom)...)
}

func UnbondCommissionStoreKey(denom string) []byte {
	return append(UnbondCommissionPrefix, []byte(denom)...)
}

func RelayFeeReceiverStorekey(denom string) []byte {
	return append(RelayFeeReceiverPrefix, []byte(denom)...)
}

func UnbondSwitchStoreKey(denom string) []byte {
	return append(UnbondSwitchPrefix, []byte(denom)...)
}
