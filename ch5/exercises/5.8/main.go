package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		body, err := fetch(url)
		if err != nil {
			log.Println(err)
			continue
		}

		doc, err := html.Parse(body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		nodes := findNodes(nil, doc, findElementByClassName("j1ei8c"))
		fmt.Printf("Found %d elements\n", len(nodes))
	}
}

func findElementByClassName(className string) func(*html.Node) *html.Node {
	return func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == className {
					return n
				}
			}
		}

		return nil
	}
}

func fetch(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch: Couldn't fetch %s got status %s: %v", url, res.Status, err)
	}

	buf := &bytes.Buffer{}
	io.Copy(buf, res.Body)

	return buf, nil
}

func findNodes(nodes []*html.Node, n *html.Node, filter func(*html.Node) *html.Node) []*html.Node {
	if filter != nil {
		if node := filter(n); n != nil {
			nodes = append(nodes, node)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = findNodes(nodes, c, filter)
	}

	return nodes
}
