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

	for key, val := range findLinks(make(map[string][]string), doc) {
		fmt.Println(key)
		fmt.Println("---------")
		for _, link := range val {
			fmt.Println(link)
		}

		fmt.Println("------------------------------------------------")
	}
}

func findLinks(links map[string][]string, n *html.Node) map[string][]string {
	if n.Type == html.ElementNode && contains(n.Data, []string{"img", "a", "style", "script", "noscript"}) {
		for _, attr := range n.Attr {
			if attr.Key == "href" || attr.Key == "src" {
				links[n.Data] = append(links[attr.Key], attr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = findLinks(links, c)
	}

	return links
}

func contains(needle string, haystack []string) bool {
	for _, cand := range haystack {
		if cand == needle {
			return true
		}
	}

	return false
}
