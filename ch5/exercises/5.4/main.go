// package description
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Panicf("findlinksex: %v", err)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		link := ""
		switch n.Data {
		case "a":
			fallthrough
		case "style":
			link = extractAttributeValue(n, "href")
		case "img":
			fallthrough
		case "script":
			link = extractAttributeValue(n, "src")
		}
		if link != "" {
			links = append(links, link)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func extractAttributeValue(n *html.Node, targetAttr string) string {
	val := ""
	for _, attr := range n.Attr {
		if attr.Key == targetAttr {
			val = attr.Val
			break
		}
	}
	return val
}
