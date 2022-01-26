package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
	sudoTypes "github.com/stafiprotocol/stafihub/x/sudo/types"
)

func (k msgServer) AddNewPool(goCtx context.Context, msg *types.MsgAddNewPool) (*types.MsgAddNewPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	//if !k.bankKeeper.HasDenomMetaData(ctx, msg.Denom) {
	//	return nil, banktypes.ErrDenomMetadataNotFound
	//}

	k.Keeper.AddPool(ctx, msg.Denom, msg.Addr)
	return &types.MsgAddNewPoolResponse{}, nil
}

func (k msgServer) RemovePool(goCtx context.Context, msg *types.MsgRemovePool) (*types.MsgRemovePoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	if !k.IsPoolExist(ctx, msg.Denom, msg.Addr) {
		return nil, types.ErrPoolNotFound
	}

	if !k.IsBondedPoolExist(ctx, msg.Denom, msg.Addr) {
		return nil, types.ErrPoolNotBonded
	}

	pipe, ok := k.Keeper.GetBondPipeline(ctx, msg.Denom, msg.Addr)
	if !ok {
		return nil, types.ErrBondPipelineNotFound
	}

	chunk := pipe.Chunk
	if chunk.Bond.Int64() != 0 || chunk.Unbond.Int64() != 0 || chunk.Active.Int64() != 0 {
		return nil, types.ErrActiveAlreadySet
	}

	k.Keeper.RemovePool(ctx, msg.Denom, msg.Addr)
	k.Keeper.RemoveBondedPool(ctx, msg.Denom, msg.Addr)

	return &types.MsgRemovePoolResponse{}, nil
}

func (k msgServer) SetEraUnbondLimit(goCtx context.Context, msg *types.MsgSetEraUnbondLimit) (*types.MsgSetEraUnbondLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	k.Keeper.SetEraUnbondLimit(ctx, msg.Denom, msg.Limit)

	return &types.MsgSetEraUnbondLimitResponse{}, nil
}

func (k msgServer) SetInitBond(goCtx context.Context, msg *types.MsgSetInitBond) (*types.MsgSetInitBondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	if !k.IsPoolExist(ctx, msg.Denom, msg.Pool) {
		return nil, types.ErrPoolNotFound
	}

	if k.IsBondedPoolExist(ctx, msg.Denom, msg.Pool) {
		return nil, types.ErrRepeatInitBond
	}

	rbalance := k.TokenToRtoken(ctx, msg.Denom, msg.Amount)
	rcoins := sdk.Coins{
		sdk.NewCoin(msg.Denom, rbalance),
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, rcoins); err != nil {
		panic(err)
	}

	rec, _ := sdk.AccAddressFromBech32(msg.Receiver)
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, rec, rcoins); err != nil {
		panic(err)
	}

	k.SetExchangeRate(ctx, msg.Denom, sdk.NewInt(0), sdk.NewInt(0))
	k.AddBondedPool(ctx, msg.Denom, msg.Pool)
	k.SetBondPipeline(ctx, types.NewBondPipeline(msg.Denom, msg.Pool))

	return &types.MsgSetInitBondResponse{}, nil
}

func (k msgServer) SetChainBondingDuration(goCtx context.Context, msg *types.MsgSetChainBondingDuration) (*types.MsgSetChainBondingDurationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	k.Keeper.SetChainBondingDuration(ctx, msg.Denom, msg.Era)
	return &types.MsgSetChainBondingDurationResponse{}, nil
}

func (k msgServer) SetPoolDetail(goCtx context.Context, msg *types.MsgSetPoolDetail) (*types.MsgSetPoolDetailResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	k.Keeper.SetPoolDetail(ctx, msg.Denom, msg.Pool, msg.SubAccounts, msg.Threshold)

	return &types.MsgSetPoolDetailResponse{}, nil
}

func (k msgServer) SetLeastBond(goCtx context.Context, msg *types.MsgSetLeastBond) (*types.MsgSetLeastBondResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	k.Keeper.SetLeastBond(ctx, msg.Denom, msg.Amount)

	return &types.MsgSetLeastBondResponse{}, nil
}

func (k msgServer) ClearCurrentEraSnapShots(goCtx context.Context, msg *types.MsgClearCurrentEraSnapShots) (*types.MsgClearCurrentEraSnapShotsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	k.Keeper.ClearCurrentEraSnapshots(ctx, msg.Denom)

	return &types.MsgClearCurrentEraSnapShotsResponse{}, nil
}

func (k msgServer) SetCommission(goCtx context.Context, msg *types.MsgSetCommission) (*types.MsgSetCommissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetCommission(ctx, msg.Commission)
	return &types.MsgSetCommissionResponse{}, nil
}

func (k msgServer) SetReceiver(goCtx context.Context, msg *types.MsgSetReceiver) (*types.MsgSetReceiverResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	receiver, _ := sdk.AccAddressFromBech32(msg.Receiver)
	k.Keeper.SetReceiver(ctx, receiver)

	return &types.MsgSetReceiverResponse{}, nil
}

func (k msgServer) SetUnbondFee(goCtx context.Context, msg *types.MsgSetUnbondFee) (*types.MsgSetUnbondFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	if _, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom); !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	if _, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Value.Denom); !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	k.Keeper.SetUnbondFee(ctx, msg.Denom, msg.Value)
	return &types.MsgSetUnbondFeeResponse{}, nil
}

func (k msgServer) SetUnbondCommission(goCtx context.Context, msg *types.MsgSetUnbondCommission) (*types.MsgSetUnbondCommissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetUnbondCommission(ctx, msg.Commission)

	return &types.MsgSetUnbondCommissionResponse{}, nil
}
