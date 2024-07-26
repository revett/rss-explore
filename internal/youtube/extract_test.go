package youtube_test

import (
	"testing"

	"github.com/revett/rss-explore/internal/youtube"
	"github.com/stretchr/testify/require"
)

func TestExtractVideoID(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		url  string
		want string
	}{
		"ValidLong": {
			url:  "https://youtube.com/watch?v=7LICrnxWd38",
			want: "7LICrnxWd38",
		},
		"ValidLongWithWWW": {
			url:  "https://www.youtube.com/watch?v=7LICrnxWd38",
			want: "7LICrnxWd38",
		},
		"ValidShort": {
			url:  "https://youtu.be/7LICrnxWd38",
			want: "7LICrnxWd38",
		},
		"ValidShortWithWWW": {
			url:  "https://www.youtu.be/7LICrnxWd38",
			want: "7LICrnxWd38",
		},
	}

	for n, testCase := range tests {
		tc := testCase

		t.Run(n, func(t *testing.T) {
			t.Parallel()

			videoID, err := youtube.ExtractVideoID(tc.url)

			require.NoError(t, err)
			require.Equal(t, tc.want, videoID)
		})
	}
}
