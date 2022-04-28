package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/rdex/types"
)

func (k msgServer) AddLiquidity(goCtx context.Context, msg *types.MsgAddLiquidity) (*types.MsgAddLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}
	swapPool, found := k.Keeper.GetSwapPool(ctx, msg.Denom)
	if !found {
		return nil, types.ErrSwapPoolNotExit
	}

	rTokenBalance := k.bankKeeper.GetBalance(ctx, userAddress, msg.Denom)
	if rTokenBalance.Amount.LT(msg.RTokenAmount) {
		return nil, types.ErrInsufficientRTokenBalance
	}
	fisBalance := k.bankKeeper.GetBalance(ctx, userAddress, utils.FisDenom)
	if fisBalance.Amount.LT(msg.FisAmount) {
		return nil, types.ErrInsufficientFisBalance
	}

	newTotalPoolUnit, addLpUnit := calPoolUnit(swapPool.TotalUnit, swapPool.FisBalance, swapPool.RTokenBalance, msg.FisAmount, msg.RTokenAmount)
	if addLpUnit.LTE(sdk.ZeroInt()) {
		return nil, types.ErrAddLpUnitZero
	}

	if msg.FisAmount.GT(sdk.ZeroInt()) {
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, sdk.NewCoins(sdk.NewCoin(utils.FisDenom, msg.FisAmount))); err != nil {
			return nil, types.ErrInsufficientFisBalance
		}
	}
	if msg.RTokenAmount.GT(sdk.ZeroInt()) {
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, sdk.NewCoins(sdk.NewCoin(msg.Denom, msg.RTokenAmount))); err != nil {
			return nil, types.ErrInsufficientRTokenBalance
		}
	}
	lpTokenCoins := sdk.NewCoins(sdk.NewCoin(types.LpTokenDenom(msg.Denom), addLpUnit))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, lpTokenCoins); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddress, lpTokenCoins); err != nil {
		return nil, err
	}

	swapPool.TotalUnit = newTotalPoolUnit
	swapPool.FisBalance = swapPool.FisBalance.Add(msg.FisAmount)
	swapPool.RTokenBalance = swapPool.RTokenBalance.Add(msg.RTokenAmount)

	k.Keeper.SetSwapPool(ctx, msg.Denom, swapPool)

	// AddLiquidity: (account, symbol, fis amount, rToken amount, new total unit, add lp unit, fis balance, rtoken balance)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddLiquidity,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyFisAmount, msg.FisAmount.String()),
			sdk.NewAttribute(types.AttributeKeyRTokenAmount, msg.RTokenAmount.String()),
			sdk.NewAttribute(types.AttributeKeyNewTotalUnit, newTotalPoolUnit.String()),
			sdk.NewAttribute(types.AttributeKeyAddLpUnit, addLpUnit.String()),
			sdk.NewAttribute(types.AttributeKeyFisBalance, swapPool.FisBalance.String()),
			sdk.NewAttribute(types.AttributeKeyRTokenBalance, swapPool.RTokenBalance.String()),
		),
	)
	return &types.MsgAddLiquidityResponse{}, nil
}
