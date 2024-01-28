package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) OpenIcaChannel(goCtx context.Context, msg *types.MsgOpenIcaChannel) (*types.MsgOpenIcaChannelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	icaPool, found := k.GetIcaPoolByDelegationAddr(ctx, msg.PoolAddress)
	if !found {
		return nil, types.ErrIcaPoolNotFound
	}

	switch msg.AccountType {
	case types.AccountTypeDelegation:
		appVersion := string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
			Version:                icatypes.Version,
			ControllerConnectionId: icaPool.DelegationAccount.CtrlConnectionId,
			HostConnectionId:       icaPool.DelegationAccount.HostConnectionId,
			Encoding:               icatypes.EncodingProtobuf,
			TxType:                 icatypes.TxTypeSDKMultiMsg,
		}))
		if err := k.Keeper.ICAControllerKeeper.RegisterInterchainAccount(ctx, icaPool.DelegationAccount.CtrlConnectionId,
			icaPool.DelegationAccount.Owner, appVersion); err != nil {
			return nil, err
		}
	case types.AccountTypeWithdraw:
		appVersion := string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
			Version:                icatypes.Version,
			ControllerConnectionId: icaPool.WithdrawalAccount.CtrlConnectionId,
			HostConnectionId:       icaPool.WithdrawalAccount.HostConnectionId,
			Encoding:               icatypes.EncodingProtobuf,
			TxType:                 icatypes.TxTypeSDKMultiMsg,
		}))

		if err := k.Keeper.ICAControllerKeeper.RegisterInterchainAccount(ctx, icaPool.WithdrawalAccount.CtrlConnectionId,
			icaPool.WithdrawalAccount.Owner, appVersion); err != nil {
			return nil, err
		}
	default:
		return nil, types.ErrUnknownAccountType
	}

	return &types.MsgOpenIcaChannelResponse{}, nil
}
