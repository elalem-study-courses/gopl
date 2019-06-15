package main

import "fmt"

func main() {
	f(3)
}

// func f(x int) {
// 	fmt.Printf("f(%d)\n", x+0/x)
// 	defer fmt.Printf("defer %d\n", x)
// 	f(x - 1)
// }

// A more resonable example
// since recursion + defer + panic is unknown how they run internally
// and defer is called at the end of the function execution.
func f(x int) {
	for {
		fmt.Printf("f(%d)\n", x+0/x)
		defer fmt.Printf("defer %d\n", x)
		f(x - 1)
	}
}
