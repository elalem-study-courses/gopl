package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type HTMLReader struct {
	str string
}

func (hr *HTMLReader) Read(p []byte) (int, error) {
	copy(p, []byte(hr.str))
	return len(hr.str), io.EOF
}

func main() {
	reader := &HTMLReader{str: "<html><body><p>Hello world</p></body></html>"}
	doc, err := html.Parse(reader)
	fmt.Println(doc.FirstChild.FirstChild.NextSibling.FirstChild.FirstChild.Data, err)
}
