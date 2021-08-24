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

	m := mapElementsCount(doc, make(map[string]int))
	for key, val := range m {
		fmt.Printf("%s: Repeated %d times\n", key, val)
	}
}

func mapElementsCount(n *html.Node, elements map[string]int) map[string]int {
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		elements = mapElementsCount(c, elements)
	}

	return elements
}
