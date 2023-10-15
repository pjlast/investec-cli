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

// profilesCmd represents the profiles command
var profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "List all profiles",
	Run: func(cmd *cobra.Command, args []string) {
		client := cli.NewClient()

		profiles, err := client.GetProfiles(context.Background())
		if err != nil {
			fmt.Println("Error while fetching accounts:", err)
		}

		for _, p := range profiles {
			fmt.Println(p.ID, p.Name, p.Default)
		}
	},
}

func init() {
	rootCmd.AddCommand(profilesCmd)
}
