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

func (k msgServer) CreateRelayer(goCtx context.Context, msg *types.MsgCreateRelayer) (*types.MsgCreateRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	if msg.Taipe == ledgertypes.ModuleName {
		_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
		if !ok {
			return nil, banktypes.ErrDenomMetadataNotFound
		}
	}


	for _, address := range msg.Addresses {
		// Check if the value already exists
		if k.Keeper.HasRelayer(ctx, msg.Taipe, msg.Denom, address) {
			return nil, types.ErrRelayerAlreadySet
		}

		k.AddRelayer(ctx, msg.Taipe, msg.Denom, address)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeRelayerAdded,
				sdk.NewAttribute(types.AttributeKeyTaipe, msg.Taipe),
				sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
				sdk.NewAttribute(types.AttributeKeyRelayer, address),
			),
		)
	}
	return &types.MsgCreateRelayerResponse{}, nil
}

func (k msgServer) DeleteRelayer(goCtx context.Context, msg *types.MsgDeleteRelayer) (*types.MsgDeleteRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	if !k.Keeper.HasRelayer(ctx, msg.Taipe, msg.Denom, msg.Address) {
		return nil, types.ErrRelayerNotFound
	}

	k.RemoveRelayer(ctx, msg.Taipe, msg.Denom, msg.Address)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRelayerRemoved,
			sdk.NewAttribute(types.AttributeKeyTaipe, msg.Taipe),
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

	if msg.Taipe == ledgertypes.ModuleName {
		_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
		if !ok {
			return nil, banktypes.ErrDenomMetadataNotFound
		}
	}

	last, _ := k.GetThreshold(ctx, msg.Taipe, msg.Denom)
	k.SetThreshold(ctx, msg.Taipe, msg.Denom, msg.Value)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeThresholdUpdated,
			sdk.NewAttribute(types.AttributeKeyLastThreshold, strconv.FormatUint(uint64(last), 10)),
			sdk.NewAttribute(types.AttributeKeyCurrentThreshold, strconv.FormatUint(uint64(msg.Value), 10)),
		),
	)
	return &types.MsgUpdateThresholdResponse{}, nil
}
