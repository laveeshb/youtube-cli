package api

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/youtube/v3"
)

func (c *Client) Publish(ctx context.Context, videoID string, scheduleAt *time.Time) error {
	resp, err := c.YT.Videos.List([]string{"status"}).Id(videoID).Do()
	if err != nil {
		return fmt.Errorf("fetching video: %w", err)
	}
	if len(resp.Items) == 0 {
		return fmt.Errorf("video %q not found", videoID)
	}

	video := resp.Items[0]
	if scheduleAt != nil {
		video.Status.PrivacyStatus = "private"
		video.Status.PublishAt = scheduleAt.UTC().Format(time.RFC3339)
	} else {
		video.Status.PrivacyStatus = "public"
		video.Status.PublishAt = ""
	}

	_, err = c.YT.Videos.Update([]string{"status"}, &youtube.Video{
		Id:     videoID,
		Status: video.Status,
	}).Do()
	if err != nil {
		return fmt.Errorf("updating video status: %w", err)
	}
	return nil
}
