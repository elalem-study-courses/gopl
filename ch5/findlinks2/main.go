package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Printf("Visiting url %s...\n", arg)
		links, err := findLinks(arg)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(strings.Join(links, "\n"))
		}

		fmt.Println("------------------------------------------------")
	}
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting: %s: %v", url, err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()

	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}
