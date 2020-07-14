package medium

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/html"

	"github.com/ByteSchneiderei/medium-rss-api/pkg/config"
	"github.com/mmcdole/gofeed"
	"github.com/patrickmn/go-cache"
)

// Medium represents blog class
type Medium struct {
	EnvVar *config.EnvVar
	Cache  *cache.Cache
}

// New instantiates new Blog
func New(e *config.EnvVar, c *cache.Cache) *Medium {
	return &Medium{EnvVar: e, Cache: c}
}

// Fetch fetches data from rss feed
func (m *Medium) Fetch() (*Response, error) {
	var response Response
	// check whether response is already stored in cache. If so, return from cache
	cachedResponse, found := m.Cache.Get(config.CacheKey)
	response, ok := cachedResponse.(Response)
	if found && ok {
		return &response, nil
	}

	fp := gofeed.NewParser()
	source := fmt.Sprintf("%s/%s", m.EnvVar.MediumRSSFeedURL, m.EnvVar.MediumProfile)
	feed, err := fp.ParseURL(source)
	if err != nil {
		return nil, err
	}

	// replace ?source=rss----c4c00d9be425---4 with empty string
	sourceRSSRegexPattern := regexp.MustCompile("\\?source=rss.*")

	posts := make([]Post, 0)
	for _, item := range feed.Items {
		guid := strings.Split(item.GUID, "/")
		content, err := tokenizeHTML(item.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, Post{
			ID:         guid[len(guid)-1:][0],
			Title:      item.Title,
			Link:       sourceRSSRegexPattern.ReplaceAllString(item.Link, ""),
			Published:  item.PublishedParsed.String(),
			Categories: item.Categories,
			Author:     item.Author.Name,
			Content:    content,
		})
	}

	response = Response{
		Title:       feed.Title,
		Description: feed.Description,
		Link:        sourceRSSRegexPattern.ReplaceAllString(feed.Link, ""),
		Posts:       posts,
	}

	m.Cache.Set(config.CacheKey, response, config.CacheDefaultExpiration)

	return &response, nil
}

func tokenizeHTML(s string) (*Node, error) {
	htmlContent := strings.NewReader(s)
	doc, err := html.Parse(htmlContent)
	if err != nil {
		return nil, err
	}

	children := make([]*Node, 0)
	body := &Node{
		Tag:      "body",
		Text:     "",
		Attrs:    nil,
		Children: children,
	}
	excludedTags := map[string]bool{
		"html": true, "head": true, "body": true,
	}

	var f func(*html.Node, *[]*Node) []*Node
	f = func(n *html.Node, nodes *[]*Node) []*Node {
		childNode := &Node{}
		childNode.Children = make([]*Node, 0)

		continueTokenization := true
		if n.Type == html.DocumentNode {
			continueTokenization = false
		}

		if n.Type == html.ElementNode {
			if found := excludedTags[n.Data]; found {
				continueTokenization = false
			}
		}

		if continueTokenization {
			if n.Type == html.ElementNode {
				attrs := make([]map[string]string, 0)
				for _, a := range n.Attr {
					attrs = append(attrs, map[string]string{"key": a.Key, "value": a.Val})
				}
				childNode.Tag = n.Data
				childNode.Attrs = attrs
			} else if n.Type == html.TextNode {
				childNode.Text = n.Data
			}
		}

		for child := n.FirstChild; child != nil; child = child.NextSibling {
			childNode.Children = f(child, &childNode.Children)
		}

		if continueTokenization {
			*nodes = append(*nodes, childNode)
		} else {
			*nodes = childNode.Children
		}

		return *nodes
	}

	body.Children = f(doc, &body.Children)

	return body, nil
}
