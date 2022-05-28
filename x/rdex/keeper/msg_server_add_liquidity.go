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

	orderTokens := sdk.Coins{msg.Token0, msg.Token1}.Sort()
	lpDenom := types.GetLpTokenDenom(msg.SwapPoolIndex)

	swapPool, found := k.Keeper.GetSwapPool(ctx, lpDenom)
	if !found {
		return nil, types.ErrSwapPoolNotExit
	}

	// check balance
	willSendToken := sdk.NewCoins()
	for _, token := range orderTokens {
		balance := k.bankKeeper.GetBalance(ctx, userAddress, token.Denom)
		if balance.Amount.LT(token.Amount) {
			return nil, types.ErrUserTokenBalanceInsufficient
		}
		if token.Amount.IsPositive() {
			willSendToken = willSendToken.Add(token)
		}
	}

	newTotalPoolUnit, addLpUnit := CalPoolUnit(swapPool.LpToken.Amount, swapPool.BaseToken.Amount, swapPool.Token.Amount, orderTokens[0].Amount, orderTokens[1].Amount)
	if !addLpUnit.IsPositive() {
		return nil, types.ErrAddLpUnitZero
	}

	//send coins
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, willSendToken); err != nil {
		return nil, types.ErrUserTokenBalanceInsufficient
	}
	lpTokenCoins := sdk.NewCoins(sdk.NewCoin(lpDenom, addLpUnit))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, lpTokenCoins); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddress, lpTokenCoins); err != nil {
		return nil, err
	}

	swapPool.LpToken.Amount = newTotalPoolUnit
	swapPool.BaseToken.Amount = swapPool.BaseToken.Amount.Add(orderTokens[0].Amount)
	swapPool.Token.Amount = swapPool.Token.Amount.Add(orderTokens[1].Amount)

	k.Keeper.SetSwapPool(ctx, lpDenom, swapPool)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddLiquidity,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyLpDenom, lpDenom),
			sdk.NewAttribute(types.AttributeKeyAddBaseToken, orderTokens[0].String()),
			sdk.NewAttribute(types.AttributeKeyAddToken, orderTokens[1].String()),
			sdk.NewAttribute(types.AttributeKeyNewTotalUnit, newTotalPoolUnit.String()),
			sdk.NewAttribute(types.AttributeKeyAddLpUnit, addLpUnit.String()),
			sdk.NewAttribute(types.AttributeKeyPoolBaseTokenBalance, swapPool.BaseToken.String()),
			sdk.NewAttribute(types.AttributeKeyPoolTokenBalance, swapPool.Token.String()),
		),
	)
	return &types.MsgAddLiquidityResponse{}, nil
}
