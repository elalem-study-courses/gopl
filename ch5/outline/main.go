// package description
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
	}
	outline(nil, doc, 0)
}

func outline(stack []string, n *html.Node, depth int) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Printf("%v\n", stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c, depth+1)
	}
}
