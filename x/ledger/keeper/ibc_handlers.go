package keeper

import (
	"encoding/hex"
	"fmt"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	icatypes "github.com/cosmos/ibc-go/v5/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	ibchost "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	proto "github.com/gogo/protobuf/proto"
	"github.com/stafihub/stafihub/x/ledger/types"
)

// Implements core logic for OnAcknowledgementPacket
func (k Keeper) OnAcknowledgement(ctx sdk.Context, modulePacket channeltypes.Packet, acknowledgement []byte) error {
	k.Logger(ctx).Debug("OnAcknowledgement start--------------------------", "acknowledgement", string(acknowledgement))

	// parse acknowledgement
	var ack channeltypes.Acknowledgement
	err := channeltypes.SubModuleCdc.UnmarshalJSON(acknowledgement, &ack)
	if err != nil {
		k.Logger(ctx).Error("Unable to unmarshal acknowledgement error", "error", err)
		return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal acknowledgement data: %s", err.Error())
	}

	var ackSuccess bool
	switch response := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		if len(response.Result) == 0 {
			return errorsmod.Wrapf(channeltypes.ErrInvalidAcknowledgement, "acknowledgement result cannot be empty")
		}
		ackSuccess = true
	case *channeltypes.Acknowledgement_Error:
		ackSuccess = false
	default:
		return errorsmod.Wrapf(channeltypes.ErrInvalidAcknowledgement, "unsupported acknowledgement response field type %T", response)
	}

	if !ackSuccess {
		// acknowledgement error
		k.Logger(ctx).Info("acknowledgement error", "ack_err", ack.GetError())
		// update interchain tx status
		propId, found := k.GetInterchainTxPropIdBySeq(ctx, modulePacket.SourcePort, modulePacket.SourceChannel, modulePacket.Sequence)
		if found {
			k.SetInterchainTxProposalStatus(ctx, propId, types.InterchainTxStatusFailed)
		}
		return nil
	}
	// acknowledgement success
	k.Logger(ctx).Debug("acknowledgement success --------------------------", "ack", hex.EncodeToString(ack.GetResult()))

	// parse txMsgData, ack includes tx exec result info
	msgTypes, _, err := ParseTxMsgData(ack.GetResult())
	if err != nil {
		k.Logger(ctx).Error("Unable to unmarshal ack.Result", "error", err, "ack.Result", ack.GetResult())
		return err
	}

	isCallBackOfMsgSetWithdrawAddress := false
	if len(msgTypes) != 0 {
		if strings.EqualFold(msgTypes[0], "/cosmos.distribution.v1beta1.MsgSetWithdrawAddress") ||
			strings.EqualFold(msgTypes[0], "/cosmos.distribution.v1beta1.MsgSetWithdrawAddressResponse") {
			isCallBackOfMsgSetWithdrawAddress = true
		}
	}

	if isCallBackOfMsgSetWithdrawAddress {
		// parse packet data
		var packetData icatypes.InterchainAccountPacketData
		err = icatypes.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &packetData)
		if err != nil {
			k.Logger(ctx).Error("unable to unmarshal modulePacket data", "error", err, "data", packetData)
			return errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal modulePacket data: %s", err.Error())
		}

		msgs, err := icatypes.DeserializeCosmosTx(k.cdc, packetData.Data)
		if err != nil {
			k.Logger(ctx).Info("Error decoding messages", "err", err)
			return err
		}
		if len(msgs) == 0 {
			return fmt.Errorf("msgs of packetData is empty")
		}

		msgSetWithdrawAddr, ok := msgs[0].(*distributiontypes.MsgSetWithdrawAddress)
		if !ok {
			errStr := "unable to cast source message to MsgSetWithdrawAddress"
			k.Logger(ctx).Error(errStr)
			return fmt.Errorf(errStr)
		}

		// update ica pool status
		icaPool, found := k.GetIcaPoolByDelegationAddr(ctx, msgSetWithdrawAddr.DelegatorAddress)
		if !found {
			return types.ErrIcaPoolNotFound
		}
		icaPool.Status = types.IcaPoolStatusSetWithdrawal

		k.SetIcaPoolDetail(ctx, icaPool)
	} else {
		// update interchain tx status
		propId, found := k.GetInterchainTxPropIdBySeq(ctx, modulePacket.SourcePort, modulePacket.SourceChannel, modulePacket.Sequence)
		if !found {
			return types.ErrInterchainTxPropIdNotFound
		}
		k.SetInterchainTxProposalStatus(ctx, propId, types.InterchainTxStatusSuccess)
	}

	k.Logger(ctx).Debug("OnAcknowledgement end --------------------------")
	return nil
}

func (k Keeper) SetWithdrawAddressOnHost(ctx sdk.Context, delegationAddrOwner, ctrlConnectionId, delegationAddr, withdrawAddr string) error {
	var msgs []sdk.Msg

	k.Logger(ctx).Info(fmt.Sprintf("Setting withdrawal address on host.delegationAddrOwner: %s DelegatorAddress: %s WithdrawAddress: %s ctrlConnectionID: %s",
		delegationAddrOwner, delegationAddr, withdrawAddr, ctrlConnectionId))
	// construct the msg
	msgs = append(msgs, &distributiontypes.MsgSetWithdrawAddress{DelegatorAddress: delegationAddr, WithdrawAddress: withdrawAddr})
	// Send the transaction through SubmitTx
	_, err := k.SubmitTxs(ctx, ctrlConnectionId, delegationAddrOwner, msgs, "MsgSetWithdrawAddress")
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Failed to SubmitTxs for %s, %s, %s", ctrlConnectionId, delegationAddrOwner, msgs)
	}
	return nil
}

// SubmitTxs submits an ICA transaction containing multiple messages
func (k Keeper) SubmitTxs(ctx sdk.Context, ctrlConnectionId, owner string, msgs []sdk.Msg, memo string) (uint64, error) {
	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		return 0, err
	}

	channelID, found := k.ICAControllerKeeper.GetActiveChannelID(ctx, ctrlConnectionId, portID)
	if !found {
		return 0, errorsmod.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
	}

	chanCap, found := k.scopedKeeper.GetCapability(ctx, ibchost.ChannelCapabilityPath(portID, channelID))
	if !found {
		return 0, errorsmod.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	data, err := icatypes.SerializeCosmosTx(k.cdc, msgs)
	if err != nil {
		return 0, err
	}

	packetData := icatypes.InterchainAccountPacketData{
		Type: icatypes.EXECUTE_TX,
		Data: data,
		Memo: memo,
	}

	timeoutTimestamp := uint64(ctx.BlockTime().Add(types.ICATxTimeout).UnixNano())
	sequence, err := k.ICAControllerKeeper.SendTx(ctx, chanCap, ctrlConnectionId, portID, packetData, timeoutTimestamp)
	if err != nil {
		return 0, err
	}

	return sequence, nil
}

// Parses ICA tx responses and returns a list of each serialized response
// The format of the raw ack differs depending on which version of ibc-go is used
// For v5 and prior, the message responses are stored under the `Data` attribute of TxMsgData
// For v6 and later, the message responses are stored under the `MsgResponse` attribute of TxMsgdata
func ParseTxMsgData(acknowledgementResult []byte) ([]string, [][]byte, error) {
	txMsgData := &sdk.TxMsgData{}
	if err := proto.Unmarshal(acknowledgementResult, txMsgData); err != nil {
		return nil, nil, errorsmod.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-27 tx message data: %s", err.Error())
	}

	// Unpack all the message responses based on the sdk version (determined from the length of txMsgData.Data)
	switch len(txMsgData.Data) {
	case 0:
		// for SDK 0.46 and above
		msgDatas := make([][]byte, len(txMsgData.MsgResponses))
		msgTypes := make([]string, len(txMsgData.MsgResponses))
		for i, msgResponse := range txMsgData.MsgResponses {
			msgDatas[i] = msgResponse.GetValue()
			msgTypes[i] = msgResponse.GetTypeUrl()
		}
		return msgTypes, msgDatas, nil
	default:
		// for SDK 0.45 and below
		var msgDatas = make([][]byte, len(txMsgData.Data))
		var msgTypes = make([]string, len(txMsgData.Data))
		for i, msgData := range txMsgData.Data {
			msgDatas[i] = msgData.Data
			msgTypes[i] = msgData.MsgType
		}
		return msgTypes, msgDatas, nil
	}
}
