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

func CreateToken(t *testing.T, ctx sdk.Context, srv types.MsgServer, owner string, denom string, supply int32) {
	token := &types.MsgCreateToken{Owner: owner, Denom: denom, MaxSupply: uint64(supply)}
	_, tokenErr := srv.CreateToken(ctx, token)
	require.NoError(t, tokenErr)
}

func TestConvertInsToTokenMsgServerNotAudit(t *testing.T) {
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

func TestConvertInsToTokenMsgServerAudit(t *testing.T) {
	k, ctx := keepertest.VaultKeeper(t)

	srv := keeper.NewMsgServerImpl(*k)

	onwer := "A"
	denom := "denom"
	recipient := "cosmos130gkej3y9j2c5exc6yptrfhqaxspq5frt3z2et"
	CreateToken(t, ctx, srv, onwer, denom, 10)
	_, adminErr := srv.CreateTokenAdmin(ctx, &types.MsgCreateTokenAdmin{Creator: onwer, Address: onwer})
	require.NoError(t, adminErr)
	_, auditErr := srv.AuditToken(ctx, &types.MsgAuditToken{TokenAdmin: onwer, Denom: denom, Pass: true})
	require.NoError(t, auditErr)

	tests := []struct {
		desc    string
		request *types.MsgConvertInsToToken
		err     error
		supply  int32
	}{
		{
			desc:    "Token not found",
			request: &types.MsgConvertInsToToken{Owner: onwer, Recipient: recipient, Amount: sdk.NewCoin("denom2", sdkmath.NewInt(1))},
			err:     sdkerrors.ErrKeyNotFound,
			supply:  0,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgConvertInsToToken{Owner: "B", Recipient: recipient, Amount: sdk.NewCoin(denom, sdkmath.NewInt(1))},
			err:     sdkerrors.ErrUnauthorized,
			supply:  10,
		},
		{
			desc:    "Mint more than Max Supply",
			request: &types.MsgConvertInsToToken{Owner: onwer, Recipient: recipient, Amount: sdk.NewCoin(denom, sdkmath.NewInt(11))},
			err:     sdkerrors.ErrInvalidRequest,
			supply:  10,
		},
		// {
		// 	desc:    "Completed",
		// 	request: &types.MsgConvertInsToToken{Sender: onwer, Recipient: recipient, Amount: sdk.NewCoin(denom, sdkmath.NewInt(1))},
		// 	supply:  9,
		// },
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := srv.ConvertInsToToken(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetToken(ctx, denom)
				require.True(t, found)
				require.Equal(t, uint64(tc.supply), rst.Supply)
			}
		})
	}
}
