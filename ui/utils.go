package ui

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func HTMLText(t string) string {
	root, err := html.Parse(strings.NewReader(html.UnescapeString(t)))
	if err != nil {
		panic(err)
	}

	var sb strings.Builder

	type state struct {
		element    string
		attributes map[string]string
	}

	var fn func(state, *html.Node)
	fn = func(s state, n *html.Node) {
		switch n.Type {
		case html.TextNode:
			text := n.Data
			switch s.element {
			case "a":
				if val, ok := s.attributes["href"]; ok {
					text = fmt.Sprintf("(%s %s)", n.Data, val)
					if n.Data == val {
						text = val
					} else if trim := strings.TrimSuffix(n.Data, "..."); strings.HasPrefix(val, trim) {
						// HN truncates long links and appends "..."
						text = val
					}
				}
			}

			sb.WriteString(text)
		case html.ElementNode:
			switch n.Data {
			case "a":
				s.element = n.Data
				s.attributes = make(map[string]string)
				for _, attr := range n.Attr {
					// discard Namespace
					s.attributes[attr.Key] = attr.Val
				}
			case "p":
				sb.WriteString("\n\n")
			}
		}

		for child := n.FirstChild; child != nil; child = child.NextSibling {
			fn(s, child)
		}
	}

	fn(state{}, root)
	return sb.String()
}
