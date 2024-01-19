package keeper

import (
	"context"

	"coss/x/vault/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AuditToken(goCtx context.Context, msg *types.MsgAuditToken) (*types.MsgAuditTokenResponse, error) {
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
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "token already audited")
	}

	if !k.IsTokenAdmin(ctx, msg.TokenAdmin) {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect token admin")
	}

	var token = types.Token{
		Owner:       valFound.Owner,
		Denom:       valFound.Denom,
		Description: valFound.Description,
		Symbol:      valFound.Symbol,
		Decimals:    valFound.Decimals,
		Url:         valFound.Url,
		MaxSupply:   valFound.MaxSupply,
		Supply:      valFound.Supply,
		Audited:     msg.Pass,
	}

	k.SetToken(ctx, token)

	return &types.MsgAuditTokenResponse{}, nil
}
