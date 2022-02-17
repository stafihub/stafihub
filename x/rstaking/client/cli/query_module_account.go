package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rstaking/types"
)

var _ = strconv.Itoa(0)

func CmdModuleAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "module-account",
		Short: "query module account",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			res := types.QueryModuleAccountResponse{
				ModuleAccount: authTypes.NewModuleAddress(types.ModuleName).String(),
			}
			return clientCtx.PrintProto(&res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
