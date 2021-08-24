package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	tokens := strings.Split(s, " ")

	transformedTokens := make([]string, 0, len(tokens))
	for _, token := range tokens {
		transformedTokens = append(transformedTokens, f(token))
	}

	return strings.Join(transformedTokens, " ")
}

func main() {
	fmt.Println(expand("This is an example of $foo", func(token string) string {
		if token != "$foo" {
			return token
		}

		return "expand"
	}))
}
