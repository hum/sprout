package sprout

import (
  "fmt"
  "strconv"

  "github.com/turnage/graw/reddit"
)

type Reddit struct {
	bot      reddit.Bot
	Conf     *Config
	Username string
	Password string
	UseAPI   bool
}

type Subreddit struct {
	Name  string
	Posts []Post
}

func (r *Reddit) Get(subreddit string, limit int) (*Subreddit, error) {
	if r.UseAPI {
		postLimit := strconv.Itoa(limit - 1)
		result, err := r.get(subreddit, postLimit)
		if err != nil {
			return nil, err
		}
    return result, nil
	}
	// TODO:
	// Implement a web crawler solution
	return nil, fmt.Errorf("Non-API post harvest for Reddit is not implemented yet. Use the API\n.")
}

func (r *Reddit) get(subreddit, limit string) (*Subreddit, error) {
  var err error

	if r.bot == nil {
		r.bot, err = createRedditBot(r.Conf)
		if err != nil {
			return nil, err
		}
	}

	params := map[string]string{
		"limit": limit,
		// TODO:
		// Add more customizable params
	}

	s := &Subreddit{}
	s.Name = subreddit

  format := "/r/%s"
	harvest, err := r.bot.ListingWithParams(fmt.Sprintf(format, s.Name), params)
	if err != nil {
		return nil, fmt.Errorf("Could not harvest posts from [%s] subreddit. %v\n", s.Name, err)
	}

	for _, post := range harvest.Posts {
		p := Post{Name: post.Title, Author: post.Author, Link: post.URL}
		s.Posts = append(s.Posts, p)
	}
	return s, nil
}

func createRedditBot(config *Config) (reddit.Bot, error) {
	if !isValidRedditConfig(config) {
		return nil, fmt.Errorf("Reddit config is invalid.")
	}

	c := reddit.BotConfig{
		Agent: config.UserAgent,
		App: reddit.App{
			ID:       config.ClientID,
			Secret:   config.ClientSecret,
			Username: config.Username,
			Password: config.Password,
		},
	}

	bot, err := reddit.NewBot(c)
	if err != nil {
		return nil, fmt.Errorf("Could not create Reddit Bot instance. %v\n", err)
	}
	return bot, nil
}

func isValidRedditConfig(c *Config) bool {
	return c.UserAgent != "" && c.ClientID != "" && c.ClientSecret != "" && c.Username != "" && c.Password != ""
}
