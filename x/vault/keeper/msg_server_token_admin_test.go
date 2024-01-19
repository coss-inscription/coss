package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "coss/testutil/keeper"
	"coss/x/vault/keeper"
	"coss/x/vault/types"
)

func TestTokenAdminMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.VaultKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	expected := &types.MsgCreateTokenAdmin{Creator: creator}
	_, err := srv.CreateTokenAdmin(wctx, expected)
	require.NoError(t, err)
	rst, found := k.GetTokenAdmin(ctx)
	require.True(t, found)
	require.Equal(t, expected.Creator, rst.Creator)
}

func TestTokenAdminMsgServerUpdate(t *testing.T) {
	creator := "A"
	address := "C"

	tests := []struct {
		desc    string
		request *types.MsgUpdateTokenAdmin
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateTokenAdmin{Creator: creator, Address: address},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTokenAdmin{Creator: "B", Address: address},
			err:     sdkerrors.ErrUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.VaultKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateTokenAdmin{Creator: creator, Address: address}
			_, err := srv.CreateTokenAdmin(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateTokenAdmin(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetTokenAdmin(ctx)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
				require.Equal(t, expected.Address, rst.Address)
			}
		})
	}
}
