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

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer money from the current active account to another account",
	Run: func(cmd *cobra.Command, args []string) {
		defaultAccount := viper.GetString("defaultAccount")
		defaultProfile := viper.GetString("defaultProfile")

		if defaultAccount == "" || defaultProfile == "" {
			fmt.Println("No default account selected")
			os.Exit(1)
			return
		}

		client := cli.NewClient()
		resp, err := client.TransferMultiple(context.Background(), defaultAccount, defaultProfile, []goinvestec.Transfer{
			{
				BeneficiaryAccountID: toAccountID,
				Amount:               amount,
				TheirReference:       reference,
				MyReference:          reference,
			},
		})

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
			return
		}

		out, err := json.Marshal(resp)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
			return
		}

		fmt.Println(string(out))
	},
}

var (
	toAccountID string
	amount      string
	reference   string
)

func init() {
	rootCmd.AddCommand(transferCmd)

	transferCmd.Flags().StringVarP(&toAccountID, "to", "t", "", "Account to transfer to")
	transferCmd.Flags().StringVarP(&amount, "amount", "a", "", "Amount to transfer")
	transferCmd.Flags().StringVarP(&reference, "reference", "r", "", "Reference for the transfer")

	transferCmd.MarkFlagRequired("to")
	transferCmd.MarkFlagRequired("amount")
	transferCmd.MarkFlagRequired("reference")
}
