package api

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/api/youtube/v3"
)

type UploadParams struct {
	FilePath    string
	Title       string
	Description string
	Tags        []string
	Thumbnail   string
	Privacy     string
	ScheduleAt  *time.Time
}

func (c *Client) Upload(ctx context.Context, p UploadParams) (string, error) {
	file, err := os.Open(p.FilePath)
	if err != nil {
		return "", fmt.Errorf("opening video file: %w", err)
	}
	defer file.Close()

	privacy := p.Privacy
	var publishAt string
	if p.ScheduleAt != nil {
		privacy = "private"
		publishAt = p.ScheduleAt.UTC().Format(time.RFC3339)
	}

	video := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       p.Title,
			Description: p.Description,
			Tags:        p.Tags,
		},
		Status: &youtube.VideoStatus{
			PrivacyStatus: privacy,
			PublishAt:     publishAt,
		},
	}

	fi, _ := file.Stat()
	total := fi.Size()
	var uploaded int64

	call := c.YT.Videos.Insert([]string{"snippet", "status"}, video)
	call.Media(file)
	call.ProgressUpdater(func(current, _ int64) {
		uploaded = current
		if total > 0 {
			pct := float64(uploaded) / float64(total) * 100
			fmt.Printf("\rUploading... %.1f%%", pct)
		}
	})

	result, err := call.Do()
	if err != nil {
		return "", fmt.Errorf("uploading video: %w", err)
	}
	fmt.Println() // newline after progress

	if p.Thumbnail != "" {
		if err := c.setThumbnail(ctx, result.Id, p.Thumbnail); err != nil {
			fmt.Printf("Warning: thumbnail upload failed: %v\n", err)
		}
	}

	return result.Id, nil
}

func (c *Client) setThumbnail(_ context.Context, videoID, thumbPath string) error {
	f, err := os.Open(thumbPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = c.YT.Thumbnails.Set(videoID).Media(f).Do()
	return err
}
