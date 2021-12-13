package keeper

import (
	"context"
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/relayers/types"
)


func (k msgServer) SubmitProposal(goCtx context.Context,  msg *types.MsgSubmitProposal) (*types.MsgSubmitProposalResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	content := msg.GetContent()

	if !k.CheckIsRelayer(ctx, content.Denom(), msg.Proposer) {
		return nil, types.ErrProposerNotRelayer
	}

	prop, err := k.Keeper.SubmitProposal(ctx, content, msg.Proposer)
	if err != nil {
		return nil, err
	}



	//prop, err := k.Keeper.SubmitProposal(ctx, msg.Content(), msg.Creator, msg.InFavour)
	//if err != nil {
	//	return nil, err
	//}

	res := &types.MsgSubmitProposalResponse{PropId: hex.EncodeToString(prop.PropId()), Status: prop.Status}
	if prop.Status != types.StatusApproved {
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
}
