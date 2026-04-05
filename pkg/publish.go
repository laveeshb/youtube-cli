package pkg

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish <video-id>",
	Short: "Publish a private or scheduled video",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		videoID := args[0]
		scheduleStr, _ := cmd.Flags().GetString("schedule")

		var scheduleAt *time.Time
		if scheduleStr != "" {
			t, err := time.Parse(time.RFC3339, scheduleStr)
			if err != nil {
				return fmt.Errorf("invalid --schedule: use RFC3339 format, e.g. 2026-04-10T15:00:00Z")
			}
			if t.Before(time.Now()) {
				return fmt.Errorf("--schedule time must be in the future")
			}
			scheduleAt = &t
		}

		client, err := newClient(cmd)
		if err != nil {
			return err
		}

		if err := client.Publish(cmd.Context(), videoID, scheduleAt); err != nil {
			return err
		}

		if scheduleAt != nil {
			fmt.Printf("Video %s scheduled to publish at %s\n", videoID, scheduleAt.Local().Format("2006-01-02 15:04:05"))
		} else {
			fmt.Printf("Video %s published: https://youtu.be/%s\n", videoID, videoID)
		}
		return nil
	},
}

func init() {
	publishCmd.Flags().String("schedule", "", "Schedule publish time (RFC3339, e.g. 2026-04-10T15:00:00Z)")
}
