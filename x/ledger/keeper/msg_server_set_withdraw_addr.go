package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafihub/stafihub/x/ledger/types"
)

func (k msgServer) SetWithdrawAddr(goCtx context.Context, msg *types.MsgSetWithdrawAddr) (*types.MsgSetWithdrawAddrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	icaPoolDetail, found := k.GetIcaPoolByDelegationAddr(ctx, msg.DelegationAddr)
	if !found {
		return nil, types.ErrIcaPoolNotFound
	}

	if icaPoolDetail.Status != types.IcaPoolStatusWithdrawCreate {
		return nil, types.ErrIcaPoolStatusUnmatch
	}

	err := k.Keeper.SetWithdrawAddressOnHost(
		ctx,
		icaPoolDetail.DelegationAccount.Owner,
		icaPoolDetail.DelegationAccount.CtrlConnectionId,
		icaPoolDetail.DelegationAccount.Address,
		icaPoolDetail.WithdrawAccount.Address)
	if err != nil {
		return nil, err
	}

	return &types.MsgSetWithdrawAddrResponse{}, nil
}
