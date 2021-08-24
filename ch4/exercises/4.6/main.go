package main

import (
	"fmt"
	"unicode"
)

func squashSpaces(s string) string {
	r := []rune(s)
	j := 0
	spaceCnt := 0
	for i := 0; i < len(r); i++ {
		if !unicode.IsSpace(r[i]) {
			spaceCnt = 0
		} else {
			spaceCnt++
		}

		if spaceCnt <= 1 {
			r[j] = r[i]
			j++
		}
	}

	return string(r[:j])
}

func main() {
	fmt.Printf("%q\n", squashSpaces("This  is a     test"))
	fmt.Printf("%q\n", squashSpaces("ديه    تجربة   بسيطة"))
}
