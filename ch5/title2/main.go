package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func title(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return "", fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("oarsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			return n.FirstChild.Data
		}

		return ""
	}

	return forEachNode(doc, visitNode), nil
}

func forEachNode(n *html.Node, callback func(n *html.Node) string) string {
	ret := ""
	if callback != nil {
		ret = callback(n)
	}

	if len(ret) == 0 {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			ret = forEachNode(c, callback)
			if len(ret) > 0 {
				break
			}
		}
	}

	return ret
}

func main() {
	for _, url := range os.Args[1:] {
		fmt.Println(title(url))
	}
}
