package account

import (
	"context"
	"fmt"

	goinvestec "github.com/pjlast/go-investec"
	"github.com/pjlast/investec-cli/cli"
	"github.com/spf13/cobra"
)

var transactionsCmd = &cobra.Command{
	Use:   "transactions [account ID]",
	Short: "Get account transactions",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("invalid args. Only [accountId] required.")
			return
		}

		client := cli.NewClient()

		txs, err := client.GetAccountTransactions(context.Background(), args[0], goinvestec.GetTransactionsOpts{})
		if err != nil {
			fmt.Println("Error while fetching accounts:", err)
		}

		for _, t := range txs {
			fmt.Println(t.Description, t.TransactionDate, t.Amount)
		}
	},
}

func init() {
	AccountCmd.AddCommand(transactionsCmd)
}
