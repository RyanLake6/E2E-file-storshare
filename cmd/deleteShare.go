package cmd

import (
	"E2E-file-storshare/config"
	"E2E-file-storshare/nextcloud"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	shareID          string
	debugDeleteShare bool
)

var deleteShareCmd = &cobra.Command{
	Use:   "delete-share",
	Short: "Delete a share link in Nextcloud.",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig()
		auth := nextcloud.NewNextcloudAuth(config.GetBaseURL())
		share := nextcloud.NewNextcloudShare(auth.BaseURL, auth.Client)
		msg, err := share.DeleteShare(shareID, config.GetToken(), debugDeleteShare)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Share link %s deleted. Message: %s", shareID, msg)
		}
	},
}

func init() {
	deleteShareCmd.Flags().StringVar(&shareID, "share-id", "", "ID of the share to delete")
	deleteShareCmd.Flags().BoolVarP(&debugDeleteShare, "debug", "d", false, "Setting debug mode") 
	rootCmd.AddCommand(deleteShareCmd)
}
