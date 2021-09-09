package main

import (
	"fmt"

	"./intset"
)

func main() {
	var x, y intset.IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	fmt.Printf("%#v\n", x.Elems())

	fmt.Println("-----------------------------")

	var a, b intset.IntSet
	a.Add(1)
	a.Add(2)
	a.Add(3)
	b.Add(2)
	b.Add(3)
	b.Add(4)
	// a.IntersectWith(&b) // {2, 3}
	// a.DifferenceWith(&b) // {1}
	// a.SymmetricDifferenceWith(&b) // {1, 4}
	a.Clear()
	fmt.Println(&a)
}
