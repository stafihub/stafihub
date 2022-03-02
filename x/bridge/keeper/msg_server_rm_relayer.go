package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/bridge/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) RmRelayer(goCtx context.Context, msg *types.MsgRmRelayer) (*types.MsgRmRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}
	accAddress, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	k.Keeper.RmRelayer(ctx, accAddress)

	return &types.MsgRmRelayerResponse{}, nil
}
