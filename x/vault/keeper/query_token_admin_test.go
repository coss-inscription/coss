package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "coss/testutil/keeper"
	"coss/testutil/nullify"
	"coss/x/vault/types"
)

func TestTokenAdminQuery(t *testing.T) {
	keeper, ctx := keepertest.VaultKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestTokenAdmin(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetTokenAdminRequest
		response *types.QueryGetTokenAdminResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTokenAdminRequest{},
			response: &types.QueryGetTokenAdminResponse{TokenAdmin: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TokenAdmin(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
