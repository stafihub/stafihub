package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) RegisterIcaPool(goCtx context.Context, msg *types.MsgRegisterIcaPool) (*types.MsgRegisterIcaPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	willUseSequence := k.GetIcapPoolNextSequence(ctx, msg.Denom)
	delegationOwner := fmt.Sprintf("%s-%d-delegation", msg.Denom, willUseSequence)
	withdrawOwner := fmt.Sprintf("%s-%d-withdraw", msg.Denom, willUseSequence)
	ctx.Logger().Info("RegisterIcaPool", "start1", msg)
	if err := k.Keeper.ICAControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, delegationOwner); err != nil {
		return nil, err
	}
	ctx.Logger().Info("RegisterIcaPool", "start2", msg)
	if err := k.Keeper.ICAControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, withdrawOwner); err != nil {
		return nil, err
	}

	ctx.Logger().Info("RegisterIcaPool", "start3", msg)
	k.SetIcaPoolDetail(ctx, &types.IcaPoolDetail{
		Denom:    msg.Denom,
		Status:   types.IcaPoolStatusInit,
		Sequence: fmt.Sprintf("%d", willUseSequence),
		DelegationAccount: &types.IcaAccount{
			Owner:            delegationOwner,
			CtrlConnectionId: msg.ConnectionId,
		},
		WithdrawAccount: &types.IcaAccount{
			Owner:            withdrawOwner,
			CtrlConnectionId: msg.ConnectionId,
		},
	})
	ctx.Logger().Info("RegisterIcaPool", "start4", msg)
	k.SetIcaPoolSequence(ctx, msg.Denom, willUseSequence)
	ctx.Logger().Info("RegisterIcaPool", "start5", willUseSequence)
	return &types.MsgRegisterIcaPoolResponse{}, nil
}
