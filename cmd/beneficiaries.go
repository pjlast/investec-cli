package cmd

import (
	"context"
	"fmt"

	"github.com/pjlast/investec-cli/cli"
	"github.com/spf13/cobra"
)

// beneficiariesCmd represents the accounts command
var beneficiariesCmd = &cobra.Command{
	Use:   "beneficiaries",
	Short: "Fetch a list of all beneficiaries",
	Run: func(cmd *cobra.Command, args []string) {
		client := cli.NewClient()

		beneficiaries, err := client.GetBeneficiaries(context.Background())
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
