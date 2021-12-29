package mirrors

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	urllib "net/url"

	"golang.org/x/net/html"
)

var (
	// Limit number of concurrent requests
	keys = make(chan bool, 20)
)

type Links struct {
	pageContent string
	urls        []*urllib.URL
	src         *urllib.URL
}

func extractLinks(url string) (*Links, error) {
	keys <- true
	res, err := http.Get(url)
	<-keys

	srcURL, _ := urllib.Parse(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("expected Response 200 code for %q got %d", url, res.StatusCode)
	}

	content, _ := ioutil.ReadAll(res.Body)
	copiedPageContent := make([]byte, len(content))
	copy(copiedPageContent, content)
	buf := bytes.NewBuffer(content)
	doc, err := html.Parse(buf)

	if err != nil {
		return nil, fmt.Errorf("couldn't parse html page for url %q: %v", url, err)
	}

	links := make([]*urllib.URL, 0)
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					url, err := res.Request.URL.Parse(attr.Val)
					if err != nil {
						log.Printf("Couldn't parse url %q", attr.Val)
						continue
					}
					links = append(links, url)
				}
			}
		}
	}

	processPage(doc, visitNode)

	return &Links{pageContent: string(copiedPageContent), urls: links, src: srcURL}, nil
}

func processPage(n *html.Node, nodeProcessor func(n *html.Node)) {
	nodeProcessor(n)

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		processPage(c, nodeProcessor)
	}
}
