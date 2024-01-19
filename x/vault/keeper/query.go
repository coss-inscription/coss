package keeper

import (
	"coss/x/vault/types"
)

var _ types.QueryServer = Keeper{}
