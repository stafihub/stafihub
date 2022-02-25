package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetEraUnbondLimit(goCtx context.Context, msg *types.MsgSetEraUnbondLimit) (*types.MsgSetEraUnbondLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
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
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	denom := msg.Coin.Denom
	_, ok := k.bankKeeper.GetDenomMetaData(ctx, denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	if k.IsBondedPoolExist(ctx, denom, msg.Pool) {
		return nil, types.ErrRepeatInitBond
	}

	if msg.Coin.Amount.GT(sdk.ZeroInt()) {
		rbalance := k.TokenToRtoken(ctx, denom, msg.Coin.Amount)
		rcoins := sdk.Coins{sdk.NewCoin(denom, rbalance)}
		if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, rcoins); err != nil {
			panic(err)
		}

		rec, _ := sdk.AccAddressFromBech32(msg.Receiver)
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, rec, rcoins); err != nil {
			panic(err)
		}
	}

	k.SetExchangeRate(ctx, denom, sdk.NewInt(0), sdk.NewInt(0))
	k.AddBondedPool(ctx, denom, msg.Pool)
	k.SetBondPipeline(ctx, types.NewBondPipeline(denom, msg.Pool))

	return &types.MsgSetInitBondResponse{}, nil
}

func (k msgServer) SetChainBondingDuration(goCtx context.Context, msg *types.MsgSetChainBondingDuration) (*types.MsgSetChainBondingDurationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
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
		return nil, sudotypes.ErrCreatorNotAdmin
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
		return nil, sudotypes.ErrCreatorNotAdmin
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
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	k.Keeper.ClearCurrentEraSnapshots(ctx, msg.Denom)

	return &types.MsgClearCurrentEraSnapShotsResponse{}, nil
}

func (k msgServer) SetStakingRewardCommission(goCtx context.Context, msg *types.MsgSetStakingRewardCommission) (*types.MsgSetStakingRewardCommissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetStakingRewardCommission(ctx, msg.Denom, msg.Commission)
	return &types.MsgSetStakingRewardCommissionResponse{}, nil
}

func (k msgServer) SetProtocolFeeReceiver(goCtx context.Context, msg *types.MsgSetProtocolFeeReceiver) (*types.MsgSetProtocolFeeReceiverResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}
	k.Keeper.SetProtocolFeeReceiver(ctx, receiver)

	return &types.MsgSetProtocolFeeReceiverResponse{}, nil
}

func (k msgServer) SetUnbondRelayFee(goCtx context.Context, msg *types.MsgSetUnbondRelayFee) (*types.MsgSetUnbondRelayFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	if _, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom); !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	if _, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Value.Denom); !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	k.Keeper.SetUnbondRelayFee(ctx, msg.Denom, msg.Value)
	return &types.MsgSetUnbondRelayFeeResponse{}, nil
}

func (k msgServer) SetUnbondCommission(goCtx context.Context, msg *types.MsgSetUnbondCommission) (*types.MsgSetUnbondCommissionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetUnbondCommission(ctx, msg.Denom, msg.Commission)

	return &types.MsgSetUnbondCommissionResponse{}, nil
}

func (k msgServer) SetRParams(goCtx context.Context, msg *types.MsgSetRParams) (*types.MsgSetRParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	rParams := types.RParams{
		Creator:     msg.GetCreator(),
		Denom:       msg.GetDenom(),
		ChainId:     msg.GetChainId(),
		NativeDenom: msg.GetNativeDenom(),
		GasPrice:    msg.GetGasPrice(),
		EraSeconds:  msg.GetEraSeconds(),
		LeastBond:   msg.LeastBond,
		Validators:  msg.GetValidators(),
	}

	k.Keeper.SetRParams(ctx, rParams)

	return &types.MsgSetRParamsResponse{}, nil
}
