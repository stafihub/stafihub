package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
	sudotypes "github.com/stafiprotocol/stafihub/x/sudo/types"
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

func (k msgServer) CreateRelayer(goCtx context.Context, msg *types.MsgCreateRelayer) (*types.MsgCreateRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	// Check if the value already exists
	if k.Keeper.IsRelayer(ctx, msg.Denom, msg.Address) {
		return nil, types.ErrRelayerAlreadySet
	}

	k.AddRelayer(ctx, msg.Denom, msg.Address)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRelayerAdded,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyRelayer, msg.Address),
		),
	)
	return &types.MsgCreateRelayerResponse{}, nil
}

func (k msgServer) DeleteRelayer(goCtx context.Context, msg *types.MsgDeleteRelayer) (*types.MsgDeleteRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	if !k.Keeper.IsRelayer(ctx, msg.Denom, msg.Address) {
		return nil, types.ErrRelayerNotFound
	}

	k.RemoveRelayer(ctx, msg.Denom, msg.Address)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRelayerRemoved,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyRelayer, msg.Address),
		),
	)
	return &types.MsgDeleteRelayerResponse{}, nil
}

func (k msgServer) UpdateThreshold(goCtx context.Context, msg *types.MsgUpdateThreshold) (*types.MsgUpdateThresholdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	lastTh := uint32(0)
	if last, ok := k.GetThreshold(ctx, msg.Denom); ok {
		lastTh = last.Value
	}

	var threshold = types.Threshold{Denom: msg.Denom, Value: msg.Value}
	k.SetThreshold(ctx, threshold)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeThresholdUpdated,
			sdk.NewAttribute(types.AttributeKeyLastThreshold, strconv.FormatUint(uint64(lastTh), 10)),
			sdk.NewAttribute(types.AttributeKeyCurrentThreshold, strconv.FormatUint(uint64(msg.Value), 10)),
		),
	)
	return &types.MsgUpdateThresholdResponse{}, nil
}
