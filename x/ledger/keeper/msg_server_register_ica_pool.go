package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	"github.com/stafihub/stafihub/x/ledger/types"
	sudotypes "github.com/stafihub/stafihub/x/sudo/types"
)

func (k msgServer) RegisterIcaPool(goCtx context.Context, msg *types.MsgRegisterIcaPool) (*types.MsgRegisterIcaPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	// ensure checkAddress work well
	_, ok := k.bankKeeper.GetDenomMetaData(ctx, msg.Denom)
	if !ok {
		return nil, banktypes.ErrDenomMetadataNotFound
	}

	willUseIndex := k.GetIcaPoolNextIndex(ctx, msg.Denom)
	delegationOwner, withdrawalOwner := types.GetOwners(msg.Denom, willUseIndex)

	connectionEnd, found := k.IBCKeeper.ConnectionKeeper.GetConnection(ctx, msg.ConnectionId)
	if !found {
		return nil, types.ErrConnectionIdNotFound
	}

	appVersion := string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
		Version:                icatypes.Version,
		ControllerConnectionId: msg.ConnectionId,
		HostConnectionId:       connectionEnd.Counterparty.ConnectionId,
		Encoding:               icatypes.EncodingProtobuf,
		TxType:                 icatypes.TxTypeSDKMultiMsg,
	}))

	if err := k.Keeper.ICAControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, delegationOwner, appVersion); err != nil {
		return nil, err
	}
	if err := k.Keeper.ICAControllerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, withdrawalOwner, appVersion); err != nil {
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
		WithdrawalAccount: &types.IcaAccount{
			Owner:            withdrawalOwner,
			CtrlConnectionId: msg.ConnectionId,
		},
	})

	k.SetIcaPoolIndex(ctx, msg.Denom, willUseIndex)

	return &types.MsgRegisterIcaPoolResponse{}, nil
}
