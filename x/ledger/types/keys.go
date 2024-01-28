package types

import (
	"encoding/binary"
	fmt "fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/utils"
)

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
	DefaultStakingRewardCommission = utils.MustNewDecFromStr("0.1")
	DefaultUnbondCommission        = utils.MustNewDecFromStr("0.002")
	DefaultUnbondRelayFee          = sdk.NewCoin(utils.FisDenom, sdk.NewInt(1000000))
	DefaultEraUnbondLimit          = uint32(200)
	ICATxTimeout                   = time.Hour * 30
)

var (
	BondedPoolPrefix                 = []byte{0x01}
	EraUnbondLimitPrefix             = []byte{0x02}
	ChainBondingDurationPrefix       = []byte{0x03}
	PoolDetailPrefix                 = []byte{0x04}
	SnapshotPrefix                   = []byte{0x06}
	CurrentEraSnapshotPrefix         = []byte{0x07}
	BondPipelinePrefix               = []byte{0x08}
	ChainEraPrefix                   = []byte{0x09}
	EraSnapshotPrefix                = []byte{0x0a}
	StakingRewardCommissionPrefix    = []byte{0x0b}
	ProtocolFeeReceiverPrefix        = []byte{0x0c}
	TotalExpectedActivePrefix        = []byte{0x0d}
	PoolUnbondPrefix                 = []byte{0x0e}
	ExchangeRateKeyPrefix            = []byte{0x0f}
	EraExchangeRateKeyPrefix         = []byte{0x10}
	UnbondFeePrefix                  = []byte{0x11}
	UnbondCommissionPrefix           = []byte{0x12}
	BondRecordPrefix                 = []byte{0x14}
	SignaturePrefix                  = []byte{0x15}
	RParamsPrefix                    = []byte{0x16}
	RValidatorIndicatorPrefix        = []byte{0x17}
	RValidatorPrefix                 = []byte{0x18}
	TotalProtocolFeePrefix           = []byte{0x19}
	RelayFeeReceiverPrefix           = []byte{0x1a}
	UnbondSwitchPrefix               = []byte{0x1b}
	PoolUnbondNextSequencePrefix     = []byte{0x1c}
	MigrateInitSealedStatePrefix     = []byte{0x1d}
	IcaPoolNextIndexPrefix           = []byte{0x1f}
	IcaPoolDetailPrefix              = []byte{0x20}
	IcaPoolDelegationAddrIndexPrefix = []byte{0x21}
	InterchainTxPropIdPrefix         = []byte{0x22}
	InterchainTxPropSeqIndexPrefix   = []byte{0x23}
	TotalExpectedFeePrefix           = []byte{0x24}
	LatestLsmBondProposalIdKey       = []byte{0x25}
)

var (
	SwitchStateClose    = []byte{0x00}
	SwitchStateOpen     = []byte{0x01}
	DelegationOwnerTail = "delegation"
	WithdrawalOwnerTail = "withdrawal"
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

// prefix + denomLen + denom + pool
func BondPipelineStoreKey(denom, pool string) []byte {
	denomLen := len([]byte(denom))
	poolLen := len([]byte(pool))

	key := make([]byte, 1+1+denomLen+poolLen)
	copy(key[0:], BondPipelinePrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	copy(key[2+denomLen:], []byte(pool))
	return key
}

// prefix + denomLen + denom + pool
func PoolDetailStoreKey(denom, pool string) []byte {
	denomLen := len([]byte(denom))
	poolLen := len([]byte(pool))

	key := make([]byte, 1+1+denomLen+poolLen)
	copy(key[0:], PoolDetailPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	copy(key[2+denomLen:], []byte(pool))
	return key
}

// prefix + denomLen + denom + poolLen + pool + era + seq
func PoolUnbondStoreKey(denom string, pool string, era, seq uint32) []byte {
	denomLen := len([]byte(denom))
	poolLen := len([]byte(pool))

	key := make([]byte, 1+1+denomLen+1+poolLen+4+4)
	copy(key[0:], PoolUnbondPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	key[1+1+denomLen] = byte(poolLen)
	copy(key[1+1+denomLen+1:], []byte(pool))

	binary.LittleEndian.PutUint32(key[1+1+denomLen+1+poolLen:], era)
	binary.LittleEndian.PutUint32(key[1+1+denomLen+1+poolLen+4:], seq)
	return key
}

// prefix + denomLen + denom + txHash
func BondRecordStoreKey(denom, txHash string) []byte {
	denomLen := len([]byte(denom))
	txHashLen := len([]byte(txHash))

	key := make([]byte, 1+1+denomLen+txHashLen)
	copy(key[0:], BondRecordPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	copy(key[2+denomLen:], []byte(txHash))
	return key
}

// prefix + denomLen + denom + 4
func IcaPoolDetailStoreKey(denom string, index uint32) []byte {
	denomLen := len([]byte(denom))

	key := make([]byte, 1+1+denomLen+4)
	copy(key[0:], IcaPoolDetailPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	binary.LittleEndian.PutUint32(key[2+denomLen:], index)

	return key
}

// prefix + delegationAddr
func IcaPoolDelegationAddrIndexStoreKey(delegationAddr string) []byte {
	delegationAddrLen := len([]byte(delegationAddr))

	key := make([]byte, 1+delegationAddrLen)
	copy(key[0:], IcaPoolDelegationAddrIndexPrefix)
	copy(key[1:], []byte(delegationAddr))

	return key
}

func GetOwners(denom string, index uint32) (string, string) {
	delegationOwner := fmt.Sprintf("%s-%d-%s", denom, index, DelegationOwnerTail)
	withdrawOwner := fmt.Sprintf("%s-%d-%s", denom, index, WithdrawalOwnerTail)
	return delegationOwner, withdrawOwner
}

func InterchainTxPropIdKey(proposalId string) []byte {
	return append(InterchainTxPropIdPrefix, []byte(proposalId)...)
}

// prefix + 1 + portIdLen + 1 + channelIdLen + 8
func InterchainTxPropSeqIndexStoreKey(ctrlPortId, ctrlChannelId string, sequence uint64) []byte {
	ctrlPortIdLen := len(ctrlPortId)
	ctrlChannelIdLen := len(ctrlChannelId)

	key := make([]byte, 1+1+ctrlPortIdLen+1+ctrlChannelIdLen+8)
	copy(key[0:], InterchainTxPropSeqIndexPrefix)
	key[1] = byte(ctrlPortIdLen)
	copy(key[2:], []byte(ctrlPortId))
	key[2+ctrlPortIdLen] = byte(ctrlChannelIdLen)
	copy(key[2+ctrlPortIdLen+1:], []byte(ctrlChannelId))

	copy(key[2+ctrlPortIdLen+1+ctrlChannelIdLen:], sdk.Uint64ToBigEndian(sequence))

	return key
}
