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

var payCmd = &cobra.Command{
	Use:   "pay",
	Short: "Pay a beneficiary from the current active account",
	Run: func(cmd *cobra.Command, args []string) {
		defaultAccount := viper.GetString("defaultAccount")
		if defaultAccount == "" {
			fmt.Println("No default account selected")
			os.Exit(1)
			return
		}

		client := cli.NewClient()
		resp, err := client.PayMultiple(context.Background(), defaultAccount, []goinvestec.Payment{
			{
				BeneficiaryID:  beneficiaryID,
				Amount:         amount,
				TheirReference: theirReference,
				MyReference:    myReference,
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
	beneficiaryID  string
	myReference    string
	theirReference string
)

func init() {
	rootCmd.AddCommand(payCmd)

	payCmd.Flags().StringVarP(&beneficiaryID, "beneficiary", "b", "", "Beneficiary to pay")
	payCmd.Flags().StringVarP(&amount, "amount", "a", "", "Payment amount")
	payCmd.Flags().StringVarP(&myReference, "myReference", "m", "", "My reference")
	payCmd.Flags().StringVarP(&theirReference, "theirReference", "t", "", "Their reference")

	payCmd.MarkFlagRequired("beneficiary")
	payCmd.MarkFlagRequired("amount")
	payCmd.MarkFlagRequired("myReference")
	payCmd.MarkFlagRequired("theirReference")
}
