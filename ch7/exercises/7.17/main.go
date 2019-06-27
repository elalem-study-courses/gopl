package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Equaler interface {
	Equal(Equaler) (bool, error)
}

type Node string

func (n Node) Equal(x Equaler) (bool, error) {
	x, ok := x.(Node)
	if !ok {
		return false, fmt.Errorf("Invalid argument type %T", x)
	}
	return x == n, nil
}

type Attribute string

func (a Attribute) Equal(x Equaler) (bool, error) {
	x, ok := x.(Attribute)
	if !ok {
		return false, fmt.Errorf("Invalid argument type %T", x)
	}
	return a == x, nil
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack = []xml.StartElement{}

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				// fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containsAll(x []xml.StartElement, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		nodeEql, err := Node(x[0].Name.Local).Equal(Node(y[0]))
		if err != nil {
			return false
		}

		attrEql := false
		for _, attr := range x[0].Attr {
			if eql, err := Attribute(attr.Value).Equal(Attribute(y[0])); err != nil {
				attrEql = false
				break
			} else if eql {
				attrEql = true
				break
			}
		}

		if nodeEql || attrEql {
			y = y[1:]
		}

		x = x[1:]
	}
	return false
}
