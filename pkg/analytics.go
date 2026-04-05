package pkg

import (
	"fmt"
	"text/tabwriter"

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
		period, _ := cmd.Flags().GetString("period")

		client, err := newClient(cmd)
		if err != nil {
			return err
		}

		stats, err := client.GetChannelAnalytics(cmd.Context(), period)
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
		fmt.Fprintf(w, "Period:\t%s\n", stats.Period)
		fmt.Fprintf(w, "Views:\t%d\n", stats.Views)
		fmt.Fprintf(w, "Watch time (min):\t%d\n", stats.EstimatedMinutesWatched)
		fmt.Fprintf(w, "Avg view duration (s):\t%d\n", stats.AverageViewDuration)
		fmt.Fprintf(w, "Subscribers gained:\t%d\n", stats.SubscribersGained)
		fmt.Fprintf(w, "Subscribers lost:\t%d\n", stats.SubscribersLost)
		return w.Flush()
	},
}

var analyticsVideoCmd = &cobra.Command{
	Use:   "video <video-id>",
	Short: "Show analytics for a specific video",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		period, _ := cmd.Flags().GetString("period")

		client, err := newClient(cmd)
		if err != nil {
			return err
		}

		stats, err := client.GetVideoAnalytics(cmd.Context(), args[0], period)
		if err != nil {
			return err
		}

		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
		fmt.Fprintf(w, "Video ID:\t%s\n", stats.VideoID)
		fmt.Fprintf(w, "Period:\t%s\n", stats.Period)
		fmt.Fprintf(w, "Views:\t%d\n", stats.Views)
		fmt.Fprintf(w, "Watch time (min):\t%d\n", stats.EstimatedMinutesWatched)
		fmt.Fprintf(w, "Avg view duration (s):\t%d\n", stats.AverageViewDuration)
		fmt.Fprintf(w, "Likes:\t%d\n", stats.Likes)
		fmt.Fprintf(w, "Comments:\t%d\n", stats.Comments)
		return w.Flush()
	},
}

func init() {
	analyticsCmd.AddCommand(analyticsChannelCmd)
	analyticsCmd.AddCommand(analyticsVideoCmd)
	analyticsCmd.PersistentFlags().String("period", "28d", "Time period (e.g. 7d, 28d, 90d)")
}
