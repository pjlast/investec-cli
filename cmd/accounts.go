/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

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
			fmt.Println("Error:", err)
			os.Exit(1)
			return
		}

		out, err := json.Marshal(accts)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
			return
		}

		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(accountsCmd)
}
