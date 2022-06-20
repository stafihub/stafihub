package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) RmRValidator(goCtx context.Context, msg *types.MsgRmRValidator) (*types.MsgRmRValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}
	rValidator := types.RValidator{
		Denom:       msg.Denom,
		PoolAddress: msg.PoolAddress,
		ValAddress:  msg.ValAddress,
	}

	if !k.Keeper.HasSelectedRValidator(ctx, &rValidator) {
		return nil, types.ErrRValidatorNotExist
	}

	k.Keeper.RemoveSelectedRValidator(ctx, &rValidator)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRmRValidator,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyPoolAddress, msg.PoolAddress),
			sdk.NewAttribute(types.AttributeKeyAddress, msg.ValAddress),
		),
	)
	return &types.MsgRmRValidatorResponse{}, nil
}
