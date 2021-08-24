package main

import (
	"fmt"
	"log"
)

type NodeState int8

const (
	Visited = iota + 1
	Explored
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
	"linear algebra":        {"calculus"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string

	seen := make(map[string]NodeState)

	var visitAll func(string)

	visitAll = func(cur string) {
		seen[cur] = Visited
		for _, prereq := range prereqs[cur] {
			if _, ok := seen[prereq]; !ok {
				visitAll(prereq)
			}

			if seen[prereq] == Visited {
				log.Fatalf("Cycle detected between %s and %s\n", prereq, cur)
			}
		}

		seen[cur] = Explored

		order = append(order, cur)
	}

	for key := range m {
		if _, ok := seen[key]; !ok {
			visitAll(key)
		}
	}

	for i := 0; i < len(order)-i; i++ {
		order[i], order[len(order)-i-1] = order[len(order)-i-1], order[i]
	}

	return order
}
