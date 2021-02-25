# Sprout
A library to collect posts/images from social media. A very early WIP. 

### Example
```go
import "github.com/hum/sprout"

sprout := sprout.New()
reddit := sprout.Reddit()

# Create config
config := &sprout.Config{
  Username: "username",
  Password: "password",
  UserAgent: "Darwin:github.com/hum/sprout:0.1.0 (by /u/username)",
  ClientID: "client_id",
  ClientSecret: "client_secret",
}

# Set config
reddit.Conf = config
reddit.UseAPI = true

# Pick subreddits to get data from
subreddits := []string{
	"funny",
	"gaming",
	"aww",
}

limit := 10 # limit the amount of posts harvested
for _, subreddit := range subreddits {
	result, err := reddit.Get(subreddit, limit)
	if err != nil {
		panic(err)
	}
	
	for _, post := range result.Posts {
		fmt.Println(post.Link)
	}
}
```

### TODO:
  - [ ] Unified parsing and structuring
  - [ ] Web crawler
  - [ ] Sites
    - [ ] Reddit
    - [ ] Instagram
    - [ ] Twitter
    - [ ] Facebook(?)
  - [ ] Tests
  - [ ] Examples

