package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)


func (k msgServer) Onboard(goCtx context.Context,  msg *types.MsgOnboard) (*types.MsgOnboardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	addrPfx, found := k.sudoKeeper.GetAddressPrefix(ctx, msg.Denom)
	if !found {
		return nil, sudotypes.ErrAddrPrefixNotExist
	}

	if !strings.HasPrefix(msg.Address, addrPfx) {
		return nil, types.ErrAddrPrefixNotMatched
	}

	if _, found := k.Keeper.RValidator(ctx, msg.Denom, msg.Address, msg.OperatorAddress); found {
		return nil, types.ErrRValidatorAlreadyExist
	}

	indicator, found := k.Keeper.RValidatorIndicator(ctx, msg.Denom)
	if found {
		if msg.Locked.Denom != indicator.Locked.Denom {
			return nil, types.ErrLockedDenomNotMatch
		}

		if !indicator.Locked.IsZero() && msg.Locked.IsLT(indicator.Locked) {
			return nil, types.ErrLockedNotEnough
		}
	}

	rv := types.NewRValidator(msg.Denom, msg.Address, msg.OperatorAddress, msg.Locked)
	k.Keeper.SetRValidator(ctx, rv)

	return &types.MsgOnboardResponse{}, nil
}

func (k msgServer) SetRValidatorIndicator(goCtx context.Context,  msg *types.MsgSetRValidatorIndicator) (*types.MsgSetRValidatorIndicatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	indicator := types.NewRValidatorIndicator(msg.Denom, msg.Commission, msg.Uptime, msg.VotingPower, msg.Locked)
	k.Keeper.SetRValidatorIndicator(ctx, indicator)

	return &types.MsgSetRValidatorIndicatorResponse{}, nil
}


