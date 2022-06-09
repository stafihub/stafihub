package keeper

import (
	"fmt"

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
	cycleSeconds := k.GetCycleSeconds(ctx, p.Denom)
	if cycleSeconds.Version != p.Cycle.Version {
		return types.ErrCycleVersionNotMatch
	}

	latestVotedCycle := k.GetLatestVotedCycle(ctx, p.Denom)
	if !(p.Cycle.Version > latestVotedCycle.Version || (p.Cycle.Version == latestVotedCycle.Version && p.Cycle.Number > latestVotedCycle.Number)) {
		return types.ErrCycleBehindLatestCycle
	}

	k.RemoveSelectedRValidator(ctx, &oldVal)
	k.AddSelectedRValidator(ctx, &newVal)
	k.SetLatestVotedCycle(ctx, p.Cycle)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateRValidator,
			sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
			sdk.NewAttribute(types.AttributeKeyOldAddress, p.OldAddress),
			sdk.NewAttribute(types.AttributeKeyNewAddress, p.NewAddress),
			sdk.NewAttribute(types.AttributeKeyCycleVersion, fmt.Sprintf("%d", p.Cycle.Version)),
			sdk.NewAttribute(types.AttributeKeyCycleNumber, fmt.Sprintf("%d", p.Cycle.Number)),
		),
	)

	return nil
}
