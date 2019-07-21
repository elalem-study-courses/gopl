// package description
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", fibN)
}

func spinner(delay time.Duration) {
	span := 0
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%*c", span, r)
			time.Sleep(delay)
		}
		span++
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
