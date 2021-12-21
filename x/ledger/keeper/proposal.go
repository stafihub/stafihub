package keeper

import (
	"bytes"

    "github.com/stafiprotocol/stafihub/x/ledger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)


func (k Keeper) ProcessSetChainEraProposal(ctx sdk.Context,  p *types.SetChainEraProposal) error {
	//if !k.IsAdminOrRelayer(ctx, msg.Denom, msg.Creator) {
	//	return types.ErrNeitherRelayerNorAdmin
	//}

	eraShot := k.CurrentEraSnapShots(ctx, p.Denom)
	if len(eraShot.ShotIds) != 0 {
		return types.ErrEraNotContinuable
	}

	if _, ok := k.relayerKeeper.LastVoter(ctx, p.Denom); !ok {
		return types.ErrLastVoterNobody
	}

	ce := k.ChainEra(ctx, p.Denom)
	if ce.Era != 0 && ce.Era + 1 != p.Era {
		return types.ErrEraSkipped
	}

	bpool, ok := k.GetBondedPoolByDenom(ctx, p.Denom)
	if ok {
		for addr, ok1 := range bpool.Addrs {
			if !ok1 {
				continue
			}

			pipe, _ := k.BondPipeLine(ctx, p.Denom, addr)
			bondShot := types.NewBondSnapshot(p.Denom, addr, p.Era, pipe.Chunk, p.Proposer)
			bsnap, err := bondShot.Marshal()
			if err != nil {
				return err
			}

			shotId := crypto.Sha256(bsnap)
			eraShot.ShotIds = append(eraShot.ShotIds, shotId)
			k.SetSnapShot(ctx, shotId, bondShot)
			//k.Keeper.AddCurrentEraSnapShot(ctx, msg.GetDenom, shotId)
			// todo add event
		}
	}

	k.SetEraSnapShot(ctx, p.Era, eraShot)
	k.SetCurrentEraSnapShot(ctx, eraShot)
	k.SetChainEra(ctx, p.Denom, p.Era)
	// todo add event

	return nil
}

func (k Keeper) ProcessBondReportProposal(ctx sdk.Context, p *types.BondReportProposal) error {
	//if !k.IsAdminOrRelayer(ctx, msg.Denom, msg.Creator) {
	//	return nil, types.ErrNeitherRelayerNorAdmin
	//}

	shot, ok := k.SnapShot(ctx, p.ShotId)
	if !ok {
		return types.ErrSnapShotNotFound
	}

	pipe, _ := k.BondPipeLine(ctx, shot.Denom, shot.Pool)
	switch p.Action {
	case types.BondOnly:
		// todo safety check
		pipe.Chunk.Bond = pipe.Chunk.Bond.Sub(shot.Chunk.Bond)
	case types.UnbondOnly:
		pipe.Chunk.Unbond = pipe.Chunk.Unbond.Sub(shot.Chunk.Unbond)
	case types.BothBondUnbond:
		pipe.Chunk.Bond = pipe.Chunk.Bond.Sub(shot.Chunk.Bond)
		pipe.Chunk.Unbond = pipe.Chunk.Unbond.Sub(shot.Chunk.Unbond)
	case types.EitherBondUnbond:
	case types.InterDeduct:
		if shot.Chunk.Bond.GT(shot.Chunk.Unbond) {
			pipe.Chunk.Bond = pipe.Chunk.Bond.Sub(shot.Chunk.Unbond)
			pipe.Chunk.Unbond = sdk.NewInt(0)
		} else {
			pipe.Chunk.Unbond = pipe.Chunk.Unbond.Sub(shot.Chunk.Bond)
			pipe.Chunk.Bond = sdk.NewInt(0)
		}
	}

	k.SetBondPipeline(ctx, pipe)
	shot.UpdateState(types.BondReported)
	k.SetSnapShot(ctx, p.ShotId, shot)
	// todo emit event

	return nil
}

func (k Keeper) ProcessBondAndReportActiveProposal(ctx sdk.Context,  p *types.BondAndReportActiveProposal) error {
	if !k.IsAdminOrRelayer(ctx, p.Denom, p.Proposer) {
		return types.ErrNeitherRelayerNorAdmin
	}

	shot, ok := k.SnapShot(ctx, p.ShotId)
	if !ok {
		return types.ErrSnapShotNotFound
	}

	if shot.BondState != types.EraUpdated {
		return types.ErrStateNotEraUpdated
	}

	if k.rateKeeper.GetRate(ctx, shot.Denom) == nil {
		return types.ErrRateIsNone
	}

	pipe, _ := k.BondPipeLine(ctx, shot.Denom, shot.Pool)
	pipe.Chunk.Bond = pipe.Chunk.Bond.Add(p.Unstaked)
	switch p.Action {
	case types.BondOnly:
		// todo safety check
		pipe.Chunk.Bond = pipe.Chunk.Bond.Sub(shot.Chunk.Bond)
	case types.UnbondOnly:
		pipe.Chunk.Unbond = pipe.Chunk.Unbond.Sub(shot.Chunk.Unbond)
	case types.BothBondUnbond:
		pipe.Chunk.Bond = pipe.Chunk.Bond.Sub(shot.Chunk.Bond)
		pipe.Chunk.Unbond = pipe.Chunk.Unbond.Sub(shot.Chunk.Unbond)
	case types.EitherBondUnbond:
	case types.InterDeduct:
		if shot.Chunk.Bond.GT(shot.Chunk.Unbond) {
			pipe.Chunk.Bond = pipe.Chunk.Bond.Sub(shot.Chunk.Unbond)
			pipe.Chunk.Unbond = sdk.NewInt(0)
		} else {
			pipe.Chunk.Unbond = pipe.Chunk.Unbond.Sub(shot.Chunk.Bond)
			pipe.Chunk.Bond = sdk.NewInt(0)
		}
	}

	receiver := k.Receiver(ctx)
	if receiver == nil {
		return types.ErrNoReceiver
	}

	eraShots := k.EraSnapShot(ctx, shot.Denom, shot.Era)
	found := false
	for _, id := range eraShots.ShotIds {
		if bytes.Equal(id, p.ShotId) {
			found = true
		}
	}
	if !found {
		return types.ErrActiveAlreadySet
	}

	currentEraShots := k.CurrentEraSnapShots(ctx, shot.Denom)
	newCurrentEraShots := types.EraSnapShot {Denom: shot.Denom, ShotIds: make([][]byte, 0)}
	found = false
	for _, id := range currentEraShots.ShotIds {
		if bytes.Equal(id, p.ShotId) {
			found = true
		} else {
			newCurrentEraShots.ShotIds = append(newCurrentEraShots.ShotIds, id)
		}
	}
	if !found {
		return types.ErrActiveAlreadySet
	}

	active := sdk.NewInt(p.Staked.Int64()).Add(p.Unstaked)
	diff := active.Sub(shot.Chunk.Active)
	if diff.GT(sdk.NewInt(0)) {
		commission := k.Commission(ctx)
		fee := commission.MulInt(diff).TruncateInt()
		rfee := k.rateKeeper.TokenToRtoken(ctx, shot.Denom, fee)
		coin := sdk.NewCoin(shot.Denom, rfee)
		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{coin}); err != nil {
			return err
		}

		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, sdk.Coins{coin}); err != nil {
			return err
		}
	}

	pipe.Chunk.Active = pipe.Chunk.Active.Add(diff)
	totalExpectedActive :=  k.TotalExpectedActive(ctx, shot.Denom, shot.Era).Add(pipe.Chunk.Active)

	shots := make([][]byte, 0)
	for _, id := range eraShots.ShotIds {
		if !bytes.Equal(id, p.ShotId) {
			shots = append(shots, id)
		}
	}

	if len(shots) == 0 {
		rtotal := k.bankKeeper.GetSupply(ctx, shot.Denom)
		k.rateKeeper.SetRate(ctx, shot.Denom, totalExpectedActive, rtotal.Amount)
	}

	k.SetEraSnapShot(ctx, shot.Era, types.EraSnapShot{Denom: shot.Denom, ShotIds: shots})
	k.SetBondPipeline(ctx, pipe)
	k.SetTotalExpectedActive(ctx, shot.Denom, shot.Era, totalExpectedActive)

	_, ok = k.PoolUnbond(ctx, shot.Denom, shot.Pool, shot.Era)
	if ok {
		shot.UpdateState(types.ActiveReported)
		// todo event
	} else {
		shot.UpdateState(types.WithdrawSkipped)
		k.SetCurrentEraSnapShot(ctx, newCurrentEraShots)
		// todo event
	}
	k.SetSnapShot(ctx, p.ShotId, shot)

	return nil
}

func (k Keeper) ProcessActiveReportProposal(ctx sdk.Context, p *types.ActiveReportProposal) error {
	//if !k.IsAdminOrRelayer(ctx, p.Denom, p.Creator) {
	//	return nil, types.ErrNeitherRelayerNorAdmin
	//}

	shot, ok := k.SnapShot(ctx, p.ShotId)
	if !ok {
		return types.ErrSnapShotNotFound
	}

	if shot.BondState != types.BondReported {
		return types.ErrStateNotBondReported
	}

	receiver := k.Receiver(ctx)
	if receiver == nil {
		return types.ErrNoReceiver
	}

	eraShots := k.EraSnapShot(ctx, shot.Denom, shot.Era)
	found := false
	for _, id := range eraShots.ShotIds {
		if bytes.Equal(id, p.ShotId) {
			found = true
		}
	}
	if !found {
		return types.ErrActiveAlreadySet
	}

	currentEraShots := k.CurrentEraSnapShots(ctx, shot.Denom)
	newCurrentEraShots := types.EraSnapShot {Denom: shot.Denom, ShotIds: make([][]byte, 0)}
	found = false
	for _, id := range currentEraShots.ShotIds {
		if bytes.Equal(id, p.ShotId) {
			found = true
		} else {
			newCurrentEraShots.ShotIds = append(newCurrentEraShots.ShotIds, id)
		}
	}
	if !found {
		return types.ErrActiveAlreadySet
	}

	active := sdk.NewInt(p.Staked.Int64()).Add(p.Unstaked)
	diff := active.Sub(shot.Chunk.Active)
	if diff.GT(sdk.NewInt(0)) {
		commission := k.Commission(ctx)
		fee := commission.MulInt(diff).TruncateInt()
		rfee := k.rateKeeper.TokenToRtoken(ctx, shot.Denom, fee)
		coin := sdk.NewCoin(shot.Denom, rfee)
		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{coin}); err != nil {
			return err
		}

		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, sdk.Coins{coin}); err != nil {
			return err
		}
	}

	pipe, _ := k.BondPipeLine(ctx, shot.Denom, shot.Pool)
	pipe.Chunk.Active = pipe.Chunk.Active.Add(diff)
	pipe.Chunk.Bond = pipe.Chunk.Bond.Add(p.Unstaked)
	totalExpectedActive :=  k.TotalExpectedActive(ctx, shot.Denom, shot.Era).Add(pipe.Chunk.Active)

	shots := make([][]byte, 0)
	for _, id := range eraShots.ShotIds {
		if !bytes.Equal(id, p.ShotId) {
			shots = append(shots, id)
		}
	}

	if len(shots) == 0 {
		rtotal := k.bankKeeper.GetSupply(ctx, shot.Denom)
		k.rateKeeper.SetRate(ctx, shot.Denom, totalExpectedActive, rtotal.Amount)
	}

	k.SetEraSnapShot(ctx, shot.Era, types.EraSnapShot{Denom: shot.Denom, ShotIds: shots})
	k.SetBondPipeline(ctx, pipe)
	k.SetTotalExpectedActive(ctx, shot.Denom, shot.Era, totalExpectedActive)

	_, ok = k.PoolUnbond(ctx, shot.Denom, shot.Pool, shot.Era)
	if ok {
		shot.UpdateState(types.ActiveReported)
		// todo event
	} else {
		shot.UpdateState(types.WithdrawSkipped)
		k.SetCurrentEraSnapShot(ctx, newCurrentEraShots)
		// todo event
	}
	k.SetSnapShot(ctx, p.ShotId, shot)

	return nil
}

func (k Keeper) ProcessWithdrawReportProposal(ctx sdk.Context, p *types.WithdrawReportProposal) error {
	// TODO: Handling the message
	_ = ctx

	return nil
}

func (k Keeper) ProcessTransferReportProposal(ctx sdk.Context,  p *types.TransferReportProposal) error {
	// TODO: Handling the message
	_ = ctx

	return nil
}