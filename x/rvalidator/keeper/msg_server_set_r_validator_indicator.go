package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) SetRValidatorIndicator(goCtx context.Context, msg *types.MsgSetRValidatorIndicator) (*types.MsgSetRValidatorIndicatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	rvi := types.RValidatorIndicator{
		Denom:      msg.Denom,
		Commission: msg.Commission,
		Uptime:     msg.Uptime,
		Locked:     msg.Locked,
	}

	k.Keeper.SetRValidatorIndicator(ctx, rvi)
	return &types.MsgSetRValidatorIndicatorResponse{}, nil
}
