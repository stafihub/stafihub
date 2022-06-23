package keeper

import (
	"context"
	"fmt"

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

	snapShots := k.ledgerKeeper.CurrentEraSnapshots(ctx, msg.Denom)
	if len(snapShots.ShotIds) > 0 {
		return nil, types.ErrLedgerIsBusyWithEra
	}

	chainEra, found := k.ledgerKeeper.GetChainEra(ctx, msg.Denom)
	if !found {
		return nil, types.ErrLedgerChainEraNotExist
	}

	k.Keeper.RemoveSelectedRValidator(ctx, &rValidator)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRmRValidator,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyPoolAddress, msg.PoolAddress),
			sdk.NewAttribute(types.AttributeKeyChainEra, fmt.Sprintf("%d", chainEra.Era)),
			sdk.NewAttribute(types.AttributeKeyAddress, msg.ValAddress),
		),
	)
	return &types.MsgRmRValidatorResponse{}, nil
}
