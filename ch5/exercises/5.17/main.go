package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func getHTMLDocument(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse html for %s: %v", url, err)
	}

	return doc, nil
}

func ElementByTagName(n *html.Node, tags ...string) []*html.Node {
	nodes := make([]*html.Node, 0)
	if n.Type == html.ElementNode && contains(tags, n.Data) {
		nodes = append(nodes, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, ElementByTagName(c, tags...)...)
	}

	return nodes
}

func contains(arr []string, target string) bool {
	for _, elem := range arr {
		if elem == target {
			return true
		}
	}
	return false
}

var (
	url  = flag.String("url", "", "The html document to be parsed")
	tags = flag.String("tags", "", "tags to be looked up 'comma seperated'")
)

func main() {
	flag.Parse()

	doc, err := getHTMLDocument(*url)
	if err != nil {
		log.Fatal(err)
	}

	nodes := ElementByTagName(doc, strings.Split(*tags, ",")...)
	for _, node := range nodes {
		var b bytes.Buffer
		html.Render(&b, node)
		fmt.Println(b.String())
	}
}
