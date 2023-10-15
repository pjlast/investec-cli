package cmd

import (
	"context"
	"fmt"

	goinvestec "github.com/pjlast/go-investec"
	"github.com/pjlast/investec-cli/cli"
	"github.com/spf13/cobra"
)

var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "Transfer money from one account to another",
	Run: func(cmd *cobra.Command, args []string) {
		client := cli.NewClient()
		resp, err := client.TransferMultiple(context.Background(), fromAccountID, profileID, []goinvestec.Transfer{
			{
				BeneficiaryAccountID: toAccountID,
				Amount:               amount,
				TheirReference:       reference,
				MyReference:          reference,
			},
		})

		if err != nil {
			fmt.Println("Something went wrong:", err)
		} else {
			fmt.Println(fmt.Sprintf("%+v", resp))
		}
	},
}

var (
	profileID     string
	fromAccountID string
	toAccountID   string
	amount        string
	reference     string
)

func init() {
	rootCmd.AddCommand(transferCmd)

	transferCmd.Flags().StringVarP(&profileID, "profile", "p", "", "Profile of the account to transfer from")
	transferCmd.Flags().StringVarP(&fromAccountID, "from", "f", "", "Account to transfer from")
	transferCmd.Flags().StringVarP(&toAccountID, "to", "t", "", "Account to transfer to")
	transferCmd.Flags().StringVarP(&amount, "amount", "a", "", "Amount to transfer")
	transferCmd.Flags().StringVarP(&reference, "reference", "r", "", "Reference for the transfer")

	transferCmd.MarkFlagRequired("profile")
	transferCmd.MarkFlagRequired("from")
	transferCmd.MarkFlagRequired("to")
	transferCmd.MarkFlagRequired("amount")
	transferCmd.MarkFlagRequired("reference")
}
