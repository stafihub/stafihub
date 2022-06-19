package keeper

import (
	"context"
	"fmt"

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

func (k msgServer) SetPoolDetail(goCtx context.Context, msg *types.MsgSetPoolDetail) (*types.MsgSetPoolDetailResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}
	err := k.Keeper.CheckAddress(ctx, msg.Denom, msg.Pool)
	if err != nil {
		return nil, err
	}

	for _, subAccount := range msg.SubAccounts {
		err := k.Keeper.CheckAddress(ctx, msg.Denom, subAccount)
		if err != nil {
			return nil, err
		}
	}

	k.Keeper.SetPoolDetail(ctx, msg.Denom, msg.Pool, msg.SubAccounts, msg.Threshold)

	if !k.IsBondedPoolExist(ctx, msg.Denom, msg.Pool) {
		k.SetExchangeRate(ctx, msg.Denom, sdk.NewInt(0), sdk.NewInt(0))
		k.AddBondedPool(ctx, msg.Denom, msg.Pool)
		k.SetBondPipeline(ctx, types.NewBondPipeline(msg.Denom, msg.Pool))
	}

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
	rparams, found := k.Keeper.GetRParams(ctx, msg.Denom)
	if !found {
		rparams.Denom = msg.Denom
	}
	rparams.LeastBond = msg.LeastBond

	k.Keeper.SetRParams(ctx, rparams)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRParamsChanged,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyGasPrice, rparams.GasPrice),
			sdk.NewAttribute(types.AttributeKeyEraSeconds, fmt.Sprintf("%d", rparams.EraSeconds)),
			sdk.NewAttribute(types.AttributeKeyOffset, fmt.Sprintf("%d", rparams.Offset)),
			sdk.NewAttribute(types.AttributeKeyBondingDuration, fmt.Sprintf("%d", rparams.BondingDuration)),
			sdk.NewAttribute(types.AttributeKeyLeastBond, rparams.LeastBond),
		),
	)

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
		Denom:           msg.GetDenom(),
		GasPrice:        msg.GetGasPrice(),
		EraSeconds:      msg.GetEraSeconds(),
		Offset:          msg.Offset,
		BondingDuration: msg.BondingDuration,
		LeastBond:       msg.LeastBond,
	}

	k.Keeper.SetRParams(ctx, rParams)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRParamsChanged,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyGasPrice, rParams.GasPrice),
			sdk.NewAttribute(types.AttributeKeyEraSeconds, fmt.Sprintf("%d", rParams.EraSeconds)),
			sdk.NewAttribute(types.AttributeKeyOffset, fmt.Sprintf("%d", rParams.Offset)),
			sdk.NewAttribute(types.AttributeKeyBondingDuration, fmt.Sprintf("%d", rParams.BondingDuration)),
			sdk.NewAttribute(types.AttributeKeyLeastBond, rParams.LeastBond),
		),
	)
	return &types.MsgSetRParamsResponse{}, nil
}
