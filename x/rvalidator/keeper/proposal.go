package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
)

func (k Keeper) ProcessUpdateRValidatorProposal(ctx sdk.Context, p *types.UpdateRValidatorProposal) error {

	oldVal := types.RValidator{
		Denom:       p.Denom,
		PoolAddress: p.PoolAddress,
		ValAddress:  p.OldAddress,
	}
	newVal := types.RValidator{
		Denom:       p.Denom,
		PoolAddress: p.PoolAddress,
		ValAddress:  p.NewAddress,
	}
	if !k.HasSelectedRValidator(ctx, &oldVal) {
		return types.ErrRValidatorNotExist
	}
	if k.HasSelectedRValidator(ctx, &newVal) {
		return types.ErrRValidatorAlreadyExist
	}
	if err := k.rBankKeeper.CheckValAddress(ctx, p.Denom, p.NewAddress); err != nil {
		return err
	}
	cycleSeconds := k.GetCycleSeconds(ctx, p.Denom)
	if cycleSeconds.Version != p.Cycle.Version {
		return types.ErrCycleVersionNotMatch
	}

	latestVotedCycle := k.GetLatestVotedCycle(ctx, p.Denom, p.PoolAddress)
	if !(p.Cycle.Version > latestVotedCycle.Version || (p.Cycle.Version == latestVotedCycle.Version && p.Cycle.Number > latestVotedCycle.Number)) {
		return types.ErrCycleBehindLatestCycle
	}
	latestDealedCycle, found := k.GetLatestDealedCycle(ctx, p.Denom, p.PoolAddress)
	if found && (latestDealedCycle.Number != latestVotedCycle.Number || latestDealedCycle.Version != latestVotedCycle.Version) {
		return types.ErrLatestVotedCycleNotDealed
	}
	snapShots := k.ledgerKeeper.CurrentEraSnapshots(ctx, p.Denom)
	if len(snapShots.ShotIds) > 0 {
		return types.ErrLedgerIsBusyWithEra
	}
	chainEra, found := k.ledgerKeeper.GetChainEra(ctx, p.Denom)
	if !found {
		return types.ErrLedgerChainEraNotExist
	}

	k.RemoveSelectedRValidator(ctx, &oldVal)
	k.AddSelectedRValidator(ctx, &newVal)
	k.SetLatestVotedCycle(ctx, p.Cycle)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateRValidator,
			sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
			sdk.NewAttribute(types.AttributeKeyPoolAddress, p.PoolAddress),
			sdk.NewAttribute(types.AttributeKeyChainEra, fmt.Sprintf("%d", chainEra.Era)),
			sdk.NewAttribute(types.AttributeKeyOldAddress, p.OldAddress),
			sdk.NewAttribute(types.AttributeKeyNewAddress, p.NewAddress),
			sdk.NewAttribute(types.AttributeKeyCycleVersion, fmt.Sprintf("%d", p.Cycle.Version)),
			sdk.NewAttribute(types.AttributeKeyCycleNumber, fmt.Sprintf("%d", p.Cycle.Number)),
			sdk.NewAttribute(types.AttributeKeyCycleSeconds, fmt.Sprintf("%d", cycleSeconds.Seconds)),
		),
	)

	return nil
}

func (k Keeper) ProcessUpdateRValidatorReportProposal(ctx sdk.Context, p *types.UpdateRValidatorReportProposal) error {
	latestVotedCycle := k.GetLatestVotedCycle(ctx, p.Denom, p.PoolAddress)
	if !(p.Cycle.Version == latestVotedCycle.Version && p.Cycle.Number == latestVotedCycle.Number) {
		return types.ErrReportCycleNotMatchLatestVotedCycle
	}

	k.SetLatestDealedCycle(ctx, p.Cycle)
	return nil
}
