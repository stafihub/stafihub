package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rdex/types"
)

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	userAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}

	lpDenom := types.GetLpTokenDenom(msg.SwapPoolIndex)

	swapPool, found := k.Keeper.GetSwapPool(ctx, lpDenom)
	if !found {
		return nil, types.ErrSwapPoolNotExit
	}
	poolBaseToken := swapPool.BaseToken
	poolToken := swapPool.Token

	inputIsBase := false
	if poolBaseToken.Denom == msg.InputToken.Denom {
		inputIsBase = true
	}

	outAmount, feeAmount := CalSwapResult(poolBaseToken.Amount, poolToken.Amount, msg.InputToken.Amount, inputIsBase)
	if !outAmount.IsPositive() {
		return nil, types.ErrSwapAmountTooFew
	}

	if outAmount.LT(msg.MinOutToken.Amount) {
		return nil, types.ErrLessThanMinOutAmount
	}

	realOutCoin := sdk.NewCoin(msg.MinOutToken.Denom, outAmount)
	if inputIsBase {
		if poolToken.Amount.LTE(outAmount) {
			return nil, types.ErrPoolTokenBalanceInsufficient
		}

		poolBaseToken.Amount = poolBaseToken.Amount.Add(msg.InputToken.Amount)
		poolToken.Amount = poolToken.Amount.Sub(outAmount)
	} else {
		if poolBaseToken.Amount.LTE(outAmount) {
			return nil, types.ErrPoolBaseTokenBalanceInsufficient
		}

		poolBaseToken.Amount = poolBaseToken.Amount.Sub(outAmount)
		poolToken.Amount = poolToken.Amount.Add(msg.InputToken.Amount)
	}

	// send coins
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, sdk.NewCoins(msg.InputToken)); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddress, sdk.NewCoins(realOutCoin)); err != nil {
		return nil, err
	}

	swapPool.BaseToken = poolBaseToken
	swapPool.Token = poolToken

	k.Keeper.SetSwapPool(ctx, lpDenom, swapPool)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSwap,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyLpDenom, lpDenom),
			sdk.NewAttribute(types.AttributeKeyInputToken, msg.InputToken.String()),
			sdk.NewAttribute(types.AttributeKeyOutputToken, realOutCoin.String()),
			sdk.NewAttribute(types.AttributeKeyFeeAmount, feeAmount.String()),
			sdk.NewAttribute(types.AttributeKeyPoolBaseTokenBalance, swapPool.BaseToken.String()),
			sdk.NewAttribute(types.AttributeKeyPoolTokenBalance, swapPool.Token.String()),
		),
	)

	return &types.MsgSwapResponse{}, nil
}
