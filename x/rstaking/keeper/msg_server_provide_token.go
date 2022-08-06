package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/utils"
	"github.com/stafihub/stafihub/x/rstaking/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) ProvideToken(goCtx context.Context, msg *types.MsgProvideToken) (*types.MsgProvideTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	if msg.Token.Denom != utils.FisDenom {
		return nil, types.ErrDenomUnmatch
	}
	if err := k.Keeper.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(msg.Token)); err != nil {
		return nil, err
	}

	return &types.MsgProvideTokenResponse{}, nil
}
