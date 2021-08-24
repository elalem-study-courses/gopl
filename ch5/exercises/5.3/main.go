package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	for _, content := range pageContents(nil, doc) {
		fmt.Println(content)
		fmt.Println("----------------------------------------")
	}
}

func pageContents(contents []string, n *html.Node) []string {
	if n.Type == html.TextNode && n.Parent.Data != "style" && n.Parent.Data != "script" && n.Parent.Data != "noscript" {
		data := strings.TrimSpace(n.Data)
		if data != "" {
			contents = append(contents, data)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		contents = pageContents(contents, c)
	}

	return contents
}
