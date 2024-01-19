package keeper

import (
	"context"

	"coss/x/vault/types"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ConvertTokenToIns(goCtx context.Context, msg *types.MsgConvertTokenToIns) (*types.MsgConvertTokenToInsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valFound, isFound := k.GetToken(ctx, msg.Amount.Denom)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "token not found")
	}
	if !valFound.Audited {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "token not audited")
	}
	if msg.Amount.Amount.GT(sdkmath.NewIntFromUint64(valFound.MaxSupply)) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "insufficient token supply")
	}

	var burnCoins sdk.Coins

	burnCoins = burnCoins.Add(msg.Amount)
	senderAddress, addrErr := sdk.AccAddressFromBech32(msg.Sender)
	if addrErr != nil {
		return nil, addrErr
	}
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, senderAddress, types.ModuleName, burnCoins)
	if err != nil {
		return nil, err
	}
	// burn coins
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnCoins)
	if err != nil {
		return nil, err
	}

	var token = types.Token{
		Owner:       valFound.Owner,
		Denom:       valFound.Denom,
		Description: valFound.Description,
		Symbol:      valFound.Symbol,
		Decimals:    valFound.Decimals,
		Url:         valFound.Url,
		MaxSupply:   valFound.MaxSupply,
		Supply:      valFound.Supply - msg.Amount.Amount.Uint64(),
		Audited:     valFound.Audited,
	}

	k.SetToken(ctx, token)

	return &types.MsgConvertTokenToInsResponse{}, nil
}
