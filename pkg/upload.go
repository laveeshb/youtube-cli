package pkg

import (
	"fmt"
	"os"
	"time"

	"github.com/laveeshb/youtube-cli/internal/api"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload <file>",
	Short: "Upload a video to YouTube",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath := args[0]
		if _, err := os.Stat(filePath); err != nil {
			return fmt.Errorf("file not found: %s", filePath)
		}

		title, _ := cmd.Flags().GetString("title")
		description, _ := cmd.Flags().GetString("description")
		tags, _ := cmd.Flags().GetStringSlice("tags")
		thumbnail, _ := cmd.Flags().GetString("thumbnail")
		privacy, _ := cmd.Flags().GetString("privacy")
		scheduleStr, _ := cmd.Flags().GetString("schedule")

		validPrivacy := map[string]bool{"public": true, "private": true, "unlisted": true}
		if !validPrivacy[privacy] {
			return fmt.Errorf("invalid --privacy %q: must be public, private, or unlisted", privacy)
		}

		var scheduleAt *time.Time
		if scheduleStr != "" {
			t, err := time.Parse(time.RFC3339, scheduleStr)
			if err != nil {
				return fmt.Errorf("invalid --schedule: use RFC3339 format, e.g. 2026-04-10T15:00:00Z")
			}
			if t.Before(time.Now()) {
				return fmt.Errorf("--schedule time must be in the future")
			}
			if privacy == "public" {
				return fmt.Errorf("cannot use --privacy=public with --schedule; scheduled videos must start as private")
			}
			scheduleAt = &t
		}

		client, err := newClient(cmd)
		if err != nil {
			return err
		}

		videoID, err := client.Upload(cmd.Context(), api.UploadParams{
			FilePath:    filePath,
			Title:       title,
			Description: description,
			Tags:        tags,
			Thumbnail:   thumbnail,
			Privacy:     privacy,
			ScheduleAt:  scheduleAt,
		})
		if err != nil {
			return err
		}

		fmt.Printf("Uploaded: https://youtu.be/%s\n", videoID)
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
