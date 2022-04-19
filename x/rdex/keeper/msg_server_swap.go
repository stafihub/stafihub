package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/rdex/types"
)

func (k msgServer) Swap(goCtx context.Context, msg *types.MsgSwap) (*types.MsgSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	userAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}

	swapPool, found := k.Keeper.GetSwapPool(ctx, msg.Denom)
	if found {
		return nil, types.ErrSwapPoolAlreadyExist
	}

	outAmount, _ := calSwapResult(swapPool.FisBalance, swapPool.RTokenBalance, msg.InputAmount, msg.InputIsFis)
	if outAmount.LTE(sdk.ZeroInt()) {
		return nil, types.ErrSwapAmountTooFew
	}

	if outAmount.LT(msg.MinOutAmount) {
		return nil, types.ErrLessThanMinOutAmount
	}

	if msg.InputIsFis {
		fisBalance := k.bankKeeper.GetBalance(ctx, userAddress, utils.FisDenom)
		if fisBalance.Amount.LT(msg.InputAmount) {
			return nil, types.ErrInsufficientFisBalance
		}
		if swapPool.RTokenBalance.LTE(outAmount) {
			return nil, types.ErrInsufficientRTokenBalance
		}

		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, sdk.NewCoins(sdk.NewCoin(utils.FisDenom, msg.InputAmount))); err != nil {
			return nil, err
		}
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddress, sdk.NewCoins(sdk.NewCoin(msg.Denom, outAmount))); err != nil {
			return nil, err
		}

		swapPool.FisBalance = swapPool.FisBalance.Add(msg.InputAmount)
		swapPool.RTokenBalance = swapPool.RTokenBalance.Sub(outAmount)
	} else {
		rTokenBalance := k.bankKeeper.GetBalance(ctx, userAddress, msg.Denom)
		if rTokenBalance.Amount.LT(msg.InputAmount) {
			return nil, types.ErrInsufficientRTokenBalance
		}
		if swapPool.FisBalance.LTE(outAmount) {
			return nil, types.ErrInsufficientFisBalance
		}

		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, sdk.NewCoins(sdk.NewCoin(msg.Denom, msg.InputAmount))); err != nil {
			return nil, err
		}
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddress, sdk.NewCoins(sdk.NewCoin(utils.FisDenom, outAmount))); err != nil {
			return nil, err
		}

		swapPool.FisBalance = swapPool.FisBalance.Sub(outAmount)
		swapPool.RTokenBalance = swapPool.RTokenBalance.Add(msg.InputAmount)
	}

	k.Keeper.SetSwapPool(ctx, msg.Denom, swapPool)

	return &types.MsgSwapResponse{}, nil
}
