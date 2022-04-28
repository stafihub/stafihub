package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/rdex/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}

	_, found := k.Keeper.GetSwapPool(ctx, msg.Denom)
	if found {
		return nil, types.ErrSwapPoolAlreadyExist
	}
	rTokenBalance := k.bankKeeper.GetBalance(ctx, userAddress, msg.Denom)
	if rTokenBalance.Amount.LT(msg.RTokenAmount) {
		return nil, types.ErrInsufficientRTokenBalance
	}
	fisBalance := k.bankKeeper.GetBalance(ctx, userAddress, utils.FisDenom)
	if fisBalance.Amount.LT(msg.FisAmount) {
		return nil, types.ErrInsufficientFisBalance
	}

	poolTotalUnit, lpUnit := calPoolUnit(sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), msg.FisAmount, msg.RTokenAmount)

	sendCoins := sdk.NewCoins(sdk.NewCoin(msg.Denom, msg.RTokenAmount), sdk.NewCoin(utils.FisDenom, msg.FisAmount))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, sendCoins); err != nil {
		return nil, err
	}

	lpTokenCoins := sdk.NewCoins(sdk.NewCoin(types.LpTokenDenom(msg.Denom), lpUnit))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, lpTokenCoins); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddress, lpTokenCoins); err != nil {
		return nil, err
	}

	swapPool := types.SwapPool{
		Denom:         msg.Denom,
		RTokenBalance: msg.RTokenAmount,
		FisBalance:    msg.FisAmount,
		TotalUnit:     poolTotalUnit,
	}

	k.Keeper.SetSwapPool(ctx, msg.Denom, &swapPool)

	// CreatePool: (account, symbol, fis amount, rToken amount, new total unit, add lp unit)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyFisAmount, msg.FisAmount.String()),
			sdk.NewAttribute(types.AttributeKeyRTokenAmount, msg.RTokenAmount.String()),
			sdk.NewAttribute(types.AttributeKeyNewTotalUnit, poolTotalUnit.String()),
			sdk.NewAttribute(types.AttributeKeyAddLpUnit, lpUnit.String()),
		),
	)

	return &types.MsgCreatePoolResponse{}, nil
}
