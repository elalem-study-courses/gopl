package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node interface {
	String() string
}

type CharData string

func (cd CharData) String() string {
	return string(cd)
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e *Element) String() string {
	attrs := make([]string, 0)
	for _, attr := range e.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, attr.Name.Local, attr.Value))
	}
	str := fmt.Sprintf("<%s %s>", e.Type.Local, strings.Join(attrs, " "))

	for _, child := range e.Children {
		str += child.String()
	}

	str += fmt.Sprintf("</%s>", e.Type.Local)

	return str
}

func main() {
	doc := xml.NewDecoder(os.Stdin)
	var stack = []Node{}

	for {
		tok, err := doc.Token()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlprint: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elem := &Element{
				Attr:     tok.Attr,
				Children: []Node{},
				Type:     tok.Name,
			}
			stack = append(stack, elem)
		case xml.EndElement:
			endStack := len(stack) - 1
			if endStack > 0 {
				parent := stack[endStack-1].(*Element)
				parent.Children = append(parent.Children, stack[endStack])
				stack = stack[:endStack]
			}
		case xml.CharData:
			if len(stack) > 1 {
				parent := stack[len(stack)-1].(*Element)
				parent.Children = append(parent.Children, CharData(tok))
			}
		}
	}
	node := stack[0].(*Element)
	fmt.Println(node)
}
