package cmd

import (
	"E2E-file-storshare/config"
	"E2E-file-storshare/nextcloud"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	baseURL string
	debugLogin   bool
)

var loginCmd = &cobra.Command{
    Use:   "login",
    Short: "Log in to your Nextcloud account.",
    Run: func(cmd *cobra.Command, args []string) {
        NextcloudAuth := nextcloud.NewNextcloudAuth(baseURL)
        pollResponse, err := NextcloudAuth.Login()
        if err != nil {
            fmt.Println("Error logging in:", err)
        }

		// Setting config file for persistent values
		config.SetToken(pollResponse.AppPassword)
		config.SetBaseURL(pollResponse.Server)
		err = config.SaveConfig()
		if err != nil {
			fmt.Println("Error saving config:", err)
		}

		if debugLogin {
			fmt.Println("poll response is: ", pollResponse)
			fmt.Println("token from config file is: ", config.GetToken())
		}
		
		if err == nil {
			fmt.Println("Login successful. Credentials saved.")
		}
    },
}

func init() {
    loginCmd.Flags().StringVar(&baseURL, "base-url", "", "Nextcloud base URL")
	loginCmd.Flags().BoolVarP(&debugLogin, "debug", "d", false, "Setting debug mode")
    rootCmd.AddCommand(loginCmd)
}
