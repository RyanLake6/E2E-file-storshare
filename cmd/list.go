package cmd

import (
	"fmt"

	"E2E-file-storshare/config"
	"E2E-file-storshare/nextcloud"

	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	var (
		remotePath string
		debug  bool
		allDetails bool
	)

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List files in a Nextcloud directory.",
		Run: func(cmd *cobra.Command, args []string) {
			config.LoadConfig()
			auth := nextcloud.NewNextcloudAuth(config.GetBaseURL())
			files := nextcloud.NewNextcloudFiles(auth.BaseURL, auth.Client)
	
			config.LoadConfig()
			err := files.ListFiles(remotePath, config.GetToken(), allDetails)
			if err != nil {
				fmt.Println("Error:", err)
			}
	
			if debug {
				fmt.Println("remote path: ", remotePath)
				fmt.Println("token is: ", config.GetToken())
				fmt.Println("base url is: ", config.GetBaseURL())
			}
		},
	}

	listCmd.Flags().StringVar(&remotePath, "remote-path", "", "Remote path to list files")
	listCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Setting debug mode")
	listCmd.Flags().BoolVarP(&allDetails, "all-details", "a", false, "Specifying if all related should be shown")

	return listCmd
}


func init() {
	rootCmd.AddCommand(listCmd())
}
