package account

import (
	"context"
	"fmt"

	goinvestec "github.com/pjlast/go-investec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var balanceCmd = &cobra.Command{
	Use: "balance [account ID]",
	Short: "Get the account balance",
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

		balance, err := cli.GetAccountBalance(context.Background(), args[0])
		if err != nil {
			fmt.Println("Error while fetching accounts:", err)
		}

		fmt.Println(balance.Currency, balance.CurrentBalance)
	},
}

func init() {
	AccountCmd.AddCommand(balanceCmd)
}
