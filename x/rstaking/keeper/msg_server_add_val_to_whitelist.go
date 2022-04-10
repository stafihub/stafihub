package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rstaking/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) AddValToWhitelist(goCtx context.Context, msg *types.MsgAddValToWhitelist) (*types.MsgAddValToWhitelistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	valAddress, err := sdk.ValAddressFromBech32(msg.ValAddress)
	if err != nil {
		return nil, err
	}
	if k.Keeper.HasValAddressInWhitelist(ctx, valAddress) {
		return nil, types.ErrValAlreadyInWhitelist
	}

	k.AddValAddressToWhitelist(ctx, valAddress)

	return &types.MsgAddValToWhitelistResponse{}, nil
}
