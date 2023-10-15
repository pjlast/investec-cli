package account

import (
	"context"
	"fmt"

	"github.com/pjlast/investec-cli/cli"
	"github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
	Use:   "balance [account ID]",
	Short: "Get the account balance",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("invalid args. Only [accountId] required.")
			return
		}

		client := cli.NewClient()

		balance, err := client.GetAccountBalance(context.Background(), args[0])
		if err != nil {
			fmt.Println("Error while fetching accounts:", err)
		}

		fmt.Println(balance.Currency, balance.CurrentBalance)
	},
}

func init() {
	AccountCmd.AddCommand(balanceCmd)
}
