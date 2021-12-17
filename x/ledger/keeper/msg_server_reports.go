package keeper

import (
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
			k.Keeper.AddSnapShot(ctx, shotId, bondShot)
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
