# Sprout
A library to collect posts/images from social media. A very early WIP. 

### Example
```go
import "github.com/hum/sprout"

sprout := sprout.New(UseAPI: false)
reddit := sprout.Reddit()

# Create config
config := &sprout.Config{
  Username: "username",
  Password: "password",
  UserAgent: "Darwin:github.com/hum/sprout:0.1.0 (by /u/username)",
  ClientID: "client_id",
  ClientSecret: "client_secret",
}

# Pick subreddits to get data from
subreddits := []string{
	"funny",
	"gaming",
	"aww",
}

posts, err := reddit.Get(subs)
if err != nil {
  panic(err)
}

for _, post := range posts {
  fmt.Println(post.Link)
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

