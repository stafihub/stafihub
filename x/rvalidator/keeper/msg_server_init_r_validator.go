package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

// init rvalidator and can only init once
func (k msgServer) InitRValidator(goCtx context.Context, msg *types.MsgInitRValidator) (*types.MsgInitRValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	rvalidatorList := k.Keeper.GetSelectedRValidatorListByDenomPoolAddress(ctx, msg.Denom, msg.PoolAddress)
	if len(rvalidatorList) > 0 {
		return nil, types.ErrRValidatorAlreadyInit
	}

	if err := k.RBankKeeper.CheckAccAddress(ctx, msg.Denom, msg.PoolAddress); err != nil {
		return nil, err
	}

	addresses := ""
	for _, address := range msg.ValAddressList {
		if err := k.RBankKeeper.CheckValAddress(ctx, msg.Denom, address); err != nil {
			return nil, err
		}
		rValidator := types.RValidator{
			Denom:       msg.Denom,
			PoolAddress: msg.PoolAddress,
			ValAddress:  address,
		}

		if k.Keeper.HasSelectedRValidator(ctx, &rValidator) {
			return nil, types.ErrRValidatorAlreadyExist
		}

		k.Keeper.AddSelectedRValidator(ctx, &rValidator)

		addresses = addresses + ":" + address
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeInitRValidator,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyPoolAddress, msg.PoolAddress),
			sdk.NewAttribute(types.AttributeKeyAddresses, addresses[1:]),
		),
	)
	return &types.MsgInitRValidatorResponse{}, nil
}
