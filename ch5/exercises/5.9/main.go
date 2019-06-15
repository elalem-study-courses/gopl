package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		if word[0] == '$' {
			words[i] = f(word)
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(expand("foo $foo", func(s string) string {
		return s[1:]
	}))
}
