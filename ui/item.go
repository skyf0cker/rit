package ui

import (
	"fmt"

	"github.com/skyf0cker/rit/reddit"
)

type PostItem struct {
	post *reddit.Post
}

func (p *PostItem) FilterValue() string {
	return p.post.Title
}

func (s PostItem) Title() string {
	return s.post.Title
	// var sb strings.Builder
	// fmt.Fprintf(&sb, "%d. %s", s.Rank+1, s.Item.Title)
	//
	// if s.URL != "" {
	// 	link, err := url.Parse(s.URL)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	//
	// 	fmt.Fprintf(&sb, " (%s)", link.Host)
	// }
	//
	// return sb.String()
}

func (s PostItem) Description() string {
	return fmt.Sprintf("  r/%s   %s 󰔓  %d 󰔑  %d",
		s.post.Subreddit,
		s.post.Author,
		s.post.Ups,
		s.post.Downs,
	)
}
