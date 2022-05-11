package types

//  Swap: (account, symbol, input amount, output amount, fee amount, input is fis, fis balance, rtoken balance)
//  Swap(AccountId, RSymbol, u128, u128, u128, bool, u128, u128),
//  CreatePool: (account, symbol, fis amount, rToken amount, new total unit, add lp unit)
//  CreatePool(AccountId, RSymbol, u128, u128, u128, u128),
//  AddLiquidity: (account, symbol, fis amount, rToken amount, new total unit, add lp unit, fis balance, rtoken balance)
//  AddLiquidity(AccountId, RSymbol, u128, u128, u128, u128, u128, u128),
//  RemoveLiquidity: (account, symbol, rm unit, swap unit, rm fis amount, rm rToken amount, input is fis, fis balance, rtoken balance)
//  RemoveLiquidity(AccountId, RSymbol, u128, u128, u128, u128, bool, u128, u128),

// sudo module event types
const (
	EventTypeSwap            = "swap"
	EventTypeCreatePool      = "create_pool"
	EventTypeAddLiquidity    = "add_liquidity"
	EventTypeRemoveLiquidity = "remove_liquidity"

	AttributeKeyAccount           = "account"
	AttributeKeyLpDenom           = "lp_denom"
	AttributeKeyInputToken        = "input_token"
	AttributeKeyOutputToken       = "output_token"
	AttributeKeyFeeAmount         = "fee_amount"
	AttributeKeyPoolTokensBalance = "pool_tokens_balance"

	AttributeKeyAddTokens    = "add_tokens"
	AttributeKeyNewTotalUnit = "new_total_unit"
	AttributeKeyAddLpUnit    = "add_lp_unit"

	AttributeKeyRemoveUnit            = "remove_unit"
	AttributeKeySwapUnit              = "swap_unit"
	AttributeKeyRemoveBaseTokenAmount = "remove_base_token_amount"
	AttributeKeyRemoveTokenAmount     = "remove_token_amount"
)
