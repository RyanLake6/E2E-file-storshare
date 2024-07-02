package cmd

import (
	"E2E-file-storshare/config"
	"E2E-file-storshare/nextcloud"
	"fmt"

	"github.com/spf13/cobra"
)

func listSharesCmd() *cobra.Command {
	var (
		debug bool
		allDetails bool
	)

	var listSharesCmd = &cobra.Command{
		Use:   "list-shares",
		Short: "List all shares in Nextcloud.",
		Run: func(cmd *cobra.Command, args []string) {
			config.LoadConfig()
			auth := nextcloud.NewNextcloudAuth(config.GetBaseURL())
			share := nextcloud.NewNextcloudShare(auth.BaseURL, auth.Client)
			err := share.ListShares(config.GetToken(), debug, allDetails)
			if err != nil {
				fmt.Println("Error:", err)
			}
		},
	}

	listSharesCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Setting debug mode") 
	listSharesCmd.Flags().BoolVarP(&allDetails, "all-details", "a", false, "Returning all shares details") 

	return listSharesCmd
}



func init() {
	
	rootCmd.AddCommand(listSharesCmd())
}
