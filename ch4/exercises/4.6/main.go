package main

import (
	"fmt"
	"unicode"
)

func squash(bytes []byte) []byte {
	runes := []rune(string(bytes))
	squashedIndex := 1
	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] && unicode.IsSpace(runes[i]) {
			continue
		}
		runes[squashedIndex] = runes[i]
		squashedIndex++
	}
	return []byte(string(runes[:squashedIndex]))
}

func main() {
	s := "This   is     Sparta"
	fmt.Printf("%q\n", squash([]byte(s)))
}
