package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rstaking/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) AddDelegatorToWhitelist(goCtx context.Context, msg *types.MsgAddDelegatorToWhitelist) (*types.MsgAddDelegatorToWhitelistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	delAddress, err := sdk.AccAddressFromBech32(msg.DelAddress)
	if err != nil {
		return nil, err
	}
	if k.Keeper.HasDelegatorAddressInWhitelist(ctx, delAddress) {
		return nil, types.ErrDelegatorAlreadyInWhitelist
	}

	k.AddDelegatorAddressToWhitelist(ctx, delAddress)

	return &types.MsgAddDelegatorToWhitelistResponse{}, nil
}
