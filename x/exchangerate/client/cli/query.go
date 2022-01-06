package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"strconv"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stafiprotocol/stafihub/x/exchangerate/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group exchangerate queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdGetExchangeRate())
	cmd.AddCommand(CmdExchangeRateAll())
	cmd.AddCommand(CmdGetEraExchangeRate())
	cmd.AddCommand(CmdEraExchangeRatesByDenom())

// this line is used by starport scaffolding # 1

	return cmd
}

func CmdGetExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-exchange-rate [denom]",
		Short: "Query get_exchange_rate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetExchangeRateRequest{
				Denom: reqDenom,
			}

			res, err := queryClient.GetExchangeRate(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdExchangeRateAll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "exchange-rate-all",
		Short: "Query exchange_rate_all",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryExchangeRateAllRequest{
			}

			res, err := queryClient.ExchangeRateAll(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetEraExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-era-exchange-rate [denom] [era]",
		Short: "Query get_era_exchange_rate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]

			argEra, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryGetEraExchangeRateRequest{
				Denom: argDenom,
				Era: uint32(argEra),
			}

			res, err := queryClient.GetEraExchangeRate(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdEraExchangeRatesByDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "era-exchange-rates-by-denom [denom]",
		Short: "Query era_exchange_rates_by_denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			params := &types.QueryEraExchangeRatesByDenomRequest{
				Denom: reqDenom,
			}

			res, err := queryClient.EraExchangeRatesByDenom(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

