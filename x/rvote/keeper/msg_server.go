package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ledgertypes "github.com/stafihub/stafihub/x/ledger/types"
	relayerstypes "github.com/stafihub/stafihub/x/relayers/types"
	rvalidatortypes "github.com/stafihub/stafihub/x/rvalidator/types"
	"github.com/stafihub/stafihub/x/rvote/types"
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

func (k msgServer) SetProposalLife(goCtx context.Context, msg *types.MsgSetProposalLife) (*types.MsgSetProposalLifeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	k.Keeper.SetProposalLife(ctx, msg.ProposalLife)

	return &types.MsgSetProposalLifeResponse{}, nil
}

func (k msgServer) SubmitProposal(goCtx context.Context, msg *types.MsgSubmitProposal) (*types.MsgSubmitProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	content := msg.GetContent()
	arena := content.ProposalRoute()

	// will use ledger relayers in this case
	if content.ProposalRoute() == rvalidatortypes.ModuleName && content.ProposalType() == rvalidatortypes.TypeUpdateRValidatorReportProposal {
		arena = ledgertypes.ModuleName
	}

	relayerFlag := k.relayerKeeper.HasRelayer(ctx, arena, content.GetDenom(), msg.Proposer)

	if !relayerFlag {
		return nil, types.ErrInvalidProposer
	}

	prop, err := k.PreSubmitProposal(ctx, content, msg.Proposer)
	if err != nil {
		return nil, err
	}
	if prop.Status == types.StatusApproved {
		return nil, types.ErrProposalAlreadyApproved
	}
	if prop.IsExpired(ctx.BlockHeight()) {
		return nil, types.ErrProposalAlreadyExpired
	}

	threshold, ok := k.relayerKeeper.GetThreshold(ctx, arena, content.GetDenom())
	if !ok {
		return nil, relayerstypes.ErrThresholdNotFound
	}

	if uint32(len(prop.Voted)) >= threshold.Value {
		prop.Status = types.StatusApproved
	}

	res := &types.MsgSubmitProposalResponse{PropId: prop.PropId(), Status: prop.Status}
	if prop.Status != types.StatusApproved {
		k.SetProposal(ctx, prop)
		return res, nil
	}

	rtr := k.Keeper.Router()
	handler := rtr.GetRoute(prop.ProposalRoute())
	cacheCtx, writeCache := ctx.CacheContext()
	if err := handler(cacheCtx, prop.GetContent()); err != nil {
		return nil, err
	}
	k.SetProposal(ctx, prop)
	// The cached context is created with a new EventManager. However, since
	// the proposal handler execution was successful, we want to track/keep
	// any events emitted, so we re-emit to "merge" the events into the
	// original Context's EventManager.
	ctx.EventManager().EmitEvents(cacheCtx.EventManager().Events())

	// write state to the underlying multi-store
	writeCache()
	return res, nil
}

func (k msgServer) PreSubmitProposal(ctx sdk.Context, content types.Content, proposer string) (*types.Proposal, error) {
	propId := content.GetPropId()
	prop, ok := k.GetProposal(ctx, propId)
	if !ok {
		prop = &types.Proposal{
			Status:     types.StatusInitiated,
			StartBlock: ctx.BlockHeight(),
			Voted:      []string{proposer},
		}
		prop.ExpireBlock = prop.StartBlock + k.ProposalLife(ctx)
		if err := prop.SetContent(content); err != nil {
			return nil, err
		}
	} else {
		if prop.HasVoted(proposer) {
			return nil, relayerstypes.ErrAlreadyVoted
		}
		prop.Voted = append(prop.Voted, proposer)
	}
	return prop, nil
}
