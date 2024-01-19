package keeper

import (
	"context"

	"coss/x/vault/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TokenAdmin(goCtx context.Context, req *types.QueryGetTokenAdminRequest) (*types.QueryGetTokenAdminResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetTokenAdmin(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTokenAdminResponse{TokenAdmin: val}, nil
}
