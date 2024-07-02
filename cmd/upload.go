package cmd

import (
	"E2E-file-storshare/config"
	"E2E-file-storshare/nextcloud"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	localPathUpload  string
	remotePathUpload string
	err        error
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file to Nextcloud.",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig()
		auth := nextcloud.NewNextcloudAuth(config.GetBaseURL())
		files := nextcloud.NewNextcloudFiles(auth.BaseURL, auth.Client)
		err = files.UploadFile(localPathUpload, remotePathUpload, config.GetToken())
		if err != nil {
			fmt.Println("Upload failed:", err)
		}
	},
}

func init() {
	uploadCmd.Flags().StringVar(&localPathUpload, "local-path", "", "Local file path to upload")
	uploadCmd.Flags().StringVar(&remotePathUpload, "remote-path", "", "Remote path to upload the file")
	rootCmd.AddCommand(uploadCmd)
}
