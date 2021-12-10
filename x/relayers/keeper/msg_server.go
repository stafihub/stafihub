package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (ms msgServer) IsAdmin(ctx sdk.Context, creator string) bool {
	addr, _ := sdk.AccAddressFromBech32(creator)
	return ms.sudoKeeper.IsAdmin(ctx, addr)
}

func (k msgServer) CreateRelayer(goCtx context.Context,  msg *types.MsgCreateRelayer) (*types.MsgCreateRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsAdmin(ctx, msg.Creator) {
		return nil, types.ErrCreatorNotAdmin
	}

	// Check if the value already exists
	if k.CheckIsRelayer(ctx, msg.Denom, msg.Address) {
		return nil, types.ErrRelayerAlreadySet
	}

	var relayer = &types.Relayer{
		Denom: msg.Denom,
		Address: msg.Address,
	}

	k.SetRelayer(ctx, relayer)
	return &types.MsgCreateRelayerResponse{}, nil
}

func (k msgServer) DeleteRelayer(goCtx context.Context,  msg *types.MsgDeleteRelayer) (*types.MsgDeleteRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsAdmin(ctx, msg.Creator) {
		return nil, types.ErrCreatorNotAdmin
	}

	if !k.CheckIsRelayer(ctx, msg.Denom, msg.Address) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "relayer not exist")
	}

	k.RemoveRelayer(ctx, msg.Denom, msg.Address)
	return &types.MsgDeleteRelayerResponse{}, nil
}

func (k msgServer) UpdateThreshold(goCtx context.Context,  msg *types.MsgUpdateThreshold) (*types.MsgUpdateThresholdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsAdmin(ctx, msg.Creator) {
		return nil, types.ErrCreatorNotAdmin
	}

	var threshold = types.Threshold{
		Denom: msg.Denom,
		Value: msg.Value,
	}

	k.SetThreshold(ctx, &threshold)

	return &types.MsgUpdateThresholdResponse{}, nil
}

func (k msgServer) SetProposalLife(goCtx context.Context,  msg *types.MsgSetProposalLife) (*types.MsgSetProposalLifeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsAdmin(ctx, msg.Creator) {
		return nil, types.ErrCreatorNotAdmin
	}

	k.Keeper.SetProposalLife(ctx, msg.ProposalLife)

	return &types.MsgSetProposalLifeResponse{}, nil
}
