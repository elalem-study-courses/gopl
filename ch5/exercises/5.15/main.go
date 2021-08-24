package main

import "fmt"

func max(initial int, rest ...int) int {
	m := initial
	for _, val := range rest {
		if val > m {
			m = val
		}
	}

	return m
}

func main() {
	fmt.Println(max(2, 1, 4, 3, 2, 1, 5, 6, 4))
}
