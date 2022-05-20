package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rdex/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) CreatePool(goCtx context.Context, msg *types.MsgCreatePool) (*types.MsgCreatePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	userAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}
	orderTokens := sdk.Coins{msg.Token0, msg.Token1}.Sort()
	lpDenom := types.GetLpTokenDenom(orderTokens)

	// check swap pool
	_, found := k.Keeper.GetSwapPool(ctx, lpDenom)
	if found {
		return nil, types.ErrSwapPoolAlreadyExist
	}
	// check balance
	for _, token := range orderTokens {
		balance := k.bankKeeper.GetBalance(ctx, userAddress, token.Denom)
		if balance.Amount.LT(token.Amount) {
			return nil, types.ErrUserTokenBalanceInsufficient
		}
	}
	poolTotalUnit, addLpUnit := CalPoolUnit(sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), orderTokens[0].Amount, orderTokens[1].Amount)
	if !addLpUnit.IsPositive() {
		return nil, types.ErrAddLpUnitZero
	}

	// send coins
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, orderTokens); err != nil {
		return nil, err
	}
	lpTokenCoins := sdk.NewCoins(sdk.NewCoin(lpDenom, addLpUnit))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, lpTokenCoins); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddress, lpTokenCoins); err != nil {
		return nil, err
	}

	swapPool := types.SwapPool{
		LpToken:   sdk.NewCoin(lpDenom, poolTotalUnit),
		BaseToken: orderTokens[0],
		Token:     orderTokens[1],
	}

	k.Keeper.SetSwapPool(ctx, lpDenom, &swapPool)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyLpDenom, lpDenom),
			sdk.NewAttribute(types.AttributeKeyNewTotalUnit, poolTotalUnit.String()),
			sdk.NewAttribute(types.AttributeKeyAddLpUnit, addLpUnit.String()),
			sdk.NewAttribute(types.AttributeKeyPoolBaseTokenBalance, swapPool.BaseToken.String()),
			sdk.NewAttribute(types.AttributeKeyPoolTokenBalance, swapPool.Token.String()),
		),
	)

	return &types.MsgCreatePoolResponse{}, nil
}
