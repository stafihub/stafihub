package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/flags"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stafihub/stafihub/x/ledger/types"
)

func CmdListExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-exchange-rate",
		Short: "list all ExchangeRate",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryExchangeRateAllRequest{}

			res, err := queryClient.ExchangeRateAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-exchange-rate [denom]",
		Short: "shows a ExchangeRate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetExchangeRateRequest{
				Denom: args[0],
			}

			res, err := queryClient.GetExchangeRate(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowEraExchangeRate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-era-exchange-rate [denom] [era]",
		Short: "shows a EraExchangeRate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]

			argEra, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetEraExchangeRateRequest{
				Denom: argDenom,
				Era:   uint32(argEra),
			}

			res, err := queryClient.GetEraExchangeRate(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdEraExchangeRateByDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "era-exchange-rate-by-denom [denom]",
		Short: "Query EraExchangeRateByDenom",
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
