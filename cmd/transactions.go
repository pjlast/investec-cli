package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	goinvestec "github.com/pjlast/go-investec"
	"github.com/pjlast/investec-cli/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var transactionsCmd = &cobra.Command{
	Use:   "transactions",
	Short: "Get transactions for the current active account",
	Run: func(cmd *cobra.Command, args []string) {
		defaultAccount := viper.GetString("defaultAccount")
		if defaultAccount == "" {
			fmt.Println("No active account selected")
			os.Exit(1)
			return
		}

		client := cli.NewClient()

		txs, err := client.GetAccountTransactions(context.Background(), defaultAccount, goinvestec.GetTransactionsOpts{})
		if err != nil {
			fmt.Println("Error while fetching transactions:", err)
			os.Exit(1)
			return
		}

		out, err := json.Marshal(txs)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
			return
		}

		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(transactionsCmd)
}
