package youtube

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/option"
	yt "google.golang.org/api/youtube/v3"
)

// FetchYouTubeChannelID uses the YouTube API to get the YouTube channel ID from
// a video ID.
func FetchYouTubeChannelID(videoID string) (string, error) {
	apiKey := os.Getenv("YOUTUBE_API_KEY")

	youtubeService, err := yt.NewService(
		context.Background(), option.WithAPIKey(apiKey),
	)
	if err != nil {
		return "", fmt.Errorf("creating youtube api service: %w", err)
	}

	videosListReq := youtubeService.Videos.List(
		[]string{
			"snippet",
		},
	)
	videosListReq.Id(videoID)

	videosListResp, err := videosListReq.Do()
	if err != nil {
		return "", fmt.Errorf("listing videos from youtube api: %w", err)
	}

	if len(videosListResp.Items) != 1 {
		return "", fmt.Errorf(
			"expected one video in api response, got %d", len(videosListResp.Items),
		)
	}

	return videosListResp.Items[0].Snippet.ChannelId, nil
}
