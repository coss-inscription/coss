package keeper

import (
	"coss/x/vault/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTokenAdmin set tokenAdmin in the store
func (k Keeper) SetTokenAdmin(ctx sdk.Context, tokenAdmin types.TokenAdmin) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenAdminKey))
	b := k.cdc.MustMarshal(&tokenAdmin)
	store.Set([]byte{0}, b)
}

// GetTokenAdmin returns tokenAdmin
func (k Keeper) GetTokenAdmin(ctx sdk.Context) (val types.TokenAdmin, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TokenAdminKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
