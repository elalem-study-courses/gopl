package main

import "fmt"

func main() {
	fmt.Println(f(10))
}

type Integer struct {
	value int
}

func f(x int) (ret Integer) {
	defer func() {
		ret = recover().(Integer)
	}()
	panic(Integer{value: x})
	return
}
