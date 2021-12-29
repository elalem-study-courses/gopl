package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type Endpoint struct {
	url   string
	depth int
}

var tokens = make(chan struct{}, 20)

func Extract(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("got Status %d expected 200")
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s failed: %v", url, err)
	}

	links := make([]string, 0)

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}

				link, err := res.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}

				links = append(links, link.String())
			}
		}
	}

	forEachNode(doc, visitNode, nil)

	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := Extract(url)
	<-tokens
	if err != nil {
		log.Println(err)
	}

	return list
}

var depth = flag.Int("depth", 1, "Specify the maximum depth for the crawler")

func main() {
	flag.Parse()

	worklist := make(chan []Endpoint)

	go func() {
		endpoints := []Endpoint{}
		for _, url := range os.Args[1:] {
			endpoints = append(endpoints, Endpoint{url: url, depth: 0})
		}
		worklist <- endpoints
	}()

	seen := make(map[string]bool)

	n := 1

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link.url] && link.depth < *depth {
				seen[link.url] = true
				n++
				go func(link Endpoint) {
					links := crawl(link.url)
					endpoints := make([]Endpoint, 0)
					for _, url := range links {
						endpoints = append(endpoints, Endpoint{url: url, depth: link.depth + 1})
					}
					worklist <- endpoints
				}(link)
			}
		}
	}
}
