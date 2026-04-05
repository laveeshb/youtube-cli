package api

import (
	"context"
	"fmt"

	"google.golang.org/api/youtube/v3"
)

func (c *Client) ListPlaylists(_ context.Context) ([]*youtube.Playlist, error) {
	var playlists []*youtube.Playlist
	pageToken := ""
	for {
		call := c.YT.Playlists.List([]string{"snippet", "contentDetails"}).
			Mine(true).
			MaxResults(50)
		if pageToken != "" {
			call.PageToken(pageToken)
		}
		resp, err := call.Do()
		if err != nil {
			return nil, fmt.Errorf("listing playlists: %w", err)
		}
		playlists = append(playlists, resp.Items...)
		if resp.NextPageToken == "" {
			break
		}
		pageToken = resp.NextPageToken
	}
	return playlists, nil
}

func (c *Client) AddToPlaylist(_ context.Context, playlistID, videoID string) error {
	item := &youtube.PlaylistItem{
		Snippet: &youtube.PlaylistItemSnippet{
			PlaylistId: playlistID,
			ResourceId: &youtube.ResourceId{
				Kind:    "youtube#video",
				VideoId: videoID,
			},
		},
	}
	_, err := c.YT.PlaylistItems.Insert([]string{"snippet"}, item).Do()
	if err != nil {
		return fmt.Errorf("adding to playlist: %w", err)
	}
	return nil
}
