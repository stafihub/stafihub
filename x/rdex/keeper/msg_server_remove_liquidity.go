package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stafihub/stafihub/x/rdex/types"
)

func (k msgServer) RemoveLiquidity(goCtx context.Context, msg *types.MsgRemoveLiquidity) (*types.MsgRemoveLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	userAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}
	minOutTokens := msg.MinOutTokens.Sort()
	lpDenom := types.GetLpTokenDenom(minOutTokens)

	swapPool, found := k.Keeper.GetSwapPool(ctx, lpDenom)
	if !found {
		return nil, types.ErrSwapPoolNotExit
	}
	poolBaseToken := swapPool.Tokens[0]
	poolToken := swapPool.Tokens[1]
	poolLpToken := swapPool.LpToken

	userLpTokenBalance := k.bankKeeper.GetBalance(ctx, userAddress, lpDenom)
	if userLpTokenBalance.Amount.LT(msg.RmUnit) {
		return nil, types.ErrInsufficientLpTokenBalance
	}
	moduleAddress := authTypes.NewModuleAddress(types.ModuleName)
	poolBaseTokenBalance := k.bankKeeper.GetBalance(ctx, moduleAddress, poolBaseToken.Denom)
	poolTokenBalance := k.bankKeeper.GetBalance(ctx, moduleAddress, poolToken.Denom)

	if !msg.RmUnit.IsPositive() || msg.RmUnit.GT(poolLpToken.Amount) || msg.SwapUnit.GT(msg.RmUnit) {
		return nil, types.ErrUnitAmount
	}

	inputIsBase := false
	if poolBaseToken.Denom == msg.InputTokenDenom {
		inputIsBase = true
	}

	rmBaseTokenAmount, rmTokenAmount, swapInputAmount := CalRemoveAmount(poolLpToken.Amount, msg.RmUnit, msg.SwapUnit, poolBaseToken.Amount, poolToken.Amount, inputIsBase)
	poolLpToken.Amount = poolLpToken.Amount.Sub(msg.RmUnit)
	poolBaseToken.Amount = poolBaseToken.Amount.Sub(rmBaseTokenAmount)
	poolToken.Amount = poolToken.Amount.Sub(rmTokenAmount)

	if swapInputAmount.IsPositive() {
		swapResult, _ := CalSwapResult(poolBaseToken.Amount, poolToken.Amount, swapInputAmount, inputIsBase)
		if !swapResult.IsPositive() {
			return nil, types.ErrSwapAmountTooFew
		}

		if inputIsBase {
			if swapResult.GTE(poolToken.Amount) {
				return nil, types.ErrPoolRTokenBalanceInsufficient
			}

			poolBaseToken.Amount = poolBaseToken.Amount.Add(swapInputAmount)
			poolToken.Amount = poolToken.Amount.Sub(swapResult)

			rmBaseTokenAmount = rmBaseTokenAmount.Sub(swapInputAmount)
			rmTokenAmount = rmTokenAmount.Add(swapResult)
		} else {
			if swapResult.GTE(poolBaseToken.Amount) {
				return nil, types.ErrPoolFisBalanceInsufficient
			}

			poolToken.Amount = poolToken.Amount.Add(swapInputAmount)
			poolBaseToken.Amount = poolBaseToken.Amount.Sub(swapResult)

			rmTokenAmount = rmTokenAmount.Sub(swapInputAmount)
			rmBaseTokenAmount = rmBaseTokenAmount.Add(swapResult)
		}
	}

	if rmBaseTokenAmount.LT(minOutTokens[0].Amount) || rmTokenAmount.LT(minOutTokens[1].Amount) {
		return nil, types.ErrLessThanMinOutAmount
	}

	if rmBaseTokenAmount.GT(poolBaseTokenBalance.Amount) {
		return nil, types.ErrPoolFisBalanceInsufficient
	}

	if rmTokenAmount.GT(poolTokenBalance.Amount) {
		return nil, types.ErrPoolRTokenBalanceInsufficient
	}

	if (poolBaseToken.Amount.IsZero() && !poolToken.Amount.IsZero()) || (poolToken.Amount.IsZero() && !poolBaseToken.Amount.IsZero()) {
		return nil, types.ErrPoolOneSideZero
	}

	willSendCoin := sdk.NewCoins()
	if rmBaseTokenAmount.IsPositive() {
		willSendCoin = willSendCoin.Add(sdk.NewCoin(poolBaseToken.Denom, rmBaseTokenAmount))
	}
	if rmTokenAmount.IsPositive() {
		willSendCoin = willSendCoin.Add(sdk.NewCoin(poolToken.Denom, rmTokenAmount))
	}

	if willSendCoin.IsAllPositive() {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddress, willSendCoin); err != nil {
			return nil, err
		}
	}

	willBurnLp := sdk.NewCoin(lpDenom, msg.RmUnit)
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, sdk.NewCoins(willBurnLp)); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(willBurnLp)); err != nil {
		return nil, err
	}

	swapPool.Tokens[0] = poolBaseToken
	swapPool.Tokens[1] = poolToken
	swapPool.LpToken = poolLpToken

	k.SetSwapPool(ctx, lpDenom, swapPool)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRemoveLiquidity,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyLpDenom, lpDenom),
			sdk.NewAttribute(types.AttributeKeyRemoveUnit, msg.RmUnit.String()),
			sdk.NewAttribute(types.AttributeKeySwapUnit, msg.SwapUnit.String()),
			sdk.NewAttribute(types.AttributeKeyNewTotalUnit, swapPool.LpToken.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyRemoveBaseTokenAmount, rmBaseTokenAmount.String()),
			sdk.NewAttribute(types.AttributeKeyRemoveTokenAmount, rmTokenAmount.String()),
			sdk.NewAttribute(types.AttributeKeyPoolTokensBalance, swapPool.Tokens.String()),
		),
	)
	return &types.MsgRemoveLiquidityResponse{}, nil
}
