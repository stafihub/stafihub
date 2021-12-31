package keeper

import (
	"context"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/rvote/types"
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

func (k msgServer) SetProposalLife(goCtx context.Context,  msg *types.MsgSetProposalLife) (*types.MsgSetProposalLifeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetProposalLife(ctx, msg.ProposalLife)

	return &types.MsgSetProposalLifeResponse{}, nil
}

func (k msgServer) SubmitProposal(goCtx context.Context,  msg *types.MsgSubmitProposal) (*types.MsgSubmitProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	content := msg.GetContent()

	adminFlag := k.sudoKeeper.IsAdmin(ctx, msg.Proposer)
	relayerFlag := k.relayerKeeper.IsRelayer(ctx, content.GetDenom(), msg.Proposer)

	if !adminFlag && !relayerFlag {
		return nil, types.ErrInvalidProposer
	}

	prop, err := k.Keeper.SubmitProposal(ctx, content, msg.Proposer)
	if err != nil {
		return nil, err
	}

	res := &types.MsgSubmitProposalResponse{PropId: hex.EncodeToString(prop.PropId()), Status: prop.Status}
	if !adminFlag && prop.Status != types.StatusApproved {
		return res, nil
	}

	rtr := k.Keeper.Router()
	handler := rtr.GetRoute(prop.ProposalRoute())
	cacheCtx, writeCache := ctx.CacheContext()
	if err := handler(cacheCtx, prop.GetContent()); err != nil {
		return nil, err
	}
	// The cached context is created with a new EventManager. However, since
	// the proposal handler execution was successful, we want to track/keep
	// any events emitted, so we re-emit to "merge" the events into the
	// original Context's EventManager.
	ctx.EventManager().EmitEvents(cacheCtx.EventManager().Events())

	// write state to the underlying multi-store
	writeCache()
	return res, nil

	return &types.MsgSubmitProposalResponse{}, nil
}
