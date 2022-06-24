package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
	sudoTypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) AddRValidator(goCtx context.Context, msg *types.MsgAddRValidator) (*types.MsgAddRValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	if err := k.RBankKeeper.CheckValAddress(ctx, msg.Denom, msg.ValAddress); err != nil {
		return nil, err
	}

	if err := k.RBankKeeper.CheckAccAddress(ctx, msg.Denom, msg.PoolAddress); err != nil {
		return nil, err
	}

	rValidator := types.RValidator{
		Denom:       msg.Denom,
		PoolAddress: msg.PoolAddress,
		ValAddress:  msg.ValAddress,
	}

	if k.Keeper.HasSelectedRValidator(ctx, &rValidator) {
		return nil, types.ErrRValidatorAlreadyExist
	}

	snapShots := k.ledgerKeeper.CurrentEraSnapshots(ctx, msg.Denom)
	if len(snapShots.ShotIds) > 0 {
		return nil, types.ErrLedgerIsBusyWithEra
	}

	chainEra, found := k.ledgerKeeper.GetChainEra(ctx, msg.Denom)
	if !found {
		return nil, types.ErrLedgerChainEraNotExist
	}

	k.Keeper.AddSelectedRValidator(ctx, &rValidator)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddRValidator,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyPoolAddress, msg.PoolAddress),
			sdk.NewAttribute(types.AttributeKeyChainEra, fmt.Sprintf("%d", chainEra.Era)),
			sdk.NewAttribute(types.AttributeKeyAddedAddress, msg.ValAddress),
		),
	)
	return &types.MsgAddRValidatorResponse{}, nil
}
