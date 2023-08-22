package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"loan/x/loan/types"
)

var _ = strconv.Itoa(0)

func CmdBurnToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-token [denom] [amount]",
		Short: "Broadcast message burn-token",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argAmount, err := cast.ToInt64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnToken(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argAmount,
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
