// package description
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"./charcount"
)

func main() {
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

		charcount.CharCount(r, n)
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range charcount.Counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("\nlen\tcount\n")

	for i, n := range charcount.UTFLen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if charcount.Invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", charcount.Invalid)
	}
}
