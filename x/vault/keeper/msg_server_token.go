package keeper

import (
	"context"

	"coss/x/vault/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetToken(
		ctx,
		msg.Denom,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "token already exists")
	}

	var token = types.Token{
		Owner:       msg.Owner,
		Denom:       msg.Denom,
		Description: msg.Description,
		Symbol:      msg.Symbol,
		Decimals:    msg.Decimals,
		Url:         msg.Url,
		MaxSupply:   msg.MaxSupply,
	}

	k.SetToken(
		ctx,
		token,
	)
	return &types.MsgCreateTokenResponse{}, nil
}

func (k msgServer) UpdateToken(goCtx context.Context, msg *types.MsgUpdateToken) (*types.MsgUpdateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetToken(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "token not found")
	}

	if valFound.Audited {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "token audited, can not update")
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Owner != valFound.Owner && !k.IsTokenAdmin(ctx, msg.Owner) {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var token = types.Token{
		Owner:       valFound.Owner,
		Denom:       valFound.Denom,
		Description: msg.Description,
		Symbol:      msg.Symbol,
		Decimals:    msg.Decimals,
		Url:         msg.Url,
		MaxSupply:   msg.MaxSupply,
		Supply:      valFound.Supply,
		Audited:     valFound.Audited,
	}

	k.SetToken(ctx, token)

	return &types.MsgUpdateTokenResponse{}, nil
}

func (k msgServer) IsTokenAdmin(goCtx context.Context, address string) bool {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// compare address with token-admin
	valFound, isFound := k.GetTokenAdmin(ctx)
	if !isFound {
		return false
	}
	return valFound.Address == address
}
