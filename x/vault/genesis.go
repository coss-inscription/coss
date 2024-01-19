package vault

import (
	"coss/x/vault/keeper"
	"coss/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.TokenAdmin != nil {
		k.SetTokenAdmin(ctx, *genState.TokenAdmin)
	}
	// Set all the token
	for _, elem := range genState.TokenList {
		k.SetToken(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all tokenAdmin
	tokenAdmin, found := k.GetTokenAdmin(ctx)
	if found {
		genesis.TokenAdmin = &tokenAdmin
	}
	genesis.TokenList = k.GetAllToken(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
