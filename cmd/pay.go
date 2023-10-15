package cmd

import (
	"context"
	"fmt"

	goinvestec "github.com/pjlast/go-investec"
	"github.com/pjlast/investec-cli/cli"
	"github.com/spf13/cobra"
)

var payCmd = &cobra.Command{
	Use:   "pay",
	Short: "Pay a beneficiary",
	Run: func(cmd *cobra.Command, args []string) {
		client := cli.NewClient()
		resp, err := client.PayMultiple(context.Background(), fromAccountID, []goinvestec.Payment{
			{
				BeneficiaryID:  beneficiaryID,
				Amount:         amount,
				TheirReference: theirReference,
				MyReference:    myReference,
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
	beneficiaryID  string
	myReference    string
	theirReference string
)

func init() {
	rootCmd.AddCommand(payCmd)

	payCmd.Flags().StringVarP(&fromAccountID, "from", "f", "", "Account to make payment from")
	payCmd.Flags().StringVarP(&beneficiaryID, "beneficiary", "b", "", "Beneficiary to pay")
	payCmd.Flags().StringVarP(&amount, "amount", "a", "", "Payment amount")
	payCmd.Flags().StringVarP(&myReference, "myReference", "m", "", "My reference")
	payCmd.Flags().StringVarP(&theirReference, "theirReference", "t", "", "Their reference")

	payCmd.MarkFlagRequired("from")
	payCmd.MarkFlagRequired("beneficiary")
	payCmd.MarkFlagRequired("amount")
	payCmd.MarkFlagRequired("myReference")
	payCmd.MarkFlagRequired("theirReference")
}
