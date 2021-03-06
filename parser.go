package linkparser

import (
	"io"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseLinks(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	links := make([]Link, 0)
	var f func(*html.Node)

	f = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {
			hrefString := parseHref(n)
			text := parseText(n, "")
			link := Link{hrefString, text}
			links = append(links, link)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return links, nil
}

func parseHref(n *html.Node) string {
	var v string
	for _, att := range n.Attr {
		if att.Key == "href" {
			v = att.Val
			break
		}
	}
	return v
}

func parseText(n *html.Node, s string) string {

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			s += c.Data
		}
		s = parseText(c, s)
	}
	return s
}
