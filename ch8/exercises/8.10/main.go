package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"golang.org/x/net/html"
)

var tokens = make(chan struct{}, 20)
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func Extract(url string) ([]string, error) {
	if cancelled() {
		return nil, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Cancel = done
	res, err := http.DefaultClient.Do(req)
	// res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("got Status %d expected 200", res.StatusCode)
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
	if cancelled() {
		return nil
	}

	fmt.Println(url)
	tokens <- struct{}{}
	list, err := Extract(url)
	<-tokens
	if err != nil {
		log.Println(err)
	}

	return list
}

func main() {
	worklist := make(chan []string)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)

	n := 1

	var wg sync.WaitGroup

loop:
	for ; n > 0; n-- {
		select {
		case list := <-worklist:
			for _, link := range list {
				if cancelled() {
					break
				}
				if !seen[link] {
					wg.Add(1)
					seen[link] = true
					n++
					go func(link string) {
						select {
						case worklist <- crawl(link):
						case <-done:
						}
						wg.Done()
					}(link)
				}
			}
		case <-done:
			break loop
		}
	}

	wg.Wait()
	close(worklist)
	for range worklist {
	}
}
