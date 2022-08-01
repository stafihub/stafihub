package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/claim/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) WithdrawToken(goCtx context.Context, msg *types.MsgWithdrawToken) (*types.MsgWithdrawTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	if err := k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, sdk.NewCoins(msg.Token)); err != nil {
		return nil, err
	}

	return &types.MsgWithdrawTokenResponse{}, nil
}
