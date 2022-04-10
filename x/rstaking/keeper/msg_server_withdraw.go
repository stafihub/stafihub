package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stafihub/stafihub/x/rstaking/types"
)

func (k msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	moduleAddress := authTypes.NewModuleAddress(types.ModuleName)
	balance := k.bankKeeper.GetBalance(ctx, moduleAddress, msg.Amount.Denom)
	if balance.Amount.LT(msg.Amount.Amount) {
		return nil, types.ErrInsufficientFunds
	}
	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, sdk.NewCoins(msg.Amount))
	if err != nil {
		return nil, err
	}
	return &types.MsgWithdrawResponse{}, nil
}
