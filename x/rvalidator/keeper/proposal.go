package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
)

func (k Keeper) ProcessUpdateRValidatorProposal(ctx sdk.Context, p *types.UpdateRValidatorProposal) error {

	oldVal := types.RValidator{
		Denom:   p.Denom,
		Address: p.OldAddress,
	}
	newVal := types.RValidator{
		Denom:   p.Denom,
		Address: p.NewAddress,
	}
	if !k.HasSelectedRValidator(ctx, &oldVal) {
		return types.ErrRValidatorNotExist
	}
	if k.HasSelectedRValidator(ctx, &newVal) {
		return types.ErrRValidatorAlreadyExist
	}

	latestVotedCycle := k.GetLatestVotedCycle(ctx, p.Denom)
	if p.Cycle <= latestVotedCycle {
		return types.ErrCycleBehindLatestCycle
	}

	k.RemoveSelectedRValidator(ctx, &oldVal)
	k.AddSelectedRValidator(ctx, &newVal)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRmRValidator,
			sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
			sdk.NewAttribute(types.AttributeKeyOldAddress, p.OldAddress),
			sdk.NewAttribute(types.AttributeKeyNewAddress, p.NewAddress),
		),
	)

	return nil
}
