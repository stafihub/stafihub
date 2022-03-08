package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stafihub/stafihub/utils"
	xBridgeTypes "github.com/stafihub/stafihub/x/bridge/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) MigrateInit(goCtx context.Context, msg *types.MsgMigrateInit) (*types.MsgMigrateInitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	//should return if exist&&exchangeRate != 1
	rate, found := k.GetExchangeRate(ctx, msg.Denom)
	if found && !rate.Value.Equal(utils.OneDec()) {
		return nil, types.ErrExchangeRateAlreadyExist
	}

	k.MigrateExchangeRate(ctx, msg.Denom, msg.ExchangeRate)

	shouldMintCoins := sdk.NewCoins(sdk.NewCoin(msg.Denom, msg.TotalSupply))
	moduleAddress := authTypes.NewModuleAddress(xBridgeTypes.ModuleName)
	balance := k.bankKeeper.GetBalance(ctx, moduleAddress, msg.Denom)
	if balance.Amount.GT(sdk.ZeroInt()) {
		k.bankKeeper.BurnCoins(ctx, xBridgeTypes.ModuleName, sdk.NewCoins(balance))
	}
	k.bankKeeper.MintCoins(ctx, xBridgeTypes.ModuleName, shouldMintCoins)

	return &types.MsgMigrateInitResponse{}, nil
}
