package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/mining/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) RmMiningProvider(goCtx context.Context, msg *types.MsgRmMiningProvider) (*types.MsgRmMiningProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	userAddr, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		return nil, err
	}

	if !k.Keeper.HasMiningProvider(ctx, userAddr) {
		return nil, types.ErrMiningProviderNotExist
	}

	k.Keeper.RemoveMiningProvider(ctx, userAddr)
	return &types.MsgRmMiningProviderResponse{}, nil
}
