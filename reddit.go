package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	_redditOAuthURL = "https://oauth.reddit.com"
	_redditURL      = "https://www.reddit.com"
)

type Reddit struct {
	c *resty.Client

	cfg   *Config
	token *AccessToken
}

func NewReddit(cfg *Config) *Reddit {
	return &Reddit{
		c: resty.New().
			SetHeaders(map[string]string{
				"Content-Type": "application/json",
				"Accept":       "application/json",
			}),
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

func (r *Reddit) GetAccessToken() (*AccessToken, error) {
	var (
		token     AccessToken
		redditErr Error
	)

	resp, err := r.c.R().
		SetBasicAuth(r.cfg.Credential.Id, r.cfg.Credential.Secret).
		SetFormData(map[string]string{
			"grant_type": "password",
			"username":   r.cfg.Credential.Username,
			"password":   r.cfg.Credential.Password,
		}).
		SetResult(&token).
		SetError(&redditErr).
		Post(_redditURL + "/api/v1/access_token")
	if err != nil {
		return nil, fmt.Errorf("get reddit access token failed, error: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("get reddit access token failed, error: %s", redditErr.Error)
	}

	return &token, nil
}
