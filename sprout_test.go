package sprout_test

import (
	"github.com/hum/sprout"
	"os"
	"testing"
)

func TestReddit(t *testing.T) {
	cfg := &sprout.Config{
		Username:     os.Getenv("REDDIT_USERNAME"),
		Password:     os.Getenv("REDDIT_PASSWORD"),
		UserAgent:    os.Getenv("REDDIT_USER_AGENT"),
		ClientID:     os.Getenv("REDDIT_CLIENT_ID"),
		ClientSecret: os.Getenv("REDDIT_CLIENT_SECRET"),
	}

	sprout := sprout.New()
	reddit := sprout.Reddit()

	reddit.UseAPI = true
	reddit.Conf = cfg

	subs := []string{"memes", "gaming", "dankmemes"}

	limit := 10

	for _, s := range subs {
		result, err := reddit.Get(s, limit)
		if err != nil {
			t.Fatalf("%v\n", err)
		}

		size := len(result.Posts)
		if size > limit {
			t.Fatalf("Fetched more images than the limit amount. Expected=%d, got=%d", limit, size)
		}
	}
}
