package keeper

import (
	"context"

	"coss/x/vault/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ConvertInsToToken(goCtx context.Context, msg *types.MsgConvertInsToToken) (*types.MsgConvertInsToTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valFound, isFound := k.GetToken(
		ctx,
		msg.Amount.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "token does not exist")
	}

	if !valFound.Audited {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "token not audited, can not mint")
	}

	// Checks if the the msg owner is the same as the current owner
	if msg.Owner != valFound.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if valFound.Supply+msg.Amount.Amount.Uint64() > valFound.MaxSupply {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Cannot mint more than Max Supply")
	}

	// check recipient address
	recipientAddress, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	var mintCoins sdk.Coins

	mintCoins = mintCoins.Add(msg.Amount)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return nil, err
	}
	moduleAcct := k.accountKeeper.GetModuleAddress(types.ModuleName)
	if err := k.bankKeeper.SendCoins(ctx, moduleAcct, recipientAddress, mintCoins); err != nil {
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
		Supply:      valFound.Supply + msg.Amount.Amount.Uint64(),
		Audited:     valFound.Audited,
	}

	k.SetToken(ctx, token)

	return &types.MsgConvertInsToTokenResponse{}, nil
}
