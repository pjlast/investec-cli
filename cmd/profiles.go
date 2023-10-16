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

// profilesCmd represents the profiles command
var profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "List all profiles",
	Run: func(cmd *cobra.Command, args []string) {
		client := cli.NewClient()

		profiles, err := client.GetProfiles(context.Background())
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
			return
		}

		out, err := json.Marshal(profiles)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
			return
		}

		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(profilesCmd)
}
