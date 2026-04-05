package api

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ChannelAnalytics struct {
	Period                  string
	Views                   int64
	EstimatedMinutesWatched int64
	AverageViewDuration     int64
	SubscribersGained       int64
	SubscribersLost         int64
}

type VideoAnalytics struct {
	VideoID                 string
	Period                  string
	Views                   int64
	EstimatedMinutesWatched int64
	AverageViewDuration     int64
	Likes                   int64
	Comments                int64
}

func parsePeriod(period string) (startDate, endDate string, err error) {
	period = strings.TrimSpace(period)
	if !strings.HasSuffix(period, "d") {
		return "", "", fmt.Errorf("invalid period %q — use format like 7d, 28d, 90d", period)
	}
	days, err := strconv.Atoi(strings.TrimSuffix(period, "d"))
	if err != nil || days <= 0 {
		return "", "", fmt.Errorf("invalid period %q — use format like 7d, 28d, 90d", period)
	}
	now := time.Now().UTC()
	endDate = now.Format("2006-01-02")
	startDate = now.AddDate(0, 0, -days).Format("2006-01-02")
	return startDate, endDate, nil
}

func (c *Client) GetChannelAnalytics(_ context.Context, period string) (*ChannelAnalytics, error) {
	startDate, endDate, err := parsePeriod(period)
	if err != nil {
		return nil, err
	}

	resp, err := c.Analytics.Reports.Query().
		Ids("channel==MINE").
		StartDate(startDate).
		EndDate(endDate).
		Metrics("views,estimatedMinutesWatched,averageViewDuration,subscribersGained,subscribersLost").
		Do()
	if err != nil {
		return nil, fmt.Errorf("fetching channel analytics: %w", err)
	}

	result := &ChannelAnalytics{Period: period}
	if len(resp.Rows) > 0 {
		row := resp.Rows[0]
		result.Views = toInt(row[0])
		result.EstimatedMinutesWatched = toInt(row[1])
		result.AverageViewDuration = toInt(row[2])
		result.SubscribersGained = toInt(row[3])
		result.SubscribersLost = toInt(row[4])
	}
	return result, nil
}

func (c *Client) GetVideoAnalytics(_ context.Context, videoID, period string) (*VideoAnalytics, error) {
	startDate, endDate, err := parsePeriod(period)
	if err != nil {
		return nil, err
	}

	resp, err := c.Analytics.Reports.Query().
		Ids("channel==MINE").
		StartDate(startDate).
		EndDate(endDate).
		Metrics("views,estimatedMinutesWatched,averageViewDuration,likes,comments").
		Filters("video==" + videoID).
		Do()
	if err != nil {
		return nil, fmt.Errorf("fetching video analytics: %w", err)
	}

	result := &VideoAnalytics{VideoID: videoID, Period: period}
	if len(resp.Rows) > 0 {
		row := resp.Rows[0]
		result.Views = toInt(row[0])
		result.EstimatedMinutesWatched = toInt(row[1])
		result.AverageViewDuration = toInt(row[2])
		result.Likes = toInt(row[3])
		result.Comments = toInt(row[4])
	}
	return result, nil
}

func toInt(v interface{}) int64 {
	switch n := v.(type) {
	case float64:
		return int64(n)
	case int64:
		return n
	}
	return 0
}
