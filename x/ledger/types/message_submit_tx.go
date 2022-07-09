package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitTx = "submit_tx"

var _ sdk.Msg = &MsgSubmitTx{}

func NewMsgSubmitTx(creator string, owner string, connectionId string, msgs []sdk.Msg) (*MsgSubmitTx, error) {
	any, err := PackTxMsgAny(msgs)
	if err != nil {
		return nil, err
	}
	return &MsgSubmitTx{
		Creator:      creator,
		Owner:        owner,
		ConnectionId: connectionId,
		Msgs:         any,
	}, nil
}

func (msg *MsgSubmitTx) Route() string {
	return RouterKey
}

func (msg *MsgSubmitTx) Type() string {
	return TypeMsgSubmitTx
}

func (msg *MsgSubmitTx) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// PackTxMsgAny marshals the sdk.Msg payload to a protobuf Any type
func PackTxMsgAny(msgs []sdk.Msg) ([]*codectypes.Any, error) {

	msgAnys := make([]*codectypes.Any, len(msgs))
	var err error
	for i, msg := range msgs {
		msgAnys[i], err = codectypes.NewAnyWithValue(msg)
		if err != nil {
			return nil, err
		}
	}
	return msgAnys, nil
}

// GetTxMsg fetches the cached any message
func (msg *MsgSubmitTx) GetTxMsg(c codec.BinaryCodec) ([]sdk.Msg, error) {
	msgs := make([]sdk.Msg, len(msg.Msgs))

	for i, any := range msg.Msgs {
		var msg sdk.Msg
		err := c.UnpackAny(any, &msg)
		if err != nil {
			return nil, err
		}
		msgs[i] = msg
	}

	return msgs, nil
}
