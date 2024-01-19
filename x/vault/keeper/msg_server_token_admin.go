package keeper

import (
	"context"

	"coss/x/vault/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateTokenAdmin(goCtx context.Context, msg *types.MsgCreateTokenAdmin) (*types.MsgCreateTokenAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetTokenAdmin(ctx)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var tokenAdmin = types.TokenAdmin{
		Creator: msg.Creator,
		Address: msg.Address,
	}

	k.SetTokenAdmin(
		ctx,
		tokenAdmin,
	)
	return &types.MsgCreateTokenAdminResponse{}, nil
}

func (k msgServer) UpdateTokenAdmin(goCtx context.Context, msg *types.MsgUpdateTokenAdmin) (*types.MsgUpdateTokenAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetTokenAdmin(ctx)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var tokenAdmin = types.TokenAdmin{
		Creator: msg.Creator,
		Address: msg.Address,
	}

	k.SetTokenAdmin(ctx, tokenAdmin)

	return &types.MsgUpdateTokenAdminResponse{}, nil
}
