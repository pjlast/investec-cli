package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

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
			fmt.Println("Error:", err)
			os.Exit(1)
			return
		}

		out, err := json.Marshal(beneficiaries)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
			return
		}

		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(beneficiariesCmd)
}
