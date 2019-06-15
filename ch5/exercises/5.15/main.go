package main

import "fmt"

func max(initial int, rest ...int) int {
	maxElem := initial
	for _, elem := range rest {
		if maxElem < elem {
			maxElem = elem
		}
	}
	return maxElem
}

func min(initial int, rest ...int) int {
	minElem := initial
	for _, elem := range rest {
		if minElem > elem {
			minElem = elem
		}
	}
	return minElem
}

func main() {
	fmt.Println(max(1, 2, 3, 4, 5))
	fmt.Println(min(1, 2, 3, 4, 5))
}
