package cmd

import (
	"E2E-file-storshare/config"
	"E2E-file-storshare/nextcloud"
	"E2E-file-storshare/sharing"
	"E2E-file-storshare/utils"
	"fmt"

	"github.com/spf13/cobra"
)



func createShareCmd() *cobra.Command {
	var (
		remoteSharePath string
		permissions     int
		debug           bool
		sendEmail       bool
		emailRecipient  string
		signature       string
	)

	var createShareCmd = &cobra.Command{
		Use:   "share",
		Short: "Create a share link for a file or folder in Nextcloud.",
		Run: func(cmd *cobra.Command, args []string) {
			config.LoadConfig()
			auth := nextcloud.NewNextcloudAuth(config.GetBaseURL())
			share := nextcloud.NewNextcloudShare(auth.BaseURL, auth.Client)
			shareURL, err := share.CreateShare(remoteSharePath, nextcloud.ShareTypeLink, permissions, config.GetToken(), debug)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Share link created:", shareURL)
			}
			if sendEmail {
				htmlContent := utils.BuildShareLinkHTML(shareURL, signature)
				err := sharing.SendEmail(emailRecipient, "NextCloud Link", "", htmlContent)
				if err != nil {
					fmt.Println("failed to send email: ", err)
				}
			}
		},
	}

	createShareCmd.Flags().StringVar(&remoteSharePath, "remote-path", "", "Remote path to share")
	createShareCmd.Flags().IntVar(&permissions, "permissions", 1, "Permissions for the share")
	createShareCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Setting debug mode") 
	createShareCmd.Flags().BoolVarP(&sendEmail, "send-email", "e", false, "Setting if you wish to email this share link") 
	createShareCmd.Flags().StringVar(&signature, "signature", "s", "signature for email (your name)")
	createShareCmd.Flags().StringVar(&emailRecipient, "email-recipient", "r", "email to send to")

	return createShareCmd
}


func init() {
	rootCmd.AddCommand(createShareCmd())
}
