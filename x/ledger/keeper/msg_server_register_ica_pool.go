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

	if err := k.Keeper.ICAControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, delegationOwner); err != nil {
		return nil, err
	}
	if err := k.Keeper.ICAControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, withdrawOwner); err != nil {
		return nil, err
	}

	k.SetIcaPoolDetail(ctx, types.IcaPoolDetail{
		Denom:    msg.Denom,
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
	k.SetIcaPoolSequence(ctx, msg.Denom, willUseSequence)

	return &types.MsgRegisterIcaPoolResponse{}, nil
}
