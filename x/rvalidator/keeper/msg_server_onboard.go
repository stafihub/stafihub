package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	rBankTypes "github.com/stafihub/stafihub/x/rbank/types"
	"github.com/stafihub/stafihub/x/rvalidator/types"
)

func (k msgServer) Onboard(goCtx context.Context, msg *types.MsgOnboard) (*types.MsgOnboardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	if err := k.rbankKeeper.CheckValAddress(ctx, msg.Denom, msg.Address); err != nil {
		return nil, err
	}

	rv, found := k.Keeper.GetRValidator(ctx, msg.Denom, msg.Address)
	if found && rv.Status != types.Offboard {
		return nil, types.ErrRValidatorAlreadyOnboard
	}

	ind, found := k.Keeper.GetRValidatorIndicator(ctx, msg.Denom)
	if !found {
		return nil, types.ErrRValidatorIndicatorNotSet
	}

	if msg.Locked.Denom != ind.Locked.Denom {
		return nil, rBankTypes.ErrDenomNotMatched
	}

	if ind.Locked.IsPositive() {
		if msg.Locked.IsLT(ind.Locked) {
			return nil, types.ErrLockedNotEnough
		}

		sender, _ := sdk.AccAddressFromBech32(msg.Creator)
		bal := k.bankKeeper.GetBalance(ctx, sender, msg.Locked.Denom)
		if bal.IsLT(msg.Locked) {
			return nil, sdkerrors.ErrInsufficientFunds
		}
		if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender,
			types.ModuleName, sdk.NewCoins(msg.Locked)); err != nil {
			panic(err)
		}
	}

	rv = types.RValidator{Denom: msg.Denom, Address: msg.Address, Status: types.Onboard, Locked: msg.Locked}
	k.Keeper.AddRValidator(ctx, rv)
	k.Keeper.AddRValidatorToSet(ctx, rv)

	return &types.MsgOnboardResponse{}, nil
}
