// package description
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

var url = flag.String("url", "", `Sets the value of the url to outline "ensure the protocol exists"`)

func main() {
	flag.Parse()
	if *url != "" {
		doc, err := fetchDoc(*url)
		if err != nil {
			log.Fatalf(err.Error())
		}
		outline(doc)
	}
}

func forEachNode(n *html.Node, depth int, pre, post func(n *html.Node, depth int)) {
	if pre != nil {
		pre(n, depth)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, depth+1, pre, post)
	}

	if post != nil {
		post(n, depth)
	}
}

func startElement(n *html.Node, depth int) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node, depth int) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func fetchDoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Couldn't parse HTML for %s: %s", url, err)
	}
	resp.Body.Close()
	return doc, err
}

func outline(doc *html.Node) {
	forEachNode(doc, 0, startElement, endElement)
}
