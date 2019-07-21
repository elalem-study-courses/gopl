package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"./links"
)

var depth = flag.Int("depth", 0, "The longest depth on path to take")

var tokens = make(chan struct{}, 20)

type URL struct {
	value string
	depth int
}

func crawl(url URL) []URL {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url.value)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	urls := make([]URL, len(list))
	for _, link := range list {
		urls = append(urls, URL{value: link, depth: url.depth + 1})
	}
	return urls
}

func main() {
	flag.Parse()

	worklist := make(chan []URL)
	unseenLinks := make(chan URL)

	go func() {
		urls := []URL{}
		for _, link := range os.Args[1:] {
			urls = append(urls, URL{value: link, depth: 0})
		}
		worklist <- urls
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.value] && link.depth <= *depth {
				seen[link.value] = true
				unseenLinks <- link
			}
		}
	}
}
