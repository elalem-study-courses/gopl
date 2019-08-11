package main

import (
	"fmt"

	"./intset"
)

var x, y intset.IntSet

func main() {
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(&y)

	x.UnionWith(&y)
	fmt.Println(&x)
	fmt.Println(x.Has(9), x.Has(123))
}
