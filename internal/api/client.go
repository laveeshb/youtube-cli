package api

import (
	"context"

	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"google.golang.org/api/youtubeanalytics/v2"
)

type Client struct {
	YT        *youtube.Service
	Analytics *youtubeanalytics.Service
}

func New(ctx context.Context, ts oauth2.TokenSource) (*Client, error) {
	httpClient := oauth2.NewClient(ctx, ts)

	ytSvc, err := youtube.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}

	analyticsSvc, err := youtubeanalytics.NewService(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}

	return &Client{YT: ytSvc, Analytics: analyticsSvc}, nil
}
