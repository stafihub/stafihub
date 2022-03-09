package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stafihub/stafihub/utils"
)

const TypeMsgMigrateInit = "migrate_init"

var _ sdk.Msg = &MsgMigrateInit{}

func NewMsgMigrateInit(creator string, denom, pool string, totalSupply, active, bond, unbond sdk.Int, exchangeRate utils.Dec) *MsgMigrateInit {
	return &MsgMigrateInit{
		Creator:      creator,
		Denom:        denom,
		Pool:         pool,
		TotalSupply:  totalSupply,
		Active:       active,
		Bond:         bond,
		Unbond:       unbond,
		ExchangeRate: exchangeRate,
	}
}

func (msg *MsgMigrateInit) Route() string {
	return RouterKey
}

func (msg *MsgMigrateInit) Type() string {
	return TypeMsgMigrateInit
}

func (msg *MsgMigrateInit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMigrateInit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMigrateInit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
