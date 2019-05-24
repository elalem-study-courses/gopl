// Package description
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileMapper := make(map[string]string)

	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, "stdin", counts, fileMapper)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, arg, counts, fileMapper)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, fileMapper[line])
		}
	}
}

func countLines(f *os.File, filename string, counts map[string]int, filemapper map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text]++
		filemapper[text] = filename
	}
}
