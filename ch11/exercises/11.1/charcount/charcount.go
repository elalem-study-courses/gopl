package charcount

import (
	"unicode"
	"unicode/utf8"
)

var (
	Counts  map[rune]int
	UTFLen  [utf8.UTFMax + 1]int
	Invalid int
)

func init() {
	Counts = make(map[rune]int)
}

func CharCount(r rune, n int) {
	if r == unicode.ReplacementChar && n == 1 {
		Invalid++
	} else {
		Counts[r]++
		UTFLen[n]++
	}
}

func Reset() {
	Counts = make(map[rune]int)
	for i := range UTFLen {
		UTFLen[i] = 0
	}
}
