package cli

import (
	"strconv"

	"coss/x/vault/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdConvertInsToToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "convert-ins-to-token [amount] [recipient]",
		Short: "Broadcast message ConvertInsToToken",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			argRecipient := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgConvertInsToToken(
				clientCtx.GetFromAddress().String(),
				argAmount,
				argRecipient,
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
