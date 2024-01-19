package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConvertInsToToken = "convert_ins_to_token"

var _ sdk.Msg = &MsgConvertInsToToken{}

func NewMsgConvertInsToToken(owner string, amount sdk.Coin, recipient string) *MsgConvertInsToToken {
	return &MsgConvertInsToToken{
		Owner:     owner,
		Amount:    amount,
		Recipient: recipient,
	}
}

func (msg *MsgConvertInsToToken) Route() string {
	return RouterKey
}

func (msg *MsgConvertInsToToken) Type() string {
	return TypeMsgConvertInsToToken
}

func (msg *MsgConvertInsToToken) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgConvertInsToToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConvertInsToToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	// check recipient address
	_, errRecipient := sdk.AccAddressFromBech32(msg.Recipient)
	if errRecipient != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receipient address (%s)", err)
	}
	if msg.Amount.IsZero() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, "amount should be positive")
	}
	return nil
}
