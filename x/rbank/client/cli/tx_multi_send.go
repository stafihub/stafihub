package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/utils"
)

var _ = strconv.Itoa(0)

func CmdMultiSend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "multi-send [receivers-info-path]",
		Short: "Transfer fis to multi receivers in a transaction. Note that the number in the file is the number of fis not ufis",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Broadcast message multi-send which can be given through a receivers info file.
Example:
$ %s tx rbank multi-send path/to/receivers_info.txt  --from mykey

Where receivers_info.txt could be like this:

stafi1dl8mh892fltlskd8jjlrre7kav929zcrcm8h29 15
stafi17qw0g6ljqy0ktgzc94l8az5qlfdwutm9qlwghz 45
stafi10lmf97336rkydn5tk2lxdq8x2yttn2gr8agm6j 45
stafi1jtx2x3hgqer3l65mrf05nej6aeuv9pa9mgn9mc 15
stafi1qyt7zr3ck8dvq8cv2fhmkk2traeuw6hdcpq86p 15
stafi1hjmjef9rtkppu55rjpmf9ss4pwcepcrjwhrwqq 15

`, version.AppName)),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiversInfoPath := args[0]

			f, err := os.Open(argReceiversInfoPath)
			if err != nil {
				return err
			}
			defer f.Close()

			br := bufio.NewReader(f)

			rows := make([][]string, 0)
			for {
				line, _, c := br.ReadLine()
				if c == io.EOF {
					break
				}

				rows = append(rows, strings.Fields(strings.TrimSpace(string(line))))
			}
			fmt.Printf("total address: %d\n", len(rows))

			totalAmount := types.NewInt(0)
			outputs := []banktypes.Output{}
			for line, row := range rows {
				if len(row) != 2 {
					return fmt.Errorf("row %d format not right", line)
				}
				address, err := types.AccAddressFromBech32(row[0])
				if err != nil {
					return fmt.Errorf("address format not right, row %d", line)
				}
				amount, ok := types.NewIntFromString(row[1])
				if !ok {
					return fmt.Errorf("amount format not right, row %d", line)
				}
				// add decimals
				amount = amount.Mul(types.NewInt(1e6))

				totalAmount = totalAmount.Add(amount)

				outputs = append(outputs, banktypes.NewOutput(address, types.NewCoins(types.NewCoin(utils.FisDenom, amount))))
			}
			fmt.Printf("total send amount: %s ufis\n", totalAmount.String())

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			inputs := []banktypes.Input{{Address: clientCtx.GetFromAddress().String(), Coins: types.NewCoins(types.NewCoin(utils.FisDenom, totalAmount))}}

			msg := banktypes.NewMsgMultiSend(inputs, outputs)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
