package main

import (
	"fmt"
	"log"
)

const (
	Explored = iota + 1
	Visited
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	// "linear algebra":        {"calculus"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]int)
	var visitAll func(string)

	visitAll = func(url string) {
		if seen[url] == Explored {
			return
		}
		seen[url] = Visited
		for _, item := range m[url] {
			if item == "calculas" {
				fmt.Println(url, m[item], seen[item])
			}
			if _, ok := seen[item]; !ok {
				visitAll(item)
			} else if seen[item] == Visited {
				log.Fatalf("Cycle detected %s: %s\n", url, item)
			}
		}
		order = append(order, url)
		seen[url] = Explored
	}

	for key := range m {
		visitAll(key)
	}
	return order
}
