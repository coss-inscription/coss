package types

import (
	"testing"

	"coss/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgConvertTokenToIns_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgConvertTokenToIns
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgConvertTokenToIns{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgConvertTokenToIns{
				Sender: sample.AccAddress(),
				Amount: sample.Coin(1),
			},
		}, {
			name: "invalid coins",
			msg: MsgConvertTokenToIns{
				Sender: sample.AccAddress(),
				Amount: sample.Coin(0),
			},
			err: sdkerrors.ErrInvalidCoins,
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
