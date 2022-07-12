package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) RegisterIcaPool(goCtx context.Context, msg *types.MsgRegisterIcaPool) (*types.MsgRegisterIcaPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	willUseIndex := k.GetIcaPoolNextIndex(ctx, msg.Denom)
	delegationOwner, withdrawOwner := types.GetOwners(msg.Denom, willUseIndex)

	if err := k.Keeper.ICAControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, delegationOwner); err != nil {
		return nil, err
	}
	if err := k.Keeper.ICAControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, withdrawOwner); err != nil {
		return nil, err
	}

	k.SetIcaPoolDetail(ctx, &types.IcaPoolDetail{
		Denom:  msg.Denom,
		Status: types.IcaPoolStatusInit,
		Index:  willUseIndex,
		DelegationAccount: &types.IcaAccount{
			Owner:            delegationOwner,
			CtrlConnectionId: msg.ConnectionId,
		},
		WithdrawAccount: &types.IcaAccount{
			Owner:            withdrawOwner,
			CtrlConnectionId: msg.ConnectionId,
		},
	})

	k.SetIcaPoolIndex(ctx, msg.Denom, willUseIndex)

	return &types.MsgRegisterIcaPoolResponse{}, nil
}
