package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConvertTokenToIns = "convert_token_to_ins"

var _ sdk.Msg = &MsgConvertTokenToIns{}

func NewMsgConvertTokenToIns(sender string, amount sdk.Coin) *MsgConvertTokenToIns {
	return &MsgConvertTokenToIns{
		Sender: sender,
		Amount: amount,
	}
}

func (msg *MsgConvertTokenToIns) Route() string {
	return RouterKey
}

func (msg *MsgConvertTokenToIns) Type() string {
	return TypeMsgConvertTokenToIns
}

func (msg *MsgConvertTokenToIns) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgConvertTokenToIns) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConvertTokenToIns) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if msg.Amount.IsZero() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "invalid amount (%s)", err)
	}
	return nil
}
