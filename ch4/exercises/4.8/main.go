package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freq := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		freq[input.Text()]++
	}

	fmt.Printf("Word\t\tFrequency\n")
	for word, f := range freq {
		fmt.Printf("%q\t\t%d\n", word, f)
	}
}
