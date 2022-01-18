package types

// sudo module event types
const (
	EventTypePoolAdded        = "pool_added"
	EventTypePoolRemoved      = "pool_removed"
	EventTypeEraPoolUpdated   = "era_pool_updated"
	EventTypeBondReported     = "bond_reported"
	EventTypeActiveReported   = "active_reported"
	EventTypeWithdrawReported = "withdraw_reported"
	EventTypeTransferReported = "transfer_reported"
	EventTypeBondExecuted     = "bond_executed"
	EventTypeLiquidityUnbond  = "liquidity_unbond"

	AttributeKeyDenom        = "denom"
	AttributeKeyPool         = "pool"
	AttributeKeyLastEra      = "last_era"
	AttributeKeyCurrentEra   = "current_era"
	AttributeKeyShotId       = "shot_id"
	AttributeKeyLastVoter    = "last_voter"
	AttributeKeyBonder       = "bonder"
	AttributeKeyUnbonder     = "unbonder"
	AttributeKeyBlockhash    = "blockhash"
	AttributeKeyTxhash       = "txhash"
	AttributeKeyBalance      = "balance"
	AttributeKeyRbalance     = "rbalance"
	AttributeKeyUnBondAmount = "unbond_amount"
	AttributeKeyReceiver     = "receiver"
)
