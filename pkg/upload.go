package pkg

import (
	"fmt"

	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload <file>",
	Short: "Upload a video to YouTube",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Upload %s: not yet implemented\n", args[0])
		return nil
	},
}

func init() {
	uploadCmd.Flags().String("title", "", "Video title")
	uploadCmd.Flags().String("description", "", "Video description")
	uploadCmd.Flags().StringSlice("tags", nil, "Comma-separated tags")
	uploadCmd.Flags().String("thumbnail", "", "Path to thumbnail image")
	uploadCmd.Flags().String("privacy", "private", "Privacy status: public, private, unlisted")
	uploadCmd.Flags().String("schedule", "", "Schedule publish time (RFC3339, e.g. 2026-04-10T15:00:00Z)")
}
