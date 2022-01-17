package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/version"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/sudo/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	FlagMetadata     = "metadata"
)


// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdUpdateAdmin())
cmd.AddCommand(CmdAddDenom())
// this line is used by starport scaffolding # 1

	return cmd
}

func CmdUpdateAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-admin [address]",
		Short: "Broadcast message update_admin",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAddress, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAdmin(
				clientCtx.GetFromAddress(),
				argAddress,
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

func CmdAddDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-denom",
		Short: "Broadcast message add_denom with an denom_metadata",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Broadcast message add_denom with an denom_metadata which can be given through a metadata JSON file.

Example:
$ %s tx sudo add-denom --metadata="path/to/metadata.json" --from mykey

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
`, version.AppName),
	),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			md, err := parseMetadataFlags(cmd.Flags())
			if err != nil {
				return fmt.Errorf("failed to parse metadata: %w", err)
			}

			if err := md.Validate(); err != nil {
				return err
			}

			msg := types.NewMsgAddDenom(clientCtx.GetFromAddress(), *md)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagMetadata, "", "Metadata file path")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}