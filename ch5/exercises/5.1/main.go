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
		log.Fatal(err)
	}

	for _, link := range findLinks(nil, doc) {
		fmt.Println(link)
	}
}

func findLinks(links []string, n *html.Node) []string {
	if n != nil && n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	if n != nil {
		links = findLinks(links, n.FirstChild)
		links = findLinks(links, n.NextSibling)
	}

	return links
}
