package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ledgertypes "github.com/stafihub/stafihub/x/ledger/types"
	"github.com/stafihub/stafihub/x/relayers/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) AddRelayer(goCtx context.Context, msg *types.MsgAddRelayer) (*types.MsgAddRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	if msg.Arena == ledgertypes.ModuleName {
		_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
		if !ok {
			return nil, banktypes.ErrDenomMetadataNotFound
		}
	}

	for _, address := range msg.Addresses {
		_, err := sdk.AccAddressFromBech32(address)
		if err != nil {
			return nil, err
		}
		// Check if the value already exists
		if k.Keeper.HasRelayer(ctx, msg.Arena, msg.Denom, address) {
			return nil, types.ErrRelayerAlreadySet
		}

		k.Keeper.AddRelayer(ctx, msg.Arena, msg.Denom, address)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeRelayerAdded,
				sdk.NewAttribute(types.AttributeKeyArena, msg.Arena),
				sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
				sdk.NewAttribute(types.AttributeKeyRelayer, address),
			),
		)
	}
	return &types.MsgAddRelayerResponse{}, nil
}

func (k msgServer) DeleteRelayer(goCtx context.Context, msg *types.MsgDeleteRelayer) (*types.MsgDeleteRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	if !k.Keeper.HasRelayer(ctx, msg.Arena, msg.Denom, msg.Address) {
		return nil, types.ErrRelayerNotFound
	}

	k.RemoveRelayer(ctx, msg.Arena, msg.Denom, msg.Address)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRelayerRemoved,
			sdk.NewAttribute(types.AttributeKeyArena, msg.Arena),
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyRelayer, msg.Address),
		),
	)
	return &types.MsgDeleteRelayerResponse{}, nil
}

func (k msgServer) SetThreshold(goCtx context.Context, msg *types.MsgSetThreshold) (*types.MsgSetThresholdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	if msg.Arena == ledgertypes.ModuleName {
		_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
		if !ok {
			return nil, banktypes.ErrDenomMetadataNotFound
		}
	}

	relayers := k.GetRelayer(ctx, msg.Arena, msg.Denom)
	if int(msg.Value) > len(relayers) {
		return nil, types.ErrThresholdTooHigh
	}

	lastTh := uint32(0)
	if last, ok := k.GetThreshold(ctx, msg.Arena, msg.Denom); ok {
		lastTh = last.Value
	}

	var threshold = types.Threshold{Arena: msg.Arena, Denom: msg.Denom, Value: msg.Value}
	k.Keeper.SetThreshold(ctx, threshold)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeThresholdUpdated,
			sdk.NewAttribute(types.AttributeKeyLastThreshold, strconv.FormatUint(uint64(lastTh), 10)),
			sdk.NewAttribute(types.AttributeKeyCurrentThreshold, strconv.FormatUint(uint64(msg.Value), 10)),
		),
	)
	return &types.MsgSetThresholdResponse{}, nil
}
