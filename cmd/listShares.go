package cmd

import (
	"E2E-file-storshare/config"
	"E2E-file-storshare/nextcloud"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	debugListShare bool
	allDetailsShare bool
)

var listSharesCmd = &cobra.Command{
	Use:   "list-shares",
	Short: "List all shares in Nextcloud.",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig()
		auth := nextcloud.NewNextcloudAuth(config.GetBaseURL())
		share := nextcloud.NewNextcloudShare(auth.BaseURL, auth.Client)
		err := share.ListShares(config.GetToken(), debugListShare, allDetailsShare)
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	listSharesCmd.Flags().BoolVarP(&debugListShare, "debug", "d", false, "Setting debug mode") 
	listSharesCmd.Flags().BoolVarP(&allDetailsShare, "all-details", "a", false, "Returning all shares details") 
	rootCmd.AddCommand(listSharesCmd)
}
