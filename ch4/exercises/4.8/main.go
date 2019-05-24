// package description
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	Letter = iota
	Digit
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	var cat [2]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			cat[Letter]++
		} else {
			cat[Digit]++
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\ncat\tcount\n")
	fmt.Printf("letters\t%d\n", cat[Letter])
	fmt.Printf("digits\t%d\n", cat[Digit])

	fmt.Printf("\nlen\tcount\n")

	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
