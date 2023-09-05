package keeper

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	"github.com/tendermint/tendermint/crypto"
)

func (k Keeper) ProcessSetChainEraProposal(ctx sdk.Context, p *types.SetChainEraProposal) error {
	eraShot := k.CurrentEraSnapshots(ctx, p.Denom)
	if len(eraShot.ShotIds) != 0 {
		return types.ErrEraNotContinuable
	}
	ce, ok := k.GetChainEra(ctx, p.Denom)
	if !ok {
		ce = types.NewChainEra(p.Denom)
	}
	if ce.Era != 0 && ce.Era+1 != p.Era {
		return types.ErrEraSkipped
	}

	bpool, found := k.GetBondedPool(ctx, p.Denom)
	if !found || len(bpool.GetAddrs()) == 0 {
		return types.ErrPoolNotBonded
	}
	for _, addr := range bpool.Addrs {
		pipe, _ := k.GetBondPipeline(ctx, p.Denom, addr)
		bondShot := types.NewBondSnapshot(p.Denom, addr, p.Era, pipe.Chunk)
		bshot, err := bondShot.Marshal()
		if err != nil {
			return err
		}

		shotId := hex.EncodeToString(crypto.Sha256(bshot))
		eraShot.ShotIds = append(eraShot.ShotIds, shotId)
		k.SetSnapshot(ctx, shotId, bondShot)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeEraPoolUpdated,
				sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
				sdk.NewAttribute(types.AttributeKeyLastEra, fmt.Sprintf("%d", ce.Era)),
				sdk.NewAttribute(types.AttributeKeyCurrentEra, fmt.Sprintf("%d", p.Era)),
				sdk.NewAttribute(types.AttributeKeyShotId, shotId),
			),
		)
	}

	k.SetEraSnapshot(ctx, p.Era, eraShot)
	k.SetCurrentEraSnapshot(ctx, eraShot)
	k.SetChainEra(ctx, p.Denom, p.Era)

	return nil
}

func (k Keeper) ProcessBondReportProposal(ctx sdk.Context, p *types.BondReportProposal) error {
	shot, ok := k.Snapshot(ctx, p.ShotId)
	if !ok {
		return types.ErrSnapshotNotFound
	}

	if shot.BondState != types.EraUpdated {
		return types.ErrStateNotEraUpdated
	}

	pipe, found := k.GetBondPipeline(ctx, shot.Denom, shot.Pool)
	if !found {
		return types.ErrBondPipelineNotFound
	}
	switch p.Action {
	case types.BondOnly:
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
	k.SetSnapshot(ctx, p.ShotId, shot)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBondReported,
			sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
			sdk.NewAttribute(types.AttributeKeyShotId, p.ShotId),
		),
	)

	return nil
}

func (k Keeper) ProcessActiveReportProposal(ctx sdk.Context, p *types.ActiveReportProposal) error {
	shot, ok := k.Snapshot(ctx, p.ShotId)
	if !ok {
		return types.ErrSnapshotNotFound
	}

	if shot.BondState != types.BondReported {
		return types.ErrStateNotBondReported
	}

	protocolFeeReceiver, foundReceiver := k.GetProtocolFeeReceiver(ctx)
	if !foundReceiver {
		return types.ErrNoProtocolFeeReceiver
	}

	eraShots := k.EraSnapshot(ctx, shot.Denom, shot.Era)
	found := false
	for _, id := range eraShots.ShotIds {
		if id == p.ShotId {
			found = true
		}
	}
	if !found {
		return types.ErrActiveAlreadySet
	}

	currentEraShots := k.CurrentEraSnapshots(ctx, shot.Denom)
	newCurrentEraShots := types.EraSnapshot{Denom: shot.Denom, ShotIds: make([]string, 0)}
	found = false
	for _, id := range currentEraShots.ShotIds {
		if id == p.ShotId {
			found = true
		} else {
			newCurrentEraShots.ShotIds = append(newCurrentEraShots.ShotIds, id)
		}
	}
	if !found {
		return types.ErrActiveAlreadySet
	}

	active := p.Staked.Add(p.Unstaked)
	diff := active.Sub(shot.Chunk.Active)

	totalExpectedFee := k.TotalExpectedFee(ctx, shot.Denom, shot.Era)
	// add commission to totalExpectedFee
	if diff.IsPositive() {
		commission := k.GetStakingRewardCommission(ctx, shot.Denom)
		fee := commission.MulInt(diff).TruncateInt()
		totalExpectedFee = totalExpectedFee.Add(fee)
	}

	pipe, found := k.GetBondPipeline(ctx, shot.Denom, shot.Pool)
	if !found {
		return types.ErrBondPipelineNotFound
	}
	pipe.Chunk.Active = pipe.Chunk.Active.Add(diff)
	pipe.Chunk.Bond = pipe.Chunk.Bond.Add(p.Unstaked)
	totalExpectedActive := k.TotalExpectedActive(ctx, shot.Denom, shot.Era).Add(pipe.Chunk.Active)

	shotIds := make([]string, 0)
	for _, id := range eraShots.ShotIds {
		if id != p.ShotId {
			shotIds = append(shotIds, id)
		}
	}

	// all pool have reported
	if len(shotIds) == 0 {
		// cal and mint rfee
		oldRTotal := k.bankKeeper.GetSupply(ctx, shot.Denom)

		rfee := sdk.ZeroInt()
		subResult := totalExpectedActive.Sub(totalExpectedFee)
		if subResult.IsPositive() {
			rfee = totalExpectedFee.Mul(oldRTotal.Amount).Quo(subResult)
		}

		if rfee.IsPositive() {
			coin := sdk.NewCoin(shot.Denom, rfee)
			if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{coin}); err != nil {
				return err
			}
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, protocolFeeReceiver, sdk.Coins{coin}); err != nil {
				return err
			}
			k.IncreaseTotalProtocolFee(ctx, coin.Denom, coin.Amount)
		}

		// update exchange rate
		newRTotalAmount := oldRTotal.Amount.Add(rfee)

		k.SetExchangeRate(ctx, shot.Denom, totalExpectedActive, newRTotalAmount)
		newRate, _ := k.GetExchangeRate(ctx, shot.Denom)
		k.SetEraExchangeRate(ctx, shot.Denom, shot.Era, newRate.Value)
	}

	k.SetEraSnapshot(ctx, shot.Era, types.EraSnapshot{Denom: shot.Denom, ShotIds: shotIds})
	k.SetBondPipeline(ctx, pipe)
	k.SetTotalExpectedActive(ctx, shot.Denom, shot.Era, totalExpectedActive)
	k.SetTotalExpectedFee(ctx, shot.Denom, shot.Era, totalExpectedFee)

	// if unbondCount > 0: 1 emit event 2 update shot to activeReported 3 not update current era snapshots
	// else: 1 update shot to transfer skipped 2 update current era snapshots
	unbondCount := k.GetPoolUnbondNextSequence(ctx, shot.Denom, shot.Pool, shot.Era)
	if unbondCount > 0 {
		shot.UpdateState(types.ActiveReported)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeActiveReported,
				sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
				sdk.NewAttribute(types.AttributeKeyShotId, p.ShotId),
			),
		)
	} else {
		shot.UpdateState(types.TransferSkipped)
		k.SetCurrentEraSnapshot(ctx, newCurrentEraShots)
	}
	k.SetSnapshot(ctx, p.ShotId, shot)

	return nil
}

func (k Keeper) ProcessTransferReportProposal(ctx sdk.Context, p *types.TransferReportProposal) error {
	shot, ok := k.Snapshot(ctx, p.ShotId)
	if !ok {
		return types.ErrSnapshotNotFound
	}

	if shot.BondState != types.ActiveReported {
		return types.ErrStateNotTransferable
	}

	currentEraShots := k.CurrentEraSnapshots(ctx, shot.Denom)
	newCurrentEraShots := types.NewEraSnapshot(shot.Denom)
	found := false
	for _, id := range currentEraShots.ShotIds {
		if id == p.ShotId {
			found = true
		} else {
			newCurrentEraShots.ShotIds = append(newCurrentEraShots.ShotIds, id)
		}
	}
	if !found {
		return types.ErrTransferReported
	}

	shot.UpdateState(types.TransferReported)
	k.SetSnapshot(ctx, p.ShotId, shot)
	k.SetCurrentEraSnapshot(ctx, newCurrentEraShots)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTransferReported,
			sdk.NewAttribute(types.AttributeKeyDenom, p.Denom),
			sdk.NewAttribute(types.AttributeKeyShotId, p.ShotId),
		),
	)

	return nil
}

func (k Keeper) ProcessExecuteBondProposal(ctx sdk.Context, p *types.ExecuteBondProposal) error {
	err := k.CheckAddress(ctx, p.Denom, p.Pool)
	if err != nil {
		return err
	}

	// check bonded pool
	if !k.IsBondedPoolExist(ctx, p.Denom, p.Pool) {
		return types.ErrPoolNotBonded
	}

	// check pool status
	poolDetail, found := k.GetPoolDetail(ctx, p.Denom, p.Pool)
	if !found {
		return types.ErrPoolDetailNotFound
	}
	if poolDetail.Status != types.Active {
		return types.ErrPoolStatusUnmatch
	}

	var bonder sdk.AccAddress
	if p.State == types.LiquidityBondStateVerifyOk {
		bonder, err = sdk.AccAddressFromBech32(p.Bonder)
		if err != nil {
			return err
		}
	}

	br, found := k.GetBondRecord(ctx, p.Denom, p.Txhash)
	if found && br.State == types.LiquidityBondStateVerifyOk {
		return types.ErrLiquidityBondAlreadyExecuted
	}
	br = types.NewBondRecord(p.Denom, p.Bonder, p.Pool, p.Txhash, p.Amount, p.State)

	if br.State != types.LiquidityBondStateVerifyOk {
		k.SetBondRecord(ctx, br)
		return nil
	}

	pipe, ok := k.GetBondPipeline(ctx, p.Denom, p.Pool)
	if !ok {
		return types.ErrPoolNotBonded
	}
	pipe.Chunk.Active = pipe.Chunk.Active.Add(p.Amount)
	pipe.Chunk.Bond = pipe.Chunk.Bond.Add(p.Amount)

	rbalance := k.TokenToRtoken(ctx, p.Denom, p.Amount)
	if rbalance.GT(sdk.ZeroInt()) {
		rcoins := sdk.Coins{
			sdk.NewCoin(p.Denom, rbalance),
		}
		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, rcoins); err != nil {
			return err
		}

		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bonder, rcoins); err != nil {
			return err
		}
	}

	k.SetBondRecord(ctx, br)
	k.SetBondPipeline(ctx, pipe)
	k.mintrewardKeeper.UpdateUserClaimInfo(ctx, bonder, p.Denom, rbalance, p.Amount)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBondExecuted,
			sdk.NewAttribute(types.AttributeKeyDenom, br.Denom),
			sdk.NewAttribute(types.AttributeKeyBonder, br.Bonder),
			sdk.NewAttribute(types.AttributeKeyPool, br.Pool),
			sdk.NewAttribute(types.AttributeKeyTxhash, br.Txhash),
			sdk.NewAttribute(types.AttributeKeyBalance, br.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyRbalance, rbalance.String()),
		),
	)

	return nil
}

func (k Keeper) ProcessExecuteNativeAndLsmBondProposal(ctx sdk.Context, p *types.ExecuteNativeAndLsmBondProposal) error {
	err := k.CheckAddress(ctx, p.Denom, p.Pool)
	if err != nil {
		return err
	}

	if !k.IsBondedPoolExist(ctx, p.Denom, p.Pool) {
		return types.ErrPoolNotBonded
	}

	icaPool, found := k.GetIcaPoolByDelegationAddr(ctx, p.Pool)
	if !found {
		return types.ErrIcaPoolNotFound
	}

	eraShot := k.CurrentEraSnapshots(ctx, p.Denom)
	if len(eraShot.ShotIds) != 0 {
		return types.ErrEraNotContinuable
	}

	// check pool status
	poolDetail, found := k.GetPoolDetail(ctx, p.Denom, p.Pool)
	if !found {
		return types.ErrPoolDetailNotFound
	}
	if poolDetail.Status != types.Active {
		return types.ErrPoolStatusUnmatch
	}

	var bonder sdk.AccAddress
	if p.State == types.LiquidityBondStateVerifyOk {
		bonder, err = sdk.AccAddressFromBech32(p.Bonder)
		if err != nil {
			return err
		}
	}

	br, found := k.GetBondRecord(ctx, p.Denom, p.Txhash)
	if found && br.State == types.LiquidityBondStateVerifyOk {
		return types.ErrLiquidityBondAlreadyExecuted
	}
	totalBondAmount := p.NativeBondAmount.Add(p.LsmBondAmount)
	br = types.NewBondRecord(p.Denom, p.Bonder, p.Pool, p.Txhash, totalBondAmount, p.State)

	if br.State != types.LiquidityBondStateVerifyOk {
		k.SetBondRecord(ctx, br)
		return nil
	}

	pipe, ok := k.GetBondPipeline(ctx, p.Denom, p.Pool)
	if !ok {
		return types.ErrPoolNotBonded
	}
	pipe.Chunk.Active = pipe.Chunk.Active.Add(totalBondAmount)
	pipe.Chunk.Bond = pipe.Chunk.Bond.Add(p.NativeBondAmount)

	rbalance := k.TokenToRtoken(ctx, p.Denom, totalBondAmount)
	if rbalance.GT(sdk.ZeroInt()) {
		rcoins := sdk.Coins{
			sdk.NewCoin(p.Denom, rbalance),
		}
		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, rcoins); err != nil {
			return err
		}

		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, bonder, rcoins); err != nil {
			return err
		}
	}

	k.SetBondRecord(ctx, br)
	k.SetBondPipeline(ctx, pipe)
	k.mintrewardKeeper.UpdateUserClaimInfo(ctx, bonder, p.Denom, rbalance, totalBondAmount)

	txMsg, err := p.GetTxMsg(k.cdc)
	if err != nil {
		return err
	}
	if len(txMsg) > 0 {
		preProposalId, found := k.GetLatestLsmBondProposalId(ctx)
		if found {
			status, exist := k.GetInterchainTxProposalStatus(ctx, preProposalId)
			if !exist {
				return types.ErrInterchainTxPropIdNotFound
			}
			if status != types.InterchainTxStatusSuccess {
				return types.ErrLsmInterchainTxFailed
			}
		}

		sequence, err := k.SubmitTxs(ctx, icaPool.DelegationAccount.CtrlConnectionId, icaPool.DelegationAccount.Owner, txMsg, types.TxTypeReDeemToken.String())
		if err != nil {
			return err
		}

		k.SetInterchainTxProposalSequenceIndex(ctx, icaPool.DelegationAccount.CtrlPortId, icaPool.DelegationAccount.CtrlChannelId, sequence, p.PropId)
		k.SetInterchainTxProposalStatus(ctx, p.PropId, types.InterchainTxStatusInit)

		k.SetLatestLsmBondProposalId(ctx, p.PropId)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeNativeAndLsmBondExecuted,
			sdk.NewAttribute(types.AttributeKeyDenom, br.Denom),
			sdk.NewAttribute(types.AttributeKeyBonder, br.Bonder),
			sdk.NewAttribute(types.AttributeKeyPool, br.Pool),
			sdk.NewAttribute(types.AttributeKeyTxhash, br.Txhash),
			sdk.NewAttribute(types.AttributeKeyBalance, br.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyNativeBondAmount, p.NativeBondAmount.String()),
			sdk.NewAttribute(types.AttributeKeyLsmBondAmount, p.LsmBondAmount.String()),
			sdk.NewAttribute(types.AttributeKeyRbalance, rbalance.String()),
		),
	)

	return nil
}

func (k Keeper) ProcessInterchainTxProposal(ctx sdk.Context, p *types.InterchainTxProposal) error {
	err := k.CheckAddress(ctx, p.Denom, p.PoolAddress)
	if err != nil {
		return err
	}

	if _, ok := k.GetBondedPool(ctx, p.Denom); !ok {
		return types.ErrPoolNotBonded
	}

	ce, ok := k.GetChainEra(ctx, p.Denom)
	if !ok {
		return types.ErrChainEraNotFound
	}
	if p.Era > ce.Era {
		return types.ErrInvalidEra
	}

	icaPool, found := k.GetIcaPoolByDelegationAddr(ctx, p.PoolAddress)
	if !found {
		return types.ErrIcaPoolNotFound
	}

	txMsg, err := p.GetTxMsg(k.cdc)
	if err != nil {
		return err
	}
	if len(txMsg) == 0 {
		return types.ErrInterchainTxMsgsEmpty
	}

	if p.TxType != types.TxTypeWithdrawAddressSend { //for delegation account
		sequence, err := k.SubmitTxs(ctx, icaPool.DelegationAccount.CtrlConnectionId, icaPool.DelegationAccount.Owner, txMsg, p.TxType.String())
		if err != nil {
			return err
		}

		k.SetInterchainTxProposalSequenceIndex(ctx, icaPool.DelegationAccount.CtrlPortId, icaPool.DelegationAccount.CtrlChannelId, sequence, p.PropId)
	} else { //for withdrawal account
		sequence, err := k.SubmitTxs(ctx, icaPool.WithdrawalAccount.CtrlConnectionId, icaPool.WithdrawalAccount.Owner, txMsg, p.TxType.String())
		if err != nil {
			return err
		}

		k.SetInterchainTxProposalSequenceIndex(ctx, icaPool.WithdrawalAccount.CtrlPortId, icaPool.WithdrawalAccount.CtrlChannelId, sequence, p.PropId)
	}
	k.SetInterchainTxProposalStatus(ctx, p.PropId, types.InterchainTxStatusInit)

	return nil
}
