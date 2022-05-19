package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rdex/types"
)

func (k msgServer) AddLiquidity(goCtx context.Context, msg *types.MsgAddLiquidity) (*types.MsgAddLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}
	// check provider exist
	if k.Keeper.GetProviderSwitch(ctx) && !k.Keeper.HasProvider(ctx, userAddress) {
		return nil, types.ErrProviderNotExist
	}

	tokens := msg.Tokens.Sort()
	lpDenom := types.GetLpTokenDenom(tokens)

	swapPool, found := k.Keeper.GetSwapPool(ctx, lpDenom)
	if !found {
		return nil, types.ErrSwapPoolNotExit
	}

	// check balance
	for _, token := range tokens {
		balance := k.bankKeeper.GetBalance(ctx, userAddress, token.Denom)
		if balance.Amount.LT(token.Amount) {
			return nil, types.ErrInsufficientTokenBalance
		}
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, tokens); err != nil {
		return nil, types.ErrInsufficientTokenBalance
	}

	newTotalPoolUnit, addLpUnit := CalPoolUnit(swapPool.LpToken.Amount, swapPool.Tokens[0].Amount, swapPool.Tokens[1].Amount, tokens[0].Amount, tokens[1].Amount)
	if !addLpUnit.IsPositive() {
		return nil, types.ErrAddLpUnitZero
	}

	lpTokenCoins := sdk.NewCoins(sdk.NewCoin(lpDenom, addLpUnit))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, lpTokenCoins); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddress, lpTokenCoins); err != nil {
		return nil, err
	}

	swapPool.LpToken.Amount = newTotalPoolUnit
	swapPool.Tokens = swapPool.Tokens.Add(tokens...)

	k.Keeper.SetSwapPool(ctx, lpDenom, swapPool)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddLiquidity,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyLpDenom, lpDenom),
			sdk.NewAttribute(types.AttributeKeyAddTokens, tokens.String()),
			sdk.NewAttribute(types.AttributeKeyNewTotalUnit, newTotalPoolUnit.String()),
			sdk.NewAttribute(types.AttributeKeyAddLpUnit, addLpUnit.String()),
			sdk.NewAttribute(types.AttributeKeyPoolTokensBalance, swapPool.Tokens.String()),
		),
	)
	return &types.MsgAddLiquidityResponse{}, nil
}
