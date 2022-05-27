package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rdex/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) AddPoolCreator(goCtx context.Context, msg *types.MsgAddPoolCreator) (*types.MsgAddPoolCreatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	poolCreator, err := sdk.AccAddressFromBech32(msg.UserAddress)
	if err != nil {
		return nil, err
	}
	k.Keeper.AddPoolCreator(ctx, poolCreator)

	return &types.MsgAddPoolCreatorResponse{}, nil
}
