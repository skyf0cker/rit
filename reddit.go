package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	_redditOAuthURL = "https://oauth.reddit.com"
	_redditURL      = "https://www.reddit.com"
)

var (
	_oauthHeaders = map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	_redditHeaders = map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
		"User-Agent":   "rit/0.1 by daydreaming_neo",
	}
)

type Reddit struct {
	oauthC *resty.Client
	c      *resty.Client

	cfg *Config

	token       *AccessToken
	tokenExpire time.Time
}

func NewReddit(cfg *Config) *Reddit {
	return &Reddit{
		oauthC: resty.New().
			SetBaseURL(_redditURL).
			SetHeaders(_oauthHeaders),
		c: resty.New().
			SetBaseURL(_redditOAuthURL).
			SetHeaders(_redditHeaders),
		cfg: cfg,
	}
}

type AccessToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

type Error struct {
	Error string `json:"error"`
}

func (r *Reddit) refreshToken() error {
	now := time.Now()
	token, err := r.getAccessToken()
	if err != nil {
		return err
	}

	r.token = token
	r.tokenExpire = now.Add(time.Duration(token.ExpiresIn) * time.Second)

	// default Bearer tokena type
	r.c = r.c.SetAuthToken(r.token.Token)
	return nil
}

func (r *Reddit) preCheck() error {
	if r.token == nil || time.Now().After(r.tokenExpire) {
		return r.refreshToken()
	}

	return nil
}

func (r *Reddit) getAccessToken() (*AccessToken, error) {
	var (
		token     AccessToken
		redditErr Error
	)

	resp, err := r.oauthC.R().
		SetBasicAuth(r.cfg.Credential.Id, r.cfg.Credential.Secret).
		SetFormData(map[string]string{
			"grant_type": "password",
			"username":   r.cfg.Credential.Username,
			"password":   r.cfg.Credential.Password,
		}).
		SetResult(&token).
		SetError(&redditErr).
		Post("/api/v1/access_token")
	if err != nil {
		return nil, fmt.Errorf("get reddit access token failed, error: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("get reddit access token failed, error: %s", redditErr.Error)
	}

	return &token, nil
}

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

func (r *Reddit) GetHomePage() ([]*Post, error) {
	if err := r.preCheck(); err != nil {
		return nil, fmt.Errorf("reddit request precheck failed, error: %w", err)
	}

	var data RedditItem[Listing]
	resp, err := r.c.R().SetResult(&data).Get("/.json")
	if err != nil {
		return nil, fmt.Errorf("request reddit homepage failed, error: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("request reddit homepage failed")
	}

	var posts []*Post
	for _, post := range data.Data.Children {
		posts = append(posts, &post.Data)
	}

	return posts, nil
}
