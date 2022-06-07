package types

// sudo module event types
const (
	EventTypeSwap            = "swap"
	EventTypeCreatePool      = "create_pool"
	EventTypeAddLiquidity    = "add_liquidity"
	EventTypeRemoveLiquidity = "remove_liquidity"

	AttributeKeyAccount              = "account"
	AttributeKeyLpDenom              = "lp_denom"
	AttributeKeyInputToken           = "input_token"
	AttributeKeyOutputToken          = "output_token"
	AttributeKeyFeeAmount            = "fee_amount"
	AttributeKeyPoolBaseTokenBalance = "pool_base_token_balance"
	AttributeKeyPoolTokenBalance     = "pool_token_balance"

	AttributeKeyAddBaseToken = "add_base_token"
	AttributeKeyAddToken     = "add_token"
	AttributeKeyNewTotalUnit = "new_total_unit"
	AttributeKeyAddLpUnit    = "add_lp_unit"

	AttributeKeyRemoveUnit            = "remove_unit"
	AttributeKeySwapUnit              = "swap_unit"
	AttributeKeyRemoveBaseTokenAmount = "remove_base_token_amount"
	AttributeKeyRemoveTokenAmount     = "remove_token_amount"
)
