package cmd

import (
	"fmt"

	"E2E-file-storshare/config"
	"E2E-file-storshare/nextcloud"

	"github.com/spf13/cobra"
)

var (
	remotePath string
	debugList  bool
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List files in a Nextcloud directory.",
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig()
		auth := nextcloud.NewNextcloudAuth(config.GetBaseURL())
		files := nextcloud.NewNextcloudFiles(auth.BaseURL, auth.Client)

		config.LoadConfig()
		err := files.ListFiles(remotePath, config.GetToken())
		if err != nil {
			fmt.Println("Error:", err)
		}

		if debugList {
			fmt.Println("remote path: ", remotePath)
			fmt.Println("token is: ", config.GetToken())
			fmt.Println("base url is: ", config.GetBaseURL())
		}
	},
}

func init() {
	listCmd.Flags().StringVar(&remotePath, "remote-path", "", "Remote path to list files")
	listCmd.Flags().BoolVarP(&debugList, "debug", "d", false, "Setting debug mode") 
	rootCmd.AddCommand(listCmd)
}
