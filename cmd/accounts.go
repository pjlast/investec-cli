/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	goinvestec "github.com/pjlast/go-investec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// accountsCmd represents the accounts command
var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Fetch a list of all accounts",
	Run: func(cmd *cobra.Command, args []string) {
		cli := goinvestec.NewClient(context.Background(),
			viper.GetString("apiURL"),
			viper.GetString("clientID"),
			viper.GetString("secret"),
			viper.GetString("apiKey"),
		)

		accts, err := cli.GetAccounts(context.Background())
		if err != nil {
			fmt.Println("Error while fetching accounts:", err)
		}

		for _, acct := range accts {
			fmt.Println(acct.ID, acct.Name, acct.Number)
		}
	},
}

func init() {
	rootCmd.AddCommand(accountsCmd)
}
