package main

import (
	"fmt"
	"strings"
)

type WordCount int
type LineCount int

func (wc *WordCount) Write(p []byte) (int, error) {
	*wc += WordCount(len(strings.Split(string(p), " ")))
	return int(*wc), nil
}

func (lc *LineCount) Write(p []byte) (int, error) {
	*lc += LineCount(len(strings.Split(string(p), "\n")))
	return int(*lc), nil
}

func main() {
	var (
		wc WordCount
		lc LineCount
	)

	fmt.Fprintf(&wc, "Hello world")
	fmt.Fprintf(&lc, `This is 
	A multiline
	phrase`)

	fmt.Println(wc, lc)
}
