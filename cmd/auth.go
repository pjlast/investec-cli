/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Update the authorization config",
	Run: func(cmd *cobra.Command, args []string) {
		if clientID != "" {
			viper.Set("clientID", clientID)
		}
		if secret != "" {
			viper.Set("secret", secret)
		}
		if apiKey != "" {
			viper.Set("apiKey", apiKey)
		}
		if apiURL != "" {
			viper.Set("apiURL", apiURL)
		}

		if err := viper.WriteConfig(); err != nil {
			fmt.Println("Error writing config:", err)
			return
		}
		fmt.Println("Config updated")
	},
}

var (
	clientID string
	secret   string
	apiKey   string
	apiURL   string
)

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.Flags().StringVarP(&clientID, "clientID", "i", "", "Client ID")
	authCmd.Flags().StringVarP(&secret, "secret", "s", "", "Secret")
	authCmd.Flags().StringVarP(&apiKey, "apiKey", "k", "", "API Key")
	authCmd.Flags().StringVarP(&apiURL, "apiURL", "u", "", "API URL")
}
