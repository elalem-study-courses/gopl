package mirrors

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"sync"
	"time"

	"golang.org/x/net/html"
)

var (
	connections int
)

func init() {
	connections = runtime.NumCPU()
}

func NewCrawler(baseSite string) *Crawler {
	crawler := &Crawler{
		baseSite:        baseSite,
		seenLinks:       make(map[string]bool),
		discoveredLinks: make(chan []string, 10),
		worklist:        make(chan []string, 10),
		limiter:         make(chan struct{}, connections),
		done:            make(chan struct{}),
	}

	go func() { crawler.worklist <- []string{crawler.baseSite} }()

	return crawler
}

type Crawler struct {
	baseSite        string
	seenLinks       map[string]bool
	discoveredLinks chan []string
	worklist        chan []string
	wg              sync.WaitGroup
	errorChan       chan error
	err             error
	limiter         chan struct{}
	done            chan struct{}
	latestFetch     time.Time
}

func (c *Crawler) crawl() {
	c.latestFetch = time.Now()

	go c.generateLinks()

	go func() {
		for {
			if time.Since(c.latestFetch) > 10*time.Second {
				close(c.discoveredLinks)
				close(c.worklist)
				c.done <- struct{}{}
			}
			time.Sleep(1 * time.Second)
		}
	}()

	<-c.done
}

func (c *Crawler) generateLinks() {
	for {
		links, ok := <-c.worklist
		if !ok {
			break
		}
		for _, link := range links {
			if !c.seenLinks[link] && c.validURL(link) {
				c.seenLinks[link] = true
				go c.discoverLink(link)
			}
		}
	}

}

func (c *Crawler) discoverLink(link string) {
	doc, err := c.fetchDocument(link)
	if err != nil {
		log.Print(err)
	} else {
		links := c.discoverLinksInDocument(doc)

		c.worklist <- links
	}
}

func (c *Crawler) fetchDocument(link string) (*html.Node, error) {
	absoluteURL := c.getAbsoluteURL(link)
	fmt.Printf("Fetching %s...\n", absoluteURL)
	c.limiter <- struct{}{}
	res, err := http.Get(absoluteURL)
	c.latestFetch = time.Now()
	<-c.limiter
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse html for url %s: %v", link, err)
	}

	return doc, nil
}

func (c *Crawler) discoverLinksInDocument(node *html.Node) []string {
	links := []string{}
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	for cur := node.FirstChild; cur != nil; cur = cur.NextSibling {
		links = append(links, c.discoverLinksInDocument(cur)...)
	}

	return links
}

func (c *Crawler) validURL(link string) bool {
	base, err := url.Parse(c.baseSite)
	if err != nil {
		return false
	}

	cur, err := url.Parse(c.getAbsoluteURL(link))
	if err != nil {
		return false
	}

	return base.Hostname() == cur.Hostname()
}

func (c *Crawler) getAbsoluteURL(link string) string {
	base, err := url.Parse(c.baseSite)
	if err != nil {
		return ""
	}

	cur, err := url.Parse(link)
	if err != nil {
		return ""
	}

	return base.ResolveReference(cur).String()
}
