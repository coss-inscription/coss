package vault_test

import (
	"testing"

	keepertest "coss/testutil/keeper"
	"coss/testutil/nullify"
	"coss/x/vault"
	"coss/x/vault/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TokenAdmin: &types.TokenAdmin{
			Address: "48",
		},
		TokenList: []types.Token{
			{
				Denom: "0",
			},
			{
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VaultKeeper(t)
	vault.InitGenesis(ctx, *k, genesisState)
	got := vault.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.TokenAdmin, got.TokenAdmin)
	require.ElementsMatch(t, genesisState.TokenList, got.TokenList)
	// this line is used by starport scaffolding # genesis/test/assert
}
