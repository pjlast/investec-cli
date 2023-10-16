package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/pjlast/investec-cli/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setAccoutCmd sets the active account for the CLI.
var setAccountCmd = &cobra.Command{
	Use:   "setAccount [accountID]",
	Short: "Sets the default active account for all CLI commands.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: setAccount [accountID]")
			os.Exit(1)
			return
		}

		client := cli.NewClient()

		accts, err := client.GetAccounts(context.Background())
		if err != nil {
			fmt.Println("Error while fetching accounts:", err)
		}

		for _, acct := range accts {
			if acct.ID == args[0] {
				viper.Set("defaultAccount", acct.ID)
				viper.Set("defaultProfile", acct.ProfileID)
				viper.WriteConfig()
				return
			}
		}

		fmt.Println("Account with ID " + args[0] + " does not exist")
		os.Exit(1)
	},
}

var setAccountID string

func init() {
	rootCmd.AddCommand(setAccountCmd)
}
