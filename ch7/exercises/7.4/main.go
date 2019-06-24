package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

type MyReader struct {
	data []byte
	i    int
}

var txt = "Enter a url to parse links from it"

func (m *MyReader) Read(b []byte) (n int, err error) {
	// scanner := bufio.NewScanner(os.Stdin)
	// if len(b) == 0 {
	// 	fmt.Println(txt)
	// }

	// if read := scanner.Scan(); !read {
	// 	err = io.EOF
	// 	return
	// }
	// readBytes := scanner.Bytes()
	// fmt.Println(readBytes)
	// n = copy(b, readBytes)
	// fmt.Printf("%p\n", &b)

	if m.i >= len(m.data) {
		err = io.EOF
		return
	}

	n = copy(b, m.data[m.i:])
	m.i += n
	return
}

func NewReader(in string) io.Reader {
	return &MyReader{data: []byte(in)}
}

func FetchDocument(link string) (*html.Node, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response returned with status %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not parse html document for %s: %v", link, err)
	}

	return doc, nil
}

func main() {
	link, _ := ioutil.ReadAll(NewReader("http://www.google.com"))
	doc, err := FetchDocument(string(link))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
