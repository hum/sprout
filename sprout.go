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

type Subreddit struct {
	Name  string
	Posts []Post
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

func (r *Reddit) Get(subreddits []string, limit int) (result map[string]Subreddit, err error) {
	if r.UseAPI {
		postLimit := strconv.Itoa(limit - 1) // for some reason Reddit returns limit+1 results
		result, err = r.get(subreddits, postLimit)
		if err != nil {
			return result, fmt.Errorf("Could not execute get function: %v", err)
		}
		return
	}
	// TODO:
	// Let user use a webcrawler instead of only an API auth
	return result, fmt.Errorf("Non-API post harvest for Reddit is not implemented yet. Use the API.")
}

func createRedditBot(conf *Config) (reddit.Bot, error) {
	c, err := validateRedditConf(conf)
	if err != nil {
		return nil, fmt.Errorf("Could not validate Reddit Config: %v", err)
	}

	bot, err := reddit.NewBot(c)
	if err != nil {
		return nil, fmt.Errorf("Could not create new Reddit Bot instance: %v", err)
	}

	return bot, nil
}

func (r *Reddit) get(subreddits []string, limit string) (result map[string]Subreddit, err error) {
	if r.bot == nil {
		r.bot, err = createRedditBot(r.Conf)
		if err != nil {
			return result, fmt.Errorf("Could not create reddit bot: %v", err)
		}
	}

	result = make(map[string]Subreddit, len(subreddits))

	params := map[string]string{
		"limit": limit,
	}

	for _, subreddit := range subreddits {
		format := "/r/%s"

		sub := Subreddit{}
		sub.Name = subreddit

		harvest, err := r.bot.ListingWithParams(fmt.Sprintf(format, subreddit), params)
		if err != nil {
      fmt.Println("Error listing for ", subreddit)
			continue
		}

		for _, post := range harvest.Posts {
			// TODO:
			// Allow user to set params that filter out URLs
			// This ignores all posts that don't have an image
			if strings.Contains(post.URL, "/comments/") {
				continue
			}

			p := Post{Name: post.Title, Author: post.Author, Link: post.URL}
			sub.Posts = append(sub.Posts, p)
		}
		result[subreddit] = sub
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
