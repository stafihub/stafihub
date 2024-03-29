package types

// sudo module event types
const (
	EventTypePoolAdded                = "pool_added"
	EventTypePoolRemoved              = "pool_removed"
	EventTypeEraPoolUpdated           = "era_pool_updated"
	EventTypeBondReported             = "bond_reported"
	EventTypeActiveReported           = "active_reported"
	EventTypeWithdrawReported         = "withdraw_reported"
	EventTypeTransferReported         = "transfer_reported"
	EventTypeBondExecuted             = "bond_executed"
	EventTypeNativeAndLsmBondExecuted = "native_and_lsm_bond_executed"
	EventTypeLiquidityUnbond          = "liquidity_unbond"
	EventTypeSignatureEnough          = "signature_enough"
	EventTypeSignatureSubmitted       = "signature_submitted"
	EventTypeRParamsChanged           = "rparams_changed"
	EventTypeInitPool                 = "init_pool"
	EventTypeRemovePool               = "remove_pool"

	AttributeKeyDenom            = "denom"
	AttributeKeyPool             = "pool"
	AttributeKeyLastEra          = "last_era"
	AttributeKeyCurrentEra       = "current_era"
	AttributeKeyShotId           = "shot_id"
	AttributeKeyBonder           = "bonder"
	AttributeKeyUnbonder         = "unbonder"
	AttributeKeyBlockhash        = "blockhash"
	AttributeKeyTxhash           = "txhash"
	AttributeKeyBalance          = "balance"
	AttributeKeyNativeBondAmount = "nativeBondAmount"
	AttributeKeyLsmBondAmount    = "lsmBondAmount"
	AttributeKeyRbalance         = "rbalance"
	AttributeKeyUnBondAmount     = "unbond_amount"
	AttributeKeyExchangeAmount   = "exchange_amount"
	AttributeKeyReceiveAmount    = "receive_amount"
	AttributeKeyReceiver         = "receiver"
	AttributeKeyEra              = "era"
	AttributeKeyUnlockEra        = "unlock_era"
	AttributeKeyTxType           = "tx_type"
	AttributeKeyPropId           = "prop_id"
	AttributeKeySigner           = "signer"

	AttributeKeyGasPrice        = "gas_price"
	AttributeKeyEraSeconds      = "era_seconds"
	AttributeKeyOffset          = "offset"
	AttributeKeyBondingDuration = "bonding_duration"
	AttributeKeyLeastBond       = "least_bond"
)
