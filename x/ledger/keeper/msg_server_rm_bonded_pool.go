package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) RmBondedPool(goCtx context.Context, msg *types.MsgRmBondedPool) (*types.MsgRmBondedPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	if !k.IsBondedPoolExist(ctx, msg.Denom, msg.Address) {
		return nil, types.ErrPoolNotBonded
	}

	bondPipeLine, found := k.Keeper.GetBondPipeline(ctx, msg.Denom, msg.Address)
	if !found {
		return nil, types.ErrBondPipelineNotFound
	}

	// pass: active <= 10000 && bond == 0 && unbond == 0
	if bondPipeLine.Chunk.Active.GT(sdk.NewInt(1e4)) ||
		!bondPipeLine.Chunk.Bond.Equal(sdk.ZeroInt()) ||
		!bondPipeLine.Chunk.Unbond.Equal(sdk.ZeroInt()) {
		return nil, types.ErrBondPipelineAlreadyWork
	}

	k.Keeper.RemoveBondedPool(ctx, msg.Denom, msg.Address)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRemovePool,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyPool, msg.Address),
		),
	)
	return &types.MsgRmBondedPoolResponse{}, nil
}
