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
	subreddits, err := reddit.Get(subs, limit)
	if err != nil {
		t.Fatalf("%v\n", err)
	}

	if len(subreddits["memes"].Posts) > limit {
		t.Fatalf("The length of harvested posts does not match the set limit. For %s subreddit: got=%d, expected=%d\n", subreddits["memes"].Name, len(subreddits["memes"].Posts), limit)
	}
}
