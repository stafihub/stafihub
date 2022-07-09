package keeper

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/golang/protobuf/proto"

	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
)

// Implements core logic for OnAcknowledgementPacket
func (k Keeper) OnAcknowledgement(ctx sdk.Context, modulePacket channeltypes.Packet, acknowledgement []byte) error {
	k.Logger(ctx).Info("OnAcknowledgement start --------------------------")
	ack := channeltypes.Acknowledgement_Result{}

	err := json.Unmarshal(acknowledgement, &ack)
	if err != nil {
		ackErr := channeltypes.Acknowledgement_Error{}
		err := json.Unmarshal(acknowledgement, &ackErr)
		if err != nil {
			k.Logger(ctx).Error("Unable to unmarshal acknowledgement error", "error", err, "data", acknowledgement)
			return err
		}
		k.Logger(ctx).Error("Unable to unmarshal acknowledgement result", "error", err, "remote_err", ackErr, "data", acknowledgement)
		return err
	}
	k.Logger(ctx).Info("OnAcknowledgement start --------------------------", "ack", ack)
	txMsgData := &sdk.TxMsgData{}
	err = proto.Unmarshal(ack.Result, txMsgData)
	if err != nil {
		k.Logger(ctx).Error("Unable to unmarshal acknowledgement", "error", err, "ack", ack.Result)
		return err
	}

	k.Logger(ctx).Info("OnAcknowledgement start --------------------------", "txMsgData", txMsgData.String())

	var packetData icatypes.InterchainAccountPacketData
	err = icatypes.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &packetData)
	if err != nil {
		k.Logger(ctx).Error("unable to unmarshal acknowledgement packet data", "error", err, "data", packetData)
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal packet data: %s", err.Error())
	}

	for _, msgData := range txMsgData.Data {
		switch msgData.MsgType {
		// staking to validators
		case "/cosmos.staking.v1beta1.MsgDelegate":
			k.Logger(ctx).Info("onAcknowledgement msg delegate--------------------------")
		// unstake
		case "/cosmos.staking.v1beta1.MsgUndelegate":
		// withdrawing rewards ()
		case "/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward":
		case "/cosmos.bank.v1beta1.MsgSend":
		case "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress":
		default:
			k.Logger(ctx).Error("Unhandled acknowledgement packet", "type", msgData.MsgType)
		}
	}
	k.Logger(ctx).Info("onAcknowledgement msg delegate end --------------------------")
	return nil
}
