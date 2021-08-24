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
		log.Fatalf("outline: %v\n", err)
	}

	depthOperations := func() []func() int {
		depth := 0
		return []func() int{
			func() int {
				return depth
			},
			func() int {
				depth++
				return depth
			},
			func() int {
				depth--
				return depth
			},
		}
	}

	forEachNode(doc, startElement, endElement, depthOperations())
}

func startElement(n *html.Node, depthOperations []func() int) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depthOperations[0]()*2, "", n.Data)
		depthOperations[1]()
	}
}

func endElement(n *html.Node, depthOperations []func() int) {
	if n.Type == html.ElementNode {
		depthOperations[2]()
		fmt.Printf("%*s</%s>\n", depthOperations[0]()*2, "", n.Data)
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node, depthOperations []func() int), depthOperations []func() int) {
	if pre != nil {
		pre(n, depthOperations)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post, depthOperations)
	}

	if post != nil {
		post(n, depthOperations)
	}
}
