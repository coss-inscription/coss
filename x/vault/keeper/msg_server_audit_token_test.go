package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "coss/testutil/keeper"
	"coss/x/vault/keeper"
	"coss/x/vault/types"
)

func TestAuditTokenMsgServer(t *testing.T) {
	k, ctx := keepertest.VaultKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	// wctx := sdk.WrapSDKContext(ctx)

	admin := "admin"
	denom := "denom"

	_, adminErr := srv.CreateTokenAdmin(ctx, &types.MsgCreateTokenAdmin{Creator: admin, Address: admin})
	require.NoError(t, adminErr)
	token := &types.MsgCreateToken{Owner: "A", Denom: denom}
	_, tokenErr := srv.CreateToken(ctx, token)
	require.NoError(t, tokenErr)

	tests := []struct {
		desc    string
		request *types.MsgAuditToken
		err     error
		pass    bool
	}{
		{
			desc:    "Completed",
			request: &types.MsgAuditToken{TokenAdmin: admin, Denom: denom, Pass: true},
			pass:    true,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgAuditToken{TokenAdmin: "B", Denom: denom, Pass: true},
			err:     sdkerrors.ErrUnauthorized,
			pass:    true,
		},
		{
			desc:    "ReAudit Unauthorized",
			request: &types.MsgAuditToken{TokenAdmin: admin, Denom: denom, Pass: false},
			err:     sdkerrors.ErrUnauthorized,
			pass:    true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := srv.AuditToken(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetToken(ctx, denom)
				require.True(t, found)
				require.Equal(t, tc.pass, rst.Audited)
			}
		})
	}
}
