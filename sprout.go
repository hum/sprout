package sprout

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/turnage/graw/reddit"
)

type Sprout struct {
	reddit *Reddit
}

type Reddit struct {
	bot      reddit.Bot
	Conf     *Config
	Username string
	Password string
	UseAPI   bool
}

type Post struct {
	Name   string
	Author string
	Link   string
}

func New() *Sprout {
	return &Sprout{}
}

func (s *Sprout) Reddit() *Reddit {
	if s.reddit == nil {
		return &Reddit{}
	}
	return s.reddit
}

func (r *Reddit) Get(subreddits []string, limit int) (result []Post, err error) {
	if r.UseAPI {
		postLimit := strconv.Itoa(limit)
		result, err = r.get(subreddits, postLimit)
		if err != nil {
			return
		}
	}
	// TODO:
	// Let user use a webcrawler instead of only an API auth
	return result, fmt.Errorf("Non-API post harvest for Reddit is not implemented yet. Use the API.")
}

func createRedditBot(conf *Config) (reddit.Bot, error) {
	c, err := validateRedditConf(conf)
	if err != nil {
		return nil, err
	}

	bot, err := reddit.NewBot(c)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

func (r *Reddit) get(subreddits []string, limit string) (result []Post, err error) {
	if r.bot == nil {
		r.bot, err = createRedditBot(r.Conf)
		if err != nil {
			return
		}
	}

	params := map[string]string{
		"limit": limit,
	}

	for _, subreddit := range subreddits {
		format := "/r/%s"

		harvest, err := r.bot.ListingWithParams(fmt.Sprintf(format, subreddit), params)
		if err != nil {
			return result, err
		}

		for _, post := range harvest.Posts {
			if strings.Contains(post.URL, "/comments/") {
				continue
			}

			p := Post{Name: post.Title, Author: post.Author, Link: post.URL}
			result = append(result, p)
		}
	}
	return
}

func validateRedditConf(conf *Config) (reddit.BotConfig, error) {
	confError := "No %s string specified in config."

	if conf.UserAgent == "" {
		return reddit.BotConfig{}, fmt.Errorf(fmt.Sprintf(confError, "user-agent"))
	}

	if conf.ClientID == "" {
		return reddit.BotConfig{}, fmt.Errorf(fmt.Sprintf(confError, "client ID"))
	}

	if conf.ClientSecret == "" {
		return reddit.BotConfig{}, fmt.Errorf(fmt.Sprintf(confError, "client secret"))
	}

	if conf.Username == "" {
		return reddit.BotConfig{}, fmt.Errorf(fmt.Sprintf(confError, "username"))
	}

	if conf.Password == "" {
		return reddit.BotConfig{}, fmt.Errorf(fmt.Sprintf(confError, "password"))
	}

	return reddit.BotConfig{
		Agent: conf.UserAgent,
		App: reddit.App{
			ID:       conf.ClientID,
			Secret:   conf.ClientSecret,
			Username: conf.Username,
			Password: conf.Password,
		},
	}, nil
}
