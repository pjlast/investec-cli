package account

import "github.com/spf13/cobra"

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Commands related to a specific account",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var accountID string

func init() {
	AccountCmd.Flags().StringVarP(&accountID, "account", "a", "", "Account ID")
	AccountCmd.MarkFlagRequired("account")
}
