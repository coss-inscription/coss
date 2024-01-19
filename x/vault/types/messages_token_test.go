package types

import (
	"testing"

	"coss/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateToken_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateToken
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateToken{
				Owner: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "all good",
			msg: MsgCreateToken{
				Owner: sample.AccAddress(),
				Denom:       "denom",
				Description: "description",
				Symbol:      "SYMBOL",
				Url:         "url",
				Decimals:    6,
				MaxSupply:   1000000000000000000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateToken_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateToken
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateToken{
				Owner: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "all good",
			msg: MsgUpdateToken{
				Owner: sample.AccAddress(),
				Denom:       "denom",
				Description: "description",
				Symbol:      "SYMBOL",
				Url:         "url",
				Decimals:    18,
				MaxSupply:   1000000000000000000,
			},
		}, {
			name: "invalid denom",
			msg: MsgUpdateToken{
				Owner:       sample.AccAddress(),
				Denom:       "",
				Description: "description",
				Symbol:      "SYMBOL",
				Url:         "url",
				Decimals:    6,
				MaxSupply:   1000000000000000000,
		},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid denom",
			msg: MsgUpdateToken{
				Owner:       sample.AccAddress(),
				Denom:       "ab",
				Description: "description",
				Symbol:      "SYMBOL",
				Url:         "url",
				Decimals:    6,
				MaxSupply:   1000000000000000000,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid denom",
			msg: MsgUpdateToken{
				Owner:       sample.AccAddress(),
				Denom:       "abcdefghijklmnopqrs",
				Description: "description",
				Symbol:      "SYMBOL",
				Url:         "url",
				Decimals:    6,
				MaxSupply:   1000000000000000000,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid denom",
			msg: MsgUpdateToken{
				Owner:       sample.AccAddress(),
				Denom:       "AB",
				Description: "description",
				Symbol:      "SYMBOL",
				Url:         "url",
				Decimals:    6,
				MaxSupply:   1000000000000000000,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid description",
			msg: MsgUpdateToken{
				Owner:       sample.AccAddress(),
				Denom:       "denom",
				Description: "",
				Symbol:      "SYMBOL",
				Url:         "url",
				Decimals:    6,
				MaxSupply:   1000000000000000000,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid symbol",
			msg: MsgUpdateToken{
				Owner:       sample.AccAddress(),
				Denom:       "denom",
				Description: "description",
				Symbol:      "",
				Url:         "url",
				Decimals:    6,
				MaxSupply:   1000000000000000000,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid symbol",
			msg: MsgUpdateToken{
				Owner:       sample.AccAddress(),
				Denom:       "denom",
				Description: "description",
				Symbol:      "symbol",
				Url:         "url",
				Decimals:    6,
				MaxSupply:   1000000000000000000,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid symbol",
			msg: MsgUpdateToken{
				Owner:       sample.AccAddress(),
				Denom:       "denom",
				Description: "description",
				Symbol:      "SYMBOLSYMBOLSYMBOLSYMBOLSYMBOL",
				Url:         "url",
				Decimals:    6,
				MaxSupply:   1000000000000000000,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid url",
			msg: MsgUpdateToken{
				Owner:       sample.AccAddress(),
				Denom:       "denom",
				Description: "description",
				Symbol:      "symbol",
				Url:         "",
				Decimals:    6,
				MaxSupply:   1000000000000000000,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid decimals",
			msg: MsgUpdateToken{
				Owner: sample.AccAddress(),
				Denom:       "denom",
				Description: "description",
				Symbol:      "symbol",
				Url:         "",
				Decimals:    -1,
				MaxSupply:   1000000000000000000,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid decimals",
			msg: MsgUpdateToken{
				Owner:       sample.AccAddress(),
				Denom:       "denom",
				Description: "description",
				Symbol:      "symbol",
				Url:         "url",
				Decimals:    19,
				MaxSupply:   1000000000000000000,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid maxSupply",
			msg: MsgUpdateToken{
				Owner:       sample.AccAddress(),
				Denom:       "denom",
				Description: "description",
				Symbol:      "symbol",
				Url:         "",
				Decimals:    6,
				MaxSupply:   0,
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
