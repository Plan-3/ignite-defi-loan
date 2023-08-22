package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"loan/x/loan/types"
)

var _ = strconv.Itoa(0)

func CmdTokenMint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token-mint [coin] [denom] [denomAmount]",
		Short: "Broadcast message token-mint",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			denom := args[1]
			denomAmount, err := strconv.ParseInt(args[2], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTokenMint(
				clientCtx.GetFromAddress().String(),
				argAmount,
				denom,
				denomAmount,
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
