// package descriptin
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
		log.Panicf("count-elements: %v", err)
	}

	mappedElements := visit(make(map[string]int), doc)
	for key, value := range mappedElements {
		fmt.Printf("%s appeared %v times\n", key, value)
	}
}

func visit(mappedElements map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		mappedElements[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		mappedElements = visit(mappedElements, c)
	}
	return mappedElements
}
