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

	imgAndLinks := getElementsByTagNames(nil, doc, "img", "a")
	fmt.Println(imgAndLinks)
}

func Contains(needle string, haystack []string) bool {
	for _, token := range haystack {
		if needle == token {
			return true
		}
	}

	return false
}

func getElementsByTagNames(elements []string, n *html.Node, tags ...string) []string {
	if n.Type == html.ElementNode && Contains(n.Data, tags) {
		elements = append(elements, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elements = getElementsByTagNames(elements, c, tags...)
	}

	return elements
}
