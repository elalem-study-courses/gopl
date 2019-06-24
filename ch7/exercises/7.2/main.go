// package description
package main

import (
	"fmt"
	"io"
	"strings"
)

var text = `This is the first line
This is the second line
This is the third line`

type WriterCount struct {
	Count  *int64
	Writer io.Writer
}

func (w *WriterCount) Write(p []byte) (int, error) {
	*(w.Count) += int64(len(p))
	return w.Writer.Write(p)
}

type WordCounter int
type LineCounter int

func (w WordCounter) Write(p []byte) (int, error) {
	words := strings.FieldsFunc(string(p), func(r rune) bool {
		if r == ' ' || r == '\n' {
			return true
		}
		return false
	})
	w += WordCounter(len(words))
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

func CountingtWriter(w io.Writer) (WriterCount, *int64) {
	wrappedWriter := WriterCount{Writer: w, Count: new(int64)}
	return wrappedWriter, wrappedWriter.Count
}

func main() {
	var w WordCounter
	// var l LineCounter

	wrappedWordCounter, byteCount := CountingtWriter(w)
	fmt.Fprintf(&wrappedWordCounter, text)
	fmt.Printf("Written %d bytes \n", *byteCount)
}
