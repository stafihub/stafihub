package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetRelayGasPrice(goCtx context.Context, msg *types.MsgSetRelayGasPrice) (*types.MsgSetRelayGasPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	rparams, found := k.Keeper.GetRParams(ctx, msg.Denom)
	if !found {
		rparams.Denom = msg.Denom
	}
	rparams.GasPrice = msg.GasPrice

	k.Keeper.SetRParams(ctx, rparams)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRParamsChanged,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
		),
	)

	return &types.MsgSetRelayGasPriceResponse{}, nil
}
