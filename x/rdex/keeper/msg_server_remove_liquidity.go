package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/rdex/types"
)

func (k msgServer) RemoveLiquidity(goCtx context.Context, msg *types.MsgRemoveLiquidity) (*types.MsgRemoveLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	userAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}

	swapPool, found := k.Keeper.GetSwapPool(ctx, msg.Denom)
	if found {
		return nil, types.ErrSwapPoolAlreadyExist
	}
	moduleAddress := authTypes.NewModuleAddress(types.ModuleName)
	poolFisBalance := k.bankKeeper.GetBalance(ctx, moduleAddress, utils.FisDenom)
	poolRTokenBalance := k.bankKeeper.GetBalance(ctx, moduleAddress, msg.Denom)

	lpTokenBalance := k.bankKeeper.GetBalance(ctx, userAddress, types.LpTokenDenom(msg.Denom))
	if lpTokenBalance.Amount.LT(msg.RmUnit) {
		return nil, types.ErrInsufficientLpTokenBalance
	}

	if msg.RmUnit.LTE(sdk.ZeroInt()) || msg.RmUnit.GT(swapPool.TotalUnit) || msg.SwapUnit.GT(msg.RmUnit) {
		return nil, types.ErrUnitAmount
	}

	rmFisAmount, rmRTokenAmount, swapInputAmount := calRemoveAmount(swapPool.TotalUnit, msg.RmUnit, msg.SwapUnit, swapPool.FisBalance, swapPool.RTokenBalance, msg.InputIsFis)
	swapPool.TotalUnit = swapPool.TotalUnit.Sub(msg.RmUnit)
	swapPool.FisBalance = swapPool.FisBalance.Sub(rmFisAmount)
	swapPool.RTokenBalance = swapPool.RTokenBalance.Sub(rmRTokenAmount)

	if swapInputAmount.GT(sdk.ZeroInt()) {
		swapResult, _ := calSwapResult(swapPool.FisBalance, swapPool.RTokenBalance, swapInputAmount, msg.InputIsFis)
		if swapResult.LTE(sdk.ZeroInt()) {
			return nil, types.ErrSwapAmountTooFew
		}

		if msg.InputIsFis {
			if swapResult.GTE(swapPool.RTokenBalance) {
				return nil, types.ErrPoolRTokenBalanceInsufficient
			}

			swapPool.FisBalance = swapPool.FisBalance.Add(swapInputAmount)
			swapPool.RTokenBalance = swapPool.RTokenBalance.Sub(swapResult)

			rmFisAmount = rmFisAmount.Sub(swapInputAmount)
			rmRTokenAmount = rmRTokenAmount.Add(swapResult)
		} else {
			if swapResult.GTE(swapPool.FisBalance) {
				return nil, types.ErrPoolFisBalanceInsufficient
			}

			swapPool.RTokenBalance = swapPool.RTokenBalance.Add(swapInputAmount)
			swapPool.FisBalance = swapPool.FisBalance.Sub(swapResult)

			rmRTokenAmount = rmRTokenAmount.Sub(swapInputAmount)
			rmFisAmount = rmFisAmount.Add(swapResult)
		}
	}

	if rmFisAmount.LT(msg.MinFisOutAmount) || rmRTokenAmount.LT(msg.MinRtokenOutAmount) {
		return nil, types.ErrLessThanMinOutAmount
	}

	if rmFisAmount.GT(poolFisBalance.Amount) {
		return nil, types.ErrPoolFisBalanceInsufficient
	}

	if rmRTokenAmount.GT(poolRTokenBalance.Amount) {
		return nil, types.ErrPoolRTokenBalanceInsufficient
	}

	if (swapPool.FisBalance.IsZero() && !swapPool.RTokenBalance.IsZero()) || (swapPool.RTokenBalance.IsZero() && !swapPool.FisBalance.IsZero()) {
		return nil, types.ErrPoolOneSideZero
	}

	willSendCoin := sdk.NewCoins()
	if rmFisAmount.GT(sdk.ZeroInt()) {
		willSendCoin.Add(sdk.NewCoin(utils.FisDenom, rmFisAmount))
	}
	if rmRTokenAmount.GT(sdk.ZeroInt()) {
		willSendCoin.Add(sdk.NewCoin(msg.Denom, rmRTokenAmount))
	}
	if willSendCoin.IsAllPositive() {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddress, willSendCoin); err != nil {
			return nil, err
		}
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(types.LpTokenDenom(msg.Denom), msg.RmUnit))); err != nil {
		return nil, err
	}
	k.SetSwapPool(ctx, msg.Denom, swapPool)

	return &types.MsgRemoveLiquidityResponse{}, nil
}
