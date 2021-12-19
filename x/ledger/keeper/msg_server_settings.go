package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
	sudoTypes "github.com/stafiprotocol/stafihub/x/sudo/types"
)

func (k msgServer) AddNewPool(goCtx context.Context,  msg *types.MsgAddNewPool) (*types.MsgAddNewPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	if err := k.Keeper.AddPool(ctx, msg.Denom, msg.Addr); err != nil {
		return nil, err
	}

	return &types.MsgAddNewPoolResponse{}, nil
}

func (k msgServer) RemovePool(goCtx context.Context,  msg *types.MsgRemovePool) (*types.MsgRemovePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	pool, ok := k.Keeper.TryFindPool(ctx, msg.Denom, msg.Addr, types.PoolPrefix)
	if !ok {
		return nil, types.ErrPoolNotFound
	}

    bpool, ok := k.Keeper.TryFindPool(ctx, msg.Denom, msg.Addr, types.BondedPoolPrefix)
    if !ok {
		return nil, types.ErrPoolNotBonded
	}

	pipe, ok := k.Keeper.BondPipeLine(ctx, msg.Denom, msg.Addr)
	if !ok {
		return nil, types.ErrBondPipelineNotFound
	}

	chunk := pipe.Chunk
	if chunk.Bond.Int64() != 0 || chunk.Unbond.Int64() != 0 || chunk.Active.Int64() != 0 {
		return nil, types.ErrActiveAlreadySet
	}

	delete(pool.Addrs, msg.Addr)
	delete(bpool.Addrs, msg.Addr)

	k.Keeper.SetPool(ctx, pool, types.PoolPrefix)
	k.Keeper.SetPool(ctx, bpool, types.BondedPoolPrefix)

	return &types.MsgRemovePoolResponse{}, nil
}

func (k msgServer) SetEraUnbondLimit(goCtx context.Context,  msg *types.MsgSetEraUnbondLimit) (*types.MsgSetEraUnbondLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetEraUnbondLimit(ctx, msg.Denom, msg.Limit)

	return &types.MsgSetEraUnbondLimitResponse{}, nil
}

func (k msgServer) SetInitBond(goCtx context.Context,  msg *types.MsgSetInitBond) (*types.MsgSetInitBondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	rec, _ := sdk.AccAddressFromBech32(msg.Receiver)
	if err := k.Keeper.SetInitBond(ctx, msg.Denom, msg.Pool, msg.Amount, rec); err != nil {
		return nil, err
	}

	return &types.MsgSetInitBondResponse{}, nil
}

func (k msgServer) SetChainBondingDuration(goCtx context.Context,  msg *types.MsgSetChainBondingDuration) (*types.MsgSetChainBondingDurationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetChainBondingDuration(ctx, msg.Denom, msg.Era)
	return &types.MsgSetChainBondingDurationResponse{}, nil
}

func (k msgServer) SetPoolDetail(goCtx context.Context,  msg *types.MsgSetPoolDetail) (*types.MsgSetPoolDetailResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetPoolDetail(ctx, msg.Denom, msg.Pool, msg.SubAccounts, msg.Threshold)

	return &types.MsgSetPoolDetailResponse{}, nil
}

func (k msgServer) SetLeastBond(goCtx context.Context,  msg *types.MsgSetLeastBond) (*types.MsgSetLeastBondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetLeastBond(ctx, msg.Denom, msg.Amount)

	return &types.MsgSetLeastBondResponse{}, nil
}

func (k msgServer) ClearCurrentEraSnapShots(goCtx context.Context,  msg *types.MsgClearCurrentEraSnapShots) (*types.MsgClearCurrentEraSnapShotsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.ClearCurrentEraSnapShots(ctx, msg.Denom)

	return &types.MsgClearCurrentEraSnapShotsResponse{}, nil
}

func (k msgServer) SetCommission(goCtx context.Context,  msg *types.MsgSetCommission) (*types.MsgSetCommissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetCommission(ctx, msg.Commission)
	return &types.MsgSetCommissionResponse{}, nil
}

func (k msgServer) SetReceiver(goCtx context.Context,  msg *types.MsgSetReceiver) (*types.MsgSetReceiverResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	receiver, _ := sdk.AccAddressFromBech32(msg.Receiver)
	k.Keeper.SetReceiver(ctx, receiver)

	return &types.MsgSetReceiverResponse{}, nil
}
