package main

import "fmt"

func removeDups(s []string) []string {
	i, d := 0, 1
	for d < len(s) {
		if s[d] != s[d-1] {
			i++
			s[i] = s[d]
		}

		d++
	}

	return s[:i+1]
}

func main() {
	fmt.Println(removeDups([]string{"one", "one", "two", "two", "two", "three", "three"}))
}
