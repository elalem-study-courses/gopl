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
var id = flag.String("id", "", `id of the element to find"`)

var usage = `Finds the first element by id`

func main() {
	flag.Usage = func() {
		fmt.Printf(usage)
	}
	flag.Parse()
	if *url != "" && *id != "" {
		doc, err := fetchDoc(*url)
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Printf("%v\n", find(doc, *id))
	}
}

func forEachNode(n *html.Node, id string, pre func(n *html.Node, id string) *html.Node, post func(n *html.Node)) (foundNode *html.Node) {
	if pre != nil {
		foundNode = pre(n, id)
		if foundNode != nil {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		foundNode = forEachNode(c, id, pre, post)
		if foundNode != nil {
			return
		}
	}

	if post != nil {
		post(n)
	}

	return
}

func startElement(n *html.Node, id string) *html.Node {
	if n.Type == html.ElementNode {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				return n
			}
		}
	}
	return nil
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

func find(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, nil)
}
