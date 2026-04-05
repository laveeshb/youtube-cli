package pkg

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yt",
	Short: "A CLI tool to manage your YouTube channel",
	Long:  `youtube-cli lets you upload videos, manage metadata, schedule publishing, and view analytics — all from the terminal.`,
}

func Execute(binaryName string) error {
	rootCmd.Use = binaryName
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(publishCmd)
	rootCmd.AddCommand(analyticsCmd)
	rootCmd.AddCommand(playlistCmd)
}
