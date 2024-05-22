package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints 'Hello, world!'",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world!")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}