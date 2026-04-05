package pkg

import (
	"fmt"

	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish <video-id>",
	Short: "Publish a private or scheduled video",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Publish %s: not yet implemented\n", args[0])
		return nil
	},
}

func init() {
	publishCmd.Flags().String("schedule", "", "Schedule publish time (RFC3339, e.g. 2026-04-10T15:00:00Z)")
}
