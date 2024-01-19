package types

import (
	"testing"

	"coss/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateTokenAdmin_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTokenAdmin
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTokenAdmin{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTokenAdmin{
				Creator: sample.AccAddress(),
				Address: sample.AccAddress(),
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

func TestMsgUpdateTokenAdmin_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateTokenAdmin
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateTokenAdmin{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateTokenAdmin{
				Creator: sample.AccAddress(),
				Address: sample.AccAddress(),
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
