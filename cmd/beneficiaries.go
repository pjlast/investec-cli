package cmd

import (
	"context"
	"fmt"

	goinvestec "github.com/pjlast/go-investec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// beneficiariesCmd represents the accounts command
var beneficiariesCmd = &cobra.Command{
	Use:   "beneficiaries",
	Short: "Fetch a list of all beneficiaries",
	Run: func(cmd *cobra.Command, args []string) {
		cli := goinvestec.NewClient(context.Background(),
			viper.GetString("apiURL"),
			viper.GetString("clientID"),
			viper.GetString("secret"),
			viper.GetString("apiKey"),
		)

		beneficiaries, err := cli.GetBeneficiaries(context.Background())
		if err != nil {
			fmt.Println("Error while fetching beneficiaries:", err)
		}

		for _, b := range beneficiaries {
			fmt.Println(b.ID, b.Name, b.Bank, b.AccountNumber)
		}
	},
}

func init() {
	rootCmd.AddCommand(beneficiariesCmd)
}
