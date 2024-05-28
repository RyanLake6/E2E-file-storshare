package cmd

import (
	"E2E-file-storshare/config"
	"E2E-file-storshare/nextcloud"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	remoteSharePath string
	permissions     int
	debugShare      bool
)

var createShareCmd = &cobra.Command{
	Use:   "share",
	Short: "Create a share link for a file or folder in Nextcloud.",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig()
		auth := nextcloud.NewNextcloudAuth(config.GetBaseURL())
		share := nextcloud.NewNextcloudShare(auth.BaseURL, auth.Client)
		shareURL, err := share.CreateShare(remoteSharePath, nextcloud.ShareTypeLink, permissions, config.GetToken(), debugShare)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Share link created:", shareURL)
		}
	},
}

func init() {
	createShareCmd.Flags().StringVar(&remoteSharePath, "remote-path", "", "Remote path to share")
	createShareCmd.Flags().IntVar(&permissions, "permissions", 1, "Permissions for the share")
	createShareCmd.Flags().BoolVarP(&debugShare, "debug", "d", false, "Setting debug mode") 
	rootCmd.AddCommand(createShareCmd)
}
