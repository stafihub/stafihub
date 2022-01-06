package keeper

import (
	"bytes"
	"encoding/hex"
	"strconv"

	"github.com/stafiprotocol/stafihub/x/ledger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)


func (k Keeper) ProcessSetChainEraProposal(ctx sdk.Context,  p *types.SetChainEraProposal) error {
	eraShot := k.CurrentEraSnapShots(ctx, p.Denom)
	if len(eraShot.ShotIds) != 0 {
		return types.ErrEraNotContinuable
	}

	lv, ok := k.relayerKeeper.LastVoter(ctx, p.Denom)
	if !ok {
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
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeEraPoolUpdated,
					sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
					sdk.NewAttribute(types.AttributeKeyLastEra, strconv.FormatUint(uint64(ce.Era), 10)),
					sdk.NewAttribute(types.AttributeKeyCurrentEra, strconv.FormatUint(uint64(p.Era), 10)),
					sdk.NewAttribute(types.AttributeKeyShotId, hex.EncodeToString(shotId)),
					sdk.NewAttribute(types.AttributeKeyLastVoter, lv.Voter),
				),
			)
		}
	}

	k.SetEraSnapShot(ctx, p.Era, eraShot)
	k.SetCurrentEraSnapShot(ctx, eraShot)
	k.SetChainEra(ctx, p.Denom, p.Era)

	return nil
}

func (k Keeper) ProcessBondReportProposal(ctx sdk.Context, p *types.BondReportProposal) error {
	shot, ok := k.SnapShot(ctx, p.ShotId)
	if !ok {
		return types.ErrSnapShotNotFound
	}

	if shot.BondState != types.EraUpdated {
		return types.ErrStateNotEraUpdated
	}

	lv, ok := k.relayerKeeper.LastVoter(ctx, p.Denom)
	if !ok {
		return types.ErrLastVoterNobody
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
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBondReported,
			sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
			sdk.NewAttribute(types.AttributeKeyShotId, hex.EncodeToString(p.ShotId)),
			sdk.NewAttribute(types.AttributeKeyLastVoter, lv.Voter),
		),
	)

	return nil
}

func (k Keeper) ProcessBondAndReportActiveProposal(ctx sdk.Context,  p *types.BondAndReportActiveProposal) error {
	shot, ok := k.SnapShot(ctx, p.ShotId)
	if !ok {
		return types.ErrSnapShotNotFound
	}

	if shot.BondState != types.EraUpdated {
		return types.ErrStateNotEraUpdated
	}

	lv, ok := k.relayerKeeper.LastVoter(ctx, p.Denom)
	if !ok {
		return types.ErrLastVoterNobody
	}

	_, ok = k.GetExchangeRate(ctx, shot.Denom)
	if !ok {
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
		rfee := k.TokenToRtoken(ctx, shot.Denom, fee)
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
		k.SetExchangeRate(ctx, shot.Denom, totalExpectedActive, rtotal.Amount)
	}

	k.SetEraSnapShot(ctx, shot.Era, types.EraSnapShot{Denom: shot.Denom, ShotIds: shots})
	k.SetBondPipeline(ctx, pipe)
	k.SetTotalExpectedActive(ctx, shot.Denom, shot.Era, totalExpectedActive)

	_, ok = k.PoolUnbond(ctx, shot.Denom, shot.Pool, shot.Era)
	if ok {
		shot.UpdateState(types.ActiveReported)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeActiveReported,
				sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
				sdk.NewAttribute(types.AttributeKeyShotId, hex.EncodeToString(p.ShotId)),
				sdk.NewAttribute(types.AttributeKeyLastVoter, lv.Voter),
			),
		)
	} else {
		shot.UpdateState(types.WithdrawSkipped)
		k.SetCurrentEraSnapShot(ctx, newCurrentEraShots)
	}
	k.SetSnapShot(ctx, p.ShotId, shot)

	return nil
}

func (k Keeper) ProcessActiveReportProposal(ctx sdk.Context, p *types.ActiveReportProposal) error {
	shot, ok := k.SnapShot(ctx, p.ShotId)
	if !ok {
		return types.ErrSnapShotNotFound
	}

	if shot.BondState != types.BondReported {
		return types.ErrStateNotBondReported
	}

	lv, ok := k.relayerKeeper.LastVoter(ctx, p.Denom)
	if !ok {
		return types.ErrLastVoterNobody
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
		rfee := k.TokenToRtoken(ctx, shot.Denom, fee)
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
		k.SetExchangeRate(ctx, shot.Denom, totalExpectedActive, rtotal.Amount)
	}

	k.SetEraSnapShot(ctx, shot.Era, types.EraSnapShot{Denom: shot.Denom, ShotIds: shots})
	k.SetBondPipeline(ctx, pipe)
	k.SetTotalExpectedActive(ctx, shot.Denom, shot.Era, totalExpectedActive)

	_, ok = k.PoolUnbond(ctx, shot.Denom, shot.Pool, shot.Era)
	if ok {
		shot.UpdateState(types.ActiveReported)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeActiveReported,
				sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
				sdk.NewAttribute(types.AttributeKeyShotId, hex.EncodeToString(p.ShotId)),
				sdk.NewAttribute(types.AttributeKeyLastVoter, lv.Voter),
			),
		)
	} else {
		shot.UpdateState(types.WithdrawSkipped)
		k.SetCurrentEraSnapShot(ctx, newCurrentEraShots)
	}
	k.SetSnapShot(ctx, p.ShotId, shot)

	return nil
}

func (k Keeper) ProcessWithdrawReportProposal(ctx sdk.Context, p *types.WithdrawReportProposal) error {
	shot, ok := k.SnapShot(ctx, p.ShotId)
	if !ok {
		return types.ErrSnapShotNotFound
	}

	if shot.BondState != types.ActiveReported {
		return types.ErrStateNotActiveReported
	}

	lv, ok := k.relayerKeeper.LastVoter(ctx, p.Denom)
	if !ok {
		return types.ErrLastVoterNobody
	}

	shot.UpdateState(types.WithdrawReported)
	k.SetSnapShot(ctx, p.ShotId, shot)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeWithdrawReported,
			sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
			sdk.NewAttribute(types.AttributeKeyShotId, hex.EncodeToString(p.ShotId)),
			sdk.NewAttribute(types.AttributeKeyLastVoter, lv.Voter),
		),
	)

	return nil
}

func (k Keeper) ProcessTransferReportProposal(ctx sdk.Context,  p *types.TransferReportProposal) error {
	shot, ok := k.SnapShot(ctx, p.ShotId)
	if !ok {
		return types.ErrSnapShotNotFound
	}

	if shot.BondState != types.ActiveReported && shot.BondState != types.WithdrawReported {
		return types.ErrStateNotTransferable
	}

	lv, ok := k.relayerKeeper.LastVoter(ctx, p.Denom)
	if !ok {
		return types.ErrLastVoterNobody
	}

	currentEraShots := k.CurrentEraSnapShots(ctx, shot.Denom)
	newCurrentEraShots := types.NewEraSnapShot(shot.Denom)
	found := false
	for _, id := range currentEraShots.ShotIds {
		if bytes.Equal(id, p.ShotId) {
			found = true
		} else {
			newCurrentEraShots.ShotIds = append(newCurrentEraShots.ShotIds, id)
		}
	}
	if !found {
		return types.ErrTransferReported
	}

	shot.UpdateState(types.TransferReported)
	k.SetSnapShot(ctx, p.ShotId, shot)
	k.SetCurrentEraSnapShot(ctx, newCurrentEraShots)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTransferReported,
			sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
			sdk.NewAttribute(types.AttributeKeyShotId, hex.EncodeToString(p.ShotId)),
			sdk.NewAttribute(types.AttributeKeyLastVoter, lv.Voter),
		),
	)

	return nil
}