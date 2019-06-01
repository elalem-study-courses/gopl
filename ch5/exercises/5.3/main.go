// package description
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

const (
	IndentationPrefix = "  "
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Panicf("print-html: %v\n", err)
	}

	visit(doc, "")
}

func visit(n *html.Node, indendationPrefix string) {
	if n.Data == "style" || n.Data == "script" {
		return
	}

	if n.Type == html.ElementNode {
		fmt.Printf("%s<%s>\n", indendationPrefix, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, indendationPrefix+IndentationPrefix)
	}

	if n.Type == html.ElementNode {
		fmt.Printf("%s</%s>\n", indendationPrefix, n.Data)
	}
}
