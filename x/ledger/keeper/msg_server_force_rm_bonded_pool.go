package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) ForceRmBondedPool(goCtx context.Context, msg *types.MsgForceRmBondedPool) (*types.MsgForceRmBondedPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	if !k.IsBondedPoolExist(ctx, msg.Denom, msg.Address) {
		return nil, types.ErrPoolNotBonded
	}

	k.Keeper.RemoveBondedPool(ctx, msg.Denom, msg.Address)

	bondPipeLine, found := k.Keeper.GetBondPipeline(ctx, msg.Denom, msg.Address)
	if !found {
		return nil, types.ErrBondPipelineNotFound
	}

	bondPipeLine.Chunk.Active = sdk.ZeroInt()
	bondPipeLine.Chunk.Bond = sdk.ZeroInt()
	bondPipeLine.Chunk.Unbond = sdk.ZeroInt()

	k.Keeper.SetBondPipeline(ctx, bondPipeLine)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRemovePool,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyPool, msg.Address),
		),
	)

	return &types.MsgForceRmBondedPoolResponse{}, nil
}
