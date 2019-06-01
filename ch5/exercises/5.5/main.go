// package description
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := countWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "countwordsandimages: %v\n", err)
			continue
		}
		fmt.Printf("%s has %d words, %d images\n", url, words, images)
	}
}

func countWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("Parsing HMTL: %s", err)
		return
	}
	words, images = countWordsAndImagesInDoc(doc)
	return
}

func countWordsAndImagesInDoc(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "img":
			images++
		case "p":
			fallthrough
		case "h1":
			fallthrough
		case "h2":
			fallthrough
		case "h3":
			fallthrough
		case "h4":
			fallthrough
		case "h5":
			fallthrough
		case "h6":
			fallthrough
		case "a":
			b := bytes.Buffer{}
			html.Render(&b, n)
			words += len(strings.Split(b.String(), " "))
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		wordsCnt, imagesCnt := countWordsAndImagesInDoc(c)
		words += wordsCnt
		images += imagesCnt
	}

	return
}
