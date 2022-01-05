package types

// sudo module event types
const (
	EventTypePoolAdded      = "pool_added"
	EventTypePoolRemoved    = "pool_removed"
	EventTypeEraPoolUpdated = "era_pool_updated"
	EventTypeBondReported = "bond_reported"
	EventTypeActiveReported = "active_reported"
	EventTypeWithdrawReported = "withdraw_reported"
	EventTypeTransferReported = "transfer_reported"

	AttributeKeyDenom = "denom"
	AttributeKeyPoolAddress = "pool_address"
	AttributeKeyLastEra = "last_era"
	AttributeKeyCurrentEra = "current_era"
	AttributeKeyShotId = "shot_id"
	AttributeKeyLastVoter = "last_voter"

)
