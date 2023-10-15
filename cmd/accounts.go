/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/pjlast/investec-cli/cli"
	"github.com/spf13/cobra"
)

// accountsCmd represents the accounts command
var accountsCmd = &cobra.Command{
	Use:   "accounts",
	Short: "Fetch a list of all accounts",
	Run: func(cmd *cobra.Command, args []string) {
		client := cli.NewClient()

		accts, err := client.GetAccounts(context.Background())
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
