package types

import (
	"regexp"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateToken = "create_token"
	TypeMsgUpdateToken = "update_token"
)

var _ sdk.Msg = &MsgCreateToken{}

func NewMsgCreateToken(
	owner string,
	denom string,
	description string,
	symbol string,
	decimals int32,
	url string,
	maxSupply uint64,

) *MsgCreateToken {
	return &MsgCreateToken{
		Owner:       owner,
		Denom:       denom,
		Description: description,
		Symbol:      symbol,
		Decimals:    decimals,
		Url:         url,
		MaxSupply:   maxSupply,
	}
}

func ValidateBasic(
	owner string,
	denom string,
	description string,
	symbol string,
	decimals int32,
	url string,
	maxSupply uint64) error {

	_, err := sdk.AccAddressFromBech32(owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if denom == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "denom cannot be empty")
	} else if len(denom) > 17 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "denom cannot be longer than 17 characters")
	} else if len(denom) < 4 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "denom cannot be shorter than 4 characters")
	}
	match, _ := regexp.MatchString(`^[a-z]+$`, denom)
	if !match {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "denom must be lowercase")
	}

	if description == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "description cannot be empty")
	} else if len(description) > 1024 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "description cannot be longer than 1024 characters")
	}
	if symbol == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "symbol cannot be empty")
	} else if len(symbol) > 16 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "symbol cannot be longer than 16 characters")
	} else if len(symbol) < 3 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "symbol cannot be shorter than 3 characters")
	}
	matchSymbol, _ := regexp.MatchString(`^[A-Z]+$`, symbol)
	if !matchSymbol {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "symbol must be uppercase")
	}

	if decimals < 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "decimals cannot be negative")
	} else if decimals > 18 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "decimals cannot be greater than 18")
	}
	if url == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "url cannot be empty")
	}
	if maxSupply == 0 {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "maxSupply cannot be 0")
	}
	return nil
}

func (msg *MsgCreateToken) Route() string {
	return RouterKey
}

func (msg *MsgCreateToken) Type() string {
	return TypeMsgCreateToken
}

func (msg *MsgCreateToken) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgCreateToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateToken) ValidateBasic() error {
	return ValidateBasic(msg.Owner, msg.Denom, msg.Description, msg.Symbol, msg.Decimals, msg.Url, msg.MaxSupply)
}

var _ sdk.Msg = &MsgUpdateToken{}

func NewMsgUpdateToken(
	owner string,
	denom string,
	description string,
	symbol string,
	decimals int32,
	url string,
	maxSupply uint64,

) *MsgUpdateToken {
	return &MsgUpdateToken{
		Owner:       owner,
		Denom:       denom,
		Description: description,
		Symbol:      symbol,
		Decimals:    decimals,
		Url:         url,
		MaxSupply:   maxSupply,
	}
}

func (msg *MsgUpdateToken) Route() string {
	return RouterKey
}

func (msg *MsgUpdateToken) Type() string {
	return TypeMsgUpdateToken
}

func (msg *MsgUpdateToken) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgUpdateToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateToken) ValidateBasic() error {
	return ValidateBasic(msg.Owner, msg.Denom, msg.Description, msg.Symbol, msg.Decimals, msg.Url, msg.MaxSupply)
}
