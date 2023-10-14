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

// profilesCmd represents the profiles command
var profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "List all profiles",
	Run: func(cmd *cobra.Command, args []string) {
		cli := goinvestec.NewClient(context.Background(),
			viper.GetString("apiURL"),
			viper.GetString("clientID"),
			viper.GetString("secret"),
			viper.GetString("apiKey"),
		)

		profiles, err := cli.GetProfiles(context.Background())
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
