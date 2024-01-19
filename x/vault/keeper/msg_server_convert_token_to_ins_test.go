package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "coss/testutil/keeper"
	"coss/x/vault/keeper"
	"coss/x/vault/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestConvertTokenToInsMsgServerNotAudit(t *testing.T) {
	k, ctx := keepertest.VaultKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)

	onwer := "A"
	denom := "denom"
	recipient := "C"
	CreateToken(t, ctx, srv, onwer, denom, 1)
	_, mintErr := srv.ConvertInsToToken(ctx, &types.MsgConvertInsToToken{
		Owner:     onwer,
		Recipient: recipient,
		Amount:    sdk.NewCoin(denom, sdkmath.NewInt(1))})
	require.ErrorIs(t, mintErr, sdkerrors.ErrUnauthorized)
}

func TestConvertTokenToInsMsgServerAudit(t *testing.T) {
	k, ctx := keepertest.VaultKeeper(t)

	srv := keeper.NewMsgServerImpl(*k)

	onwer := "A"
	denom := "denom"
	sender := "cosmos130gkej3y9j2c5exc6yptrfhqaxspq5frt3z2et"
	CreateToken(t, ctx, srv, onwer, denom, 10)
	_, adminErr := srv.CreateTokenAdmin(ctx, &types.MsgCreateTokenAdmin{Creator: onwer, Address: onwer})
	require.NoError(t, adminErr)
	_, auditErr := srv.AuditToken(ctx, &types.MsgAuditToken{TokenAdmin: onwer, Denom: denom, Pass: true})
	require.NoError(t, auditErr)

	tests := []struct {
		desc    string
		request *types.MsgConvertTokenToIns
		err     error
	}{
		{
			desc:    "Token not found",
			request: &types.MsgConvertTokenToIns{Sender: sender, Amount: sdk.NewCoin("denom2", sdkmath.NewInt(1))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc:    "Insufficient token supply",
			request: &types.MsgConvertTokenToIns{Sender: sender, Amount: sdk.NewCoin(denom, sdkmath.NewInt(11))},
			err:     sdkerrors.ErrInvalidRequest,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := srv.ConvertTokenToIns(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
