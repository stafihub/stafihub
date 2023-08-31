package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) ForceUpdateBondedPool(goCtx context.Context, msg *types.MsgForceUpdateBondedPool) (*types.MsgForceUpdateBondedPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	bondPipeLine, found := k.Keeper.GetBondPipeline(ctx, msg.Denom, msg.Address)
	if !found {
		return nil, types.ErrBondPipelineNotFound
	}

	bondPipeLine.Chunk.Active = msg.Active
	bondPipeLine.Chunk.Bond = msg.Bond
	bondPipeLine.Chunk.Unbond = msg.Unbond

	k.Keeper.SetBondPipeline(ctx, bondPipeLine)

	return &types.MsgForceUpdateBondedPoolResponse{}, nil
}
