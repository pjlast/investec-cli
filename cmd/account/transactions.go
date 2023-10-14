package account

import (
	"context"
	"fmt"

	goinvestec "github.com/pjlast/go-investec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var transactionsCmd = &cobra.Command{
	Use:   "transactions [account ID]",
	Short: "Get account transactions",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("invalid args. Only [accountId] required.")
			return
		}

		cli := goinvestec.NewClient(context.Background(),
			viper.GetString("apiURL"),
			viper.GetString("clientID"),
			viper.GetString("secret"),
			viper.GetString("apiKey"),
		)

		txs, err := cli.GetAccountTransactions(context.Background(), args[0])
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
