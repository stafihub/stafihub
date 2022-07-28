package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/claim/types"
)

func (k msgServer) ProvideToken(goCtx context.Context, msg *types.MsgProvideToken) (*types.MsgProvideTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if err := k.Keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, sdk.NewCoins(msg.Token)); err != nil {
		return nil, err
	}

	return &types.MsgProvideTokenResponse{}, nil
}
