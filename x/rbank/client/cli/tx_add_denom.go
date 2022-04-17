package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/rbank/types"
)

var _ = strconv.Itoa(0)

func CmdAddDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-denom [acc-address-prefix] [val-address-prefix] [metadata-path]",
		Short: "Add metadata and addressPrefix",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Broadcast message add_denom with an denom_metadata which can be given through a metadata JSON file.
Example:
$ %s tx rbank add-denom cosmos cosmosvaloper path/to/metadata.json  --from mykey

Where metadata.json could be like this:

{
	"description": "The native staking token of the StaFiHub.",
	"denom_units": [
        {
			"denom": "ufis",
          	"exponent": 0,
          	"aliases": [
            	"microfis"
          	]
        },
        {
          "denom": "mfis",
          "exponent": 3,
          "aliases": [
            "millifis"
          ]
        },
        {
          "denom": "fis",
          "exponent": 6,
          "aliases": []
        }
      ],
      "base": "ufis",
      "display": "fis",
      "name": "",
      "symbol": ""
}
`, version.AppName)),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAccAddressPrefix := args[0]
			argValAddressPrefix := args[1]
			argMetadataPath := args[2]

			contents, err := os.ReadFile(argMetadataPath)
			if err != nil {
				return err
			}
			md := banktypes.Metadata{}
			err = json.Unmarshal(contents, &md)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddDenom(
				clientCtx.GetFromAddress().String(),
				argAccAddressPrefix,
				argValAddressPrefix,
				md,
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
