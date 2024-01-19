package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAuditToken = "audit_token"

var _ sdk.Msg = &MsgAuditToken{}

func NewMsgAuditToken(tokenAdmin string, denom string, pass bool) *MsgAuditToken {
	return &MsgAuditToken{
		TokenAdmin: tokenAdmin,
		Denom:      denom,
		Pass:       pass,
	}
}

func (msg *MsgAuditToken) Route() string {
	return RouterKey
}

func (msg *MsgAuditToken) Type() string {
	return TypeMsgAuditToken
}

func (msg *MsgAuditToken) GetSigners() []sdk.AccAddress {
	tokenAdmin, err := sdk.AccAddressFromBech32(msg.TokenAdmin)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{tokenAdmin}
}

func (msg *MsgAuditToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAuditToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.TokenAdmin)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid tokenAdmin address (%s)", err)
	}
	if msg.Denom == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "denom cannot be empty")
	}
	return nil
}
