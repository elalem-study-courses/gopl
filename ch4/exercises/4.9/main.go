// package description
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := map[string]int{}

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		r := in.Text()
		words[r]++
	}

	for word, count := range words {
		fmt.Printf("%q\t%d\n", word, count)
	}
}
