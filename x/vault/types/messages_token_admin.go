package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateTokenAdmin = "create_token_admin"
	TypeMsgUpdateTokenAdmin = "update_token_admin"
)

var _ sdk.Msg = &MsgCreateTokenAdmin{}

func NewMsgCreateTokenAdmin(creator string, address string) *MsgCreateTokenAdmin {
	return &MsgCreateTokenAdmin{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgCreateTokenAdmin) Route() string {
	return RouterKey
}

func (msg *MsgCreateTokenAdmin) Type() string {
	return TypeMsgCreateTokenAdmin
}

func (msg *MsgCreateTokenAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTokenAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTokenAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err1 := sdk.AccAddressFromBech32(msg.Address)
	if err1 != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err1)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTokenAdmin{}

func NewMsgUpdateTokenAdmin(creator string, address string) *MsgUpdateTokenAdmin {
	return &MsgUpdateTokenAdmin{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgUpdateTokenAdmin) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTokenAdmin) Type() string {
	return TypeMsgUpdateTokenAdmin
}

func (msg *MsgUpdateTokenAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTokenAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTokenAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err1 := sdk.AccAddressFromBech32(msg.Address)
	if err1 != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid address (%s)", err1)
	}
	return nil
}
