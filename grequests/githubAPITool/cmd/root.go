package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gist cli",
	Short: "CLI to push files to gist.github.com",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// cmd here
	// },
}
