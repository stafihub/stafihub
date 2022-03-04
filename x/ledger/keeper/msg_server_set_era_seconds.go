package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetEraSeconds(goCtx context.Context, msg *types.MsgSetEraSeconds) (*types.MsgSetEraSecondsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	rparams, found := k.Keeper.GetRParams(ctx, msg.Denom)
	if !found {
		rparams.Denom = msg.Denom
	}
	rparams.EraSeconds = msg.EraSeconds
	rparams.BondingDuration = msg.BondingDuration
	rparams.Offset = msg.Offset

	k.Keeper.SetRParams(ctx, rparams)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRParamsChanged,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
		),
	)
	return &types.MsgSetEraSecondsResponse{}, nil
}
