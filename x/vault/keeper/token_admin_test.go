package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "coss/testutil/keeper"
	"coss/testutil/nullify"
	"coss/x/vault/keeper"
	"coss/x/vault/types"
)

func createTestTokenAdmin(keeper *keeper.Keeper, ctx sdk.Context) types.TokenAdmin {
	item := types.TokenAdmin{}
	keeper.SetTokenAdmin(ctx, item)
	return item
}

func TestTokenAdminGet(t *testing.T) {
	keeper, ctx := keepertest.VaultKeeper(t)
	item := createTestTokenAdmin(keeper, ctx)
	rst, found := keeper.GetTokenAdmin(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}
