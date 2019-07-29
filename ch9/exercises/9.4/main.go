package main

import "fmt"

var depth int

/**
Max depth reached

Depth = 4194291

Stack overflow fatal error
*/

func main() {
	done := make(chan struct{})

	var rec func(done chan struct{})
	rec = func(done chan struct{}) {
		depth++
		fmt.Printf("\rDepth = %d", depth)
		anotherDone := make(chan struct{})
		rec(anotherDone)
		done <- struct{}{}
	}

	go rec(done)

	<-done
}
