package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "e2e-file-storshare-cli",
	Short: "E2E file storshare CLI is a tool for managing and sharing Nextcloud files with E2E encryption.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
