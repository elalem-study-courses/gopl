// package description
package main

import (
	"fmt"
	"strings"
)

var text = `This is the first line
This is the second line
This is the third line`

type WordCounter int
type LineCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	words := strings.FieldsFunc(string(p), func(r rune) bool {
		if r == ' ' || r == '\n' {
			return true
		}
		return false
	})
	*w += WordCounter(len(words))
	return len(words), nil
}

func (w *WordCounter) String() string {
	return fmt.Sprintf("Counted %v words", *w)
}

func (l *LineCounter) Write(p []byte) (int, error) {
	lines := strings.Split(string(p), "\n")
	*l += LineCounter(len(lines))
	return len(lines), nil
}

func (l *LineCounter) String() string {
	return fmt.Sprintf("Counted %d lines", *l)
}

func main() {
	var w WordCounter
	var l LineCounter

	fmt.Fprintf(&w, "%s", text)
	fmt.Println(&w)

	fmt.Fprintf(&l, "%s", text)
	fmt.Println(&l)
}
