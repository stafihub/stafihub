package types

// sudo module event types
const (
	EventTypeAddStakePool  = "add_stake_pool"
	EventTypeAddStakeItem  = "add_stake_item"
	EventTypeAddRewardPool = "add_reward_pool"
	EventTypeStake         = "stake"
	EventTypeClaimReward   = "claim_reward"
	EventTypeWithdraw      = "withdraw"

	AttributeKeyAccount         = "account"
	AttributeKeyStakeTokenDenom = "stake_token_denom"
	AttributeKeyMaxRewardPools  = "max_reward_pools"

	AttributeKeyStakeItemIndex  = "stake_item_index"
	AttributeKeyLockSecond      = "lock_second"
	AttributeKeyPowerRewardRate = "power_reward_rate"

	AttributeKeyRewardTokenDenom    = "reward_token_denom"
	AttributeKeyTotalRewardAmount   = "total_reward_amount"
	AttributeKeyRewardPerSecond     = "reward_per_second"
	AttributeKeyStartTimestamp      = "start_timestamp"
	AttributeKeyLastRewardTimestamp = "last_reward_timestamp"

	AttributeKeyStakeRecordIndex = "stake_record_index"
	AttributeKeyStakeTokenAmount = "stake_token_amount"
	AttributeKeyStakePower       = "stake_power"
	AttributeKeyEndTimestamp     = "end_timestamp"

	AttributeKeyClaimedTokens = "claimed_tokens"

	AttributeKeyWithdrawToken = "withdraw_token"
)
