package main

import "fmt"

func f(x int) {
	panic(x)
}

func main() {
	defer func() {
		fmt.Println(recover())
	}()

	f(3)
}
