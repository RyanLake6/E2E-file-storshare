package cmd

import (
	auth "E2E-file-storshare/nextcloud"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	baseURL  string
	username string
	password string
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to your Nextcloud account.",
	Run: func(cmd *cobra.Command, args []string) {
		NextcloudAuth := auth.NewNextcloudAuth(baseURL)
		err := NextcloudAuth.Login(username, password)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Login successful")
		}
	},
}

func init() {
	loginCmd.Flags().StringVar(&baseURL, "base-url", "", "Nextcloud base URL")
	loginCmd.Flags().StringVar(&username, "username", "", "Nextcloud username")
	loginCmd.Flags().StringVar(&password, "password", "", "Nextcloud password")
	rootCmd.AddCommand(loginCmd)
}
