package cmd

import (
	"E2E-file-storshare/config"
	"E2E-file-storshare/nextcloud"
	"fmt"

	"github.com/spf13/cobra"
)


func downloadCmd() *cobra.Command {
	var (
		localPath  string
		remotePath string
		err        error
	)

	var downloadCmd = &cobra.Command{
		Use:   "download",
		Short: "download a file from Nextcloud.",
		Run: func(cmd *cobra.Command, args []string) {
			config.LoadConfig()
			auth := nextcloud.NewNextcloudAuth(config.GetBaseURL())
			files := nextcloud.NewNextcloudFiles(auth.BaseURL, auth.Client)
			err = files.DownloadFile(remotePath, localPath, config.GetToken())
			if err != nil {
				fmt.Println("Download failed:", err)
			}
		},
	}

	downloadCmd.Flags().StringVar(&localPath, "local-path", "", "Local file path to download to")
	downloadCmd.Flags().StringVar(&remotePath, "remote-path", "", "Remote path to download from")

	return downloadCmd
}


func init() {
	rootCmd.AddCommand(downloadCmd())
}
