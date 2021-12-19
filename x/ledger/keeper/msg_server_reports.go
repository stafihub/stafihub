package keeper

import (
	"bytes"
	"context"

    "github.com/stafiprotocol/stafihub/x/ledger/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)


func (k msgServer) SetChainEra(goCtx context.Context,  msg *types.MsgSetChainEra) (*types.MsgSetChainEraResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsAdminOrRelayer(ctx, msg.Denom, msg.Creator) {
		return nil, types.ErrNeitherRelayerNorAdmin
	}

	eraShot := k.CurrentEraSnapShots(ctx, msg.Denom)
	if len(eraShot.ShotIds) != 0 {
		return nil, types.ErrEraNotContinuable
	}

	if _, ok := k.relayerKeeper.LastVoter(ctx, msg.Denom); !ok {
		return nil, types.ErrLastVoterNobody
	}

	ce := k.Keeper.ChainEra(ctx, msg.Denom)
	if ce.Era != 0 && ce.Era + 1 != msg.Era {
		return nil, types.ErrEraSkipped
	}

	bpool, ok := k.Keeper.GetBondedPoolByDenom(ctx, msg.Denom)
	if ok {
		for addr, ok1 := range bpool.Addrs {
			if !ok1 {
				continue
			}

			pipe, _ := k.Keeper.BondPipeLine(ctx, msg.Denom, addr)
			bondShot := types.NewBondSnapshot(msg.Denom, addr, msg.Era, pipe.Chunk, msg.Creator)
			bsnap, err := bondShot.Marshal()
			if err != nil {
				return nil, err
			}

			shotId := crypto.Sha256(bsnap)
			eraShot.ShotIds = append(eraShot.ShotIds, shotId)
			k.Keeper.SetSnapShot(ctx, shotId, bondShot)
			//k.Keeper.AddCurrentEraSnapShot(ctx, msg.Denom, shotId)
			// todo add event
		}
	}

	k.Keeper.SetEraSnapShot(ctx, msg.Era, eraShot)
	k.Keeper.SetCurrentEraSnapShot(ctx, eraShot)
	k.Keeper.SetChainEra(ctx, msg.Denom, msg.Era)
	// todo add event

	return &types.MsgSetChainEraResponse{}, nil
}

func (k msgServer) ActiveReport(goCtx context.Context,  msg *types.MsgActiveReport) (*types.MsgActiveReportResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsAdminOrRelayer(ctx, msg.Denom, msg.Creator) {
		return nil, types.ErrNeitherRelayerNorAdmin
	}

	shot, ok := k.Keeper.SnapShot(ctx, msg.ShotId)
	if !ok {
		return nil, types.ErrSnapShotNotFound
	}

	if shot.BondState != types.BondReported {
		return nil, types.ErrStateNotBondReported
	}

	receiver := k.Keeper.Receiver(ctx)
	if receiver == nil {
		return nil, types.ErrNoReceiver
	}

	eraShots := k.Keeper.EraSnapShot(ctx, shot.Denom, shot.Era)
	found := false
	for _, id := range eraShots.ShotIds {
		if bytes.Equal(id, msg.ShotId) {
			found = true
		}
	}
	if !found {
		return nil, types.ErrActiveAlreadySet
	}

	currentEraShots := k.Keeper.CurrentEraSnapShots(ctx, shot.Denom)
	newCurrentEraShots := types.EraSnapShot {Denom: shot.Denom, ShotIds: make([][]byte, 0)}
	found = false
	for _, id := range currentEraShots.ShotIds {
		if bytes.Equal(id, msg.ShotId) {
			found = true
		} else {
			newCurrentEraShots.ShotIds = append(newCurrentEraShots.ShotIds, id)
		}
	}
	if !found {
		return nil, types.ErrActiveAlreadySet
	}

	active := sdk.NewInt(msg.Staked.Int64()).Add(msg.Unstaked)
	diff := active.Sub(shot.Chunk.Active)
	if diff.GT(sdk.NewInt(0)) {
		commission := k.Keeper.Commission(ctx)
		fee := commission.MulInt(diff).TruncateInt()
		rfee := k.rateKeeper.TokenToRtoken(ctx, shot.Denom, fee)
		coin := sdk.NewCoin(shot.Denom, rfee)
		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{coin}); err != nil {
			return nil, err
		}

		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, sdk.Coins{coin}); err != nil {
			return nil, err
		}
	}

	pipe, _ := k.Keeper.BondPipeLine(ctx, shot.Denom, shot.Pool)
	pipe.Chunk.Active = pipe.Chunk.Active.Add(diff)
	pipe.Chunk.Bond = pipe.Chunk.Bond.Add(msg.Unstaked)
	totalExpectedActive :=  k.Keeper.TotalExpectedActive(ctx, shot.Denom, shot.Era).Add(pipe.Chunk.Active)

	shots := make([][]byte, 0)
	for _, id := range eraShots.ShotIds {
		if !bytes.Equal(id, msg.ShotId) {
			shots = append(shots, id)
		}
	}

	if len(shots) == 0 {
		rtotal := k.bankKeeper.GetSupply(ctx, shot.Denom)
		k.rateKeeper.SetRate(ctx, shot.Denom, totalExpectedActive, rtotal.Amount)
	}

	k.Keeper.SetEraSnapShot(ctx, shot.Era, types.EraSnapShot{Denom: shot.Denom, ShotIds: shots})
	k.Keeper.SetBondPipeline(ctx, pipe)
	k.Keeper.SetTotalExpectedActive(ctx, shot.Denom, shot.Era, totalExpectedActive)

	_, ok = k.Keeper.PoolUnbond(ctx, shot.Denom, shot.Pool, shot.Era)
	if ok {
		shot.UpdateState(types.ActiveReported)
		// todo event
	} else {
		shot.UpdateState(types.WithdrawSkipped)
		k.Keeper.SetCurrentEraSnapShot(ctx, newCurrentEraShots)
		// todo event
	}
	k.Keeper.SetSnapShot(ctx, msg.ShotId, shot)

	return &types.MsgActiveReportResponse{}, nil
}
