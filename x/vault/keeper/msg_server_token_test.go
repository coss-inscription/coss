package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "coss/testutil/keeper"
	"coss/x/vault/keeper"
	"coss/x/vault/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestTokenMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.VaultKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	owner := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateToken{
			Owner:       owner,
			Denom: strconv.Itoa(i),
			Symbol:      strconv.Itoa(i),
			Url:         strconv.Itoa(i),
			MaxSupply:   uint64(i),
			Decimals:    int32(i),
			Description: strconv.Itoa(i),
		}
		_, err := srv.CreateToken(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetToken(ctx,
			expected.Denom,
		)
		require.True(t, found)
		require.Equal(t, expected.Owner, rst.Owner)
		require.Equal(t, expected.Denom, rst.Denom)
		require.Equal(t, expected.Symbol, rst.Symbol)
		require.Equal(t, expected.Url, rst.Url)
		require.Equal(t, expected.MaxSupply, rst.MaxSupply)
		require.Equal(t, expected.Decimals, rst.Decimals)
		require.Equal(t, expected.Description, rst.Description)
		require.Equal(t, false, rst.Audited)
		require.Equal(t, uint64(0), rst.Supply)
	}
}

func TestTokenMsgServerUpdate(t *testing.T) {
	owner := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateToken
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateToken{
				Owner:       owner,
				Denom: strconv.Itoa(0),
				Symbol:      strconv.Itoa(1),
				Url:         strconv.Itoa(1),
				MaxSupply:   uint64(2),
				Decimals:    int32(1),
				Description: strconv.Itoa(1),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateToken{Owner: "B",
				Denom: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateToken{Owner: owner,
				Denom: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.VaultKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			token := &types.MsgCreateToken{Owner: owner,
				Denom:       strconv.Itoa(0),
				Symbol:      strconv.Itoa(0),
				Url:         strconv.Itoa(0),
				MaxSupply:   uint64(1),
				Decimals:    int32(0),
				Description: strconv.Itoa(0),
			}
			expected := &types.MsgCreateToken{Owner: owner,
				Denom: strconv.Itoa(0),
				Symbol:      strconv.Itoa(1),
				Url:         strconv.Itoa(1),
				MaxSupply:   uint64(2),
				Decimals:    int32(1),
				Description: strconv.Itoa(1),
			}
			_, err := srv.CreateToken(wctx, token)
			require.NoError(t, err)

			_, err = srv.UpdateToken(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetToken(ctx,
					expected.Denom,
				)
				require.True(t, found)
				require.Equal(t, expected.Owner, rst.Owner)
				require.Equal(t, expected.Denom, rst.Denom)
				require.Equal(t, expected.Symbol, rst.Symbol)
				require.Equal(t, expected.Url, rst.Url)
				require.Equal(t, expected.MaxSupply, rst.MaxSupply)
				require.Equal(t, expected.Decimals, rst.Decimals)
				require.Equal(t, expected.Description, rst.Description)
				require.Equal(t, false, rst.Audited)
				require.Equal(t, uint64(0), rst.Supply)
			}
		})
	}
}
