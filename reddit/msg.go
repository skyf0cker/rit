package reddit

type RedditItem[T any] struct {
	Kind string `json:"kind"`
	Data T      `json:"data"`
}

type Listing struct {
	Children []RedditItem[Post] `json:"children"`
}

type Post struct {
	Subreddit            string      `json:"subreddit"`
	Selftext             string      `json:"selftext"`
	AuthorFullname       string      `json:"author_fullname"`
	Clicked              bool        `json:"clicked"`
	Title                string      `json:"title"`
	Downs                int         `json:"downs"`
	Ups                  int         `json:"ups"`
	Category             interface{} `json:"category"`
	Created              float64     `json:"created"`
	SelftextHTML         string      `json:"selftext_html"`
	Likes                interface{} `json:"likes"`
	Over18               bool        `json:"over_18"`
	Visited              bool        `json:"visited"`
	RemovedBy            interface{} `json:"removed_by"`
	ID                   string      `json:"id"`
	Author               string      `json:"author"`
	Permalink            string      `json:"permalink"`
	URL                  string      `json:"url"`
	SubredditSubscribers int         `json:"subreddit_subscribers"`
	CreatedUtc           float64     `json:"created_utc"`
}
