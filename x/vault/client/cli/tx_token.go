package cli

import (
	"coss/x/vault/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreateToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-token [denom] [description] [symbol] [decimals] [url] [max-supply]",
		Short: "Create a new Token",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexDenom := args[0]

			// Get value arguments
			argDescription := args[1]
			argSymbol := args[2]
			argDecimals, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argUrl := args[4]
			argMaxSupply, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateToken(
				clientCtx.GetFromAddress().String(),
				indexDenom,
				argDescription,
				argSymbol,
				argDecimals,
				argUrl,
				argMaxSupply,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-token [denom] [description] [symbol] [decimals] [url] [max-supply]",
		Short: "Update a Token",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexDenom := args[0]

			// Get value arguments
			argDescription := args[1]
			argSymbol := args[2]
			argDecimals, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argUrl := args[4]
			argMaxSupply, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateToken(
				clientCtx.GetFromAddress().String(),
				indexDenom,
				argDescription,
				argSymbol,
				argDecimals,
				argUrl,
				argMaxSupply,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
