// cmd/root.go
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "contentful-cli",
	Short: "A CLI tool to fetch and store content from Contentful",
}

func Execute() error {
	return rootCmd.Execute()
}
