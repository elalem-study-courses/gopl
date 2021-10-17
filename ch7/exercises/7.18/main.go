package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"
)

type Node interface{}

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func main() {
	doc := xml.NewDecoder(os.Stdin)
	stack := make([]Node, 0)
	for {
		tok, err := doc.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			elem := Element{Type: tok.Name}
			if len(stack) > 0 {
				lastElement := stack[len(stack)-1].(Element)
				lastElement.Children = append(lastElement.Children, elem)
			}
			stack = append(stack, elem)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(stack) > 0 {
				lastElement := stack[len(stack)-1].(Element)
				lastElement.Children = append(lastElement.Children, tok)
			}
		case xml.Attr:
			lastElement := stack[len(stack)-1].(Element)
			lastElement.Attr = append(lastElement.Attr, tok)
		}
	}
}
