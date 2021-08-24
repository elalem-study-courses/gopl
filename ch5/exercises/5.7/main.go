package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth = 0

func main() {
	for _, url := range os.Args[1:] {
		body, err := fetch(url)
		if err != nil {
			log.Println(err)
		}
		doc, err := html.Parse(body)
		if err != nil {
			log.Fatal(err)
		}

		links := visit(nil, doc, startElement, endElement)
		fmt.Println(strings.Join(links, "\n"))
	}
}

func handleElementNode(n *html.Node) string {
	attrs := make([]string, 0)
	for _, attr := range n.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, attr.Key, attr.Val))
	}

	endTag := ">"
	if n.Data == "img" {
		endTag = "/>"
	}

	return fmt.Sprintf("%*s<%s %s %s", depth*2, " ", n.Data, strings.Join(attrs, " "), endTag)
}

func handleCommentNode(n *html.Node) string {
	return fmt.Sprintf("%*s<!-- %s -->", depth*2, " ", strings.TrimSpace(n.Data))
}

func handleTextNode(n *html.Node) string {
	return fmt.Sprintf("%*s%s", depth*2, " ", n.Data)
}

func startElement(n *html.Node) string {
	var output string
	switch n.Type {
	case html.ElementNode:
		output = handleElementNode(n)
	case html.CommentNode:
		output = handleCommentNode(n)
	case html.TextNode:
		output = handleTextNode(n)
	}

	depth++
	return output
}

func endElement(n *html.Node) string {
	depth--
	if n.Type != html.ElementNode || n.Data == "img" {
		return ""
	}

	return fmt.Sprintf("%*s</%s>", depth*2, " ", n.Data)
}

func fetch(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch: Received status %s instead of 200: %v", res.Status, err)
	}

	buf := &bytes.Buffer{}
	io.Copy(buf, res.Body)
	return buf, nil
}

func visit(links []string, n *html.Node, pre, post func(*html.Node) string) []string {
	if pre != nil {
		links = append(links, pre(n))
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c, pre, post)
	}

	if post != nil {
		out := post(n)
		if len(out) > 0 {
			links = append(links, post(n))
		}
	}

	return links
}
