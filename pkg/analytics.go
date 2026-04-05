package pkg

import (
	"fmt"

	"github.com/spf13/cobra"
)

var analyticsCmd = &cobra.Command{
	Use:   "analytics",
	Short: "View channel and video analytics",
}

var analyticsChannelCmd = &cobra.Command{
	Use:   "channel",
	Short: "Show channel-level analytics",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Analytics channel: not yet implemented")
		return nil
	},
}

var analyticsVideoCmd = &cobra.Command{
	Use:   "video <video-id>",
	Short: "Show analytics for a specific video",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Analytics video %s: not yet implemented\n", args[0])
		return nil
	},
}

func init() {
	analyticsCmd.AddCommand(analyticsChannelCmd)
	analyticsCmd.AddCommand(analyticsVideoCmd)
	analyticsCmd.PersistentFlags().String("period", "28d", "Time period (e.g. 7d, 28d, 90d)")
}
