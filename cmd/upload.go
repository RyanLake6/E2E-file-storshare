package cmd

import (
	"E2E-file-storshare/config"
	"E2E-file-storshare/nextcloud"
	"fmt"

	"github.com/spf13/cobra"
)


func uploadCmd() *cobra.Command {
	var (
		localPath  string
		remotePath string
		err        error
	)

	var uploadCmd = &cobra.Command{
		Use:   "upload",
		Short: "Upload a file to Nextcloud.",
		Run: func(cmd *cobra.Command, args []string) {
			config.LoadConfig()
			auth := nextcloud.NewNextcloudAuth(config.GetBaseURL())
			files := nextcloud.NewNextcloudFiles(auth.BaseURL, auth.Client)
			err = files.UploadFile(localPath, remotePath, config.GetToken())
			if err != nil {
				fmt.Println("Upload failed:", err)
			}
		},
	}

	uploadCmd.Flags().StringVar(&localPath, "local-path", "", "Local file path to upload")
	uploadCmd.Flags().StringVar(&remotePath, "remote-path", "", "Remote path to upload the file")

	return uploadCmd
}


func init() {
	rootCmd.AddCommand(uploadCmd())
}
