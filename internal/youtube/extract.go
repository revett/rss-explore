package youtube

import (
	"fmt"
	"net/url"
	"strings"
)

const videoIDLength = 11

// ExtractVideoID validates and extracts the YouTube video ID from a URL.
func ExtractVideoID(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("parsing string to url type: %w", err)
	}

	host := strings.TrimPrefix(parsedURL.Host, "www.")
	hostParts := strings.Split(host, ".")
	if len(hostParts) != 2 {
		return "", fmt.Errorf("invalid url host: %s", parsedURL.Host)
	}

	domain := hostParts[0]
	extension := hostParts[1]

	var videoID string

	if domain == "youtube" {
		queryValues, err := url.ParseQuery(parsedURL.RawQuery)
		if err != nil {
			return "", fmt.Errorf("parsing query values from url: %w", err)
		}

		if !queryValues.Has("v") {
			return "", fmt.Errorf(
				"url does not include video id query parameter: %s", rawURL,
			)
		}

		videoID = queryValues.Get("v")
	}

	if domain == "youtu" && extension == "be" {
		videoID = strings.TrimPrefix(parsedURL.Path, "/")
	}

	if videoID == "" {
		return "", fmt.Errorf("invalid youtube url: %s", rawURL)
	}

	if len(videoID) != videoIDLength {
		return "", fmt.Errorf("video id has incorrect length: %s", videoID)
	}

	return videoID, nil
}
