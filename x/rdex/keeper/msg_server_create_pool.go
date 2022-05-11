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
	tokens := msg.Tokens.Sort()
	lpDenom := types.GetLpTokenDenom(tokens)

	// check swap pool
	_, found := k.Keeper.GetSwapPool(ctx, lpDenom)
	if found {
		return nil, types.ErrSwapPoolAlreadyExist
	}
	// check balance
	for _, token := range tokens {
		balance := k.bankKeeper.GetBalance(ctx, userAddress, token.Denom)
		if balance.Amount.LT(token.Amount) {
			return nil, types.ErrInsufficientTokenBalance
		}
	}
	poolTotalUnit, lpUnit := calPoolUnit(sdk.ZeroInt(), sdk.ZeroInt(), sdk.ZeroInt(), tokens[0].Amount, tokens[1].Amount)
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, userAddress, types.ModuleName, tokens); err != nil {
		return nil, err
	}

	lpTokenCoins := sdk.NewCoins(sdk.NewCoin(lpDenom, lpUnit))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, lpTokenCoins); err != nil {
		return nil, err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, userAddress, lpTokenCoins); err != nil {
		return nil, err
	}

	swapPool := types.SwapPool{
		LpToken: sdk.NewCoin(lpDenom, lpUnit),
		Tokens:  tokens,
	}

	k.Keeper.SetSwapPool(ctx, lpDenom, &swapPool)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute(types.AttributeKeyAccount, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyLpDenom, lpDenom),
			sdk.NewAttribute(types.AttributeKeyAddTokens, tokens.String()),
			sdk.NewAttribute(types.AttributeKeyNewTotalUnit, poolTotalUnit.String()),
			sdk.NewAttribute(types.AttributeKeyAddLpUnit, lpUnit.String()),
			sdk.NewAttribute(types.AttributeKeyPoolTokensBalance, swapPool.Tokens.String()),
		),
	)

	return &types.MsgCreatePoolResponse{}, nil
}
