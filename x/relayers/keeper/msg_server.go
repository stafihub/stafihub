package keeper

import (
	"context"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k msgServer) CreateRelayer(goCtx context.Context,  msg *types.MsgCreateRelayer) (*types.MsgCreateRelayerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	// Check if the value already exists
	if k.Keeper.IsRelayer(ctx, msg.Denom, msg.Address) {
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

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	if !k.Keeper.IsRelayer(ctx, msg.Denom, msg.Address) {
		return nil, types.ErrRelayerNotFound
	}

	k.RemoveRelayer(ctx, msg.Denom, msg.Address)
	return &types.MsgDeleteRelayerResponse{}, nil
}

func (k msgServer) UpdateThreshold(goCtx context.Context,  msg *types.MsgUpdateThreshold) (*types.MsgUpdateThresholdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
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



	k.Keeper.SetProposalLife(ctx, msg.ProposalLife)

	return &types.MsgSetProposalLifeResponse{}, nil
}

func (k msgServer) SubmitProposal(goCtx context.Context,  msg *types.MsgSubmitProposal) (*types.MsgSubmitProposalResponse, error) {

}
