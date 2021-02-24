package sprout

import (
	"github.com/turnage/graw/reddit"
	"net/http"
	"time"
)

type Config struct {
	Username     string
	Password     string
	ClientID     string
	ClientSecret string
	UserAgent    string

	App    reddit.App
	Rate   time.Duration
	Client *http.Client
}
