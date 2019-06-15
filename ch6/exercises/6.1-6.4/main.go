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

	x.Remove(123)
	fmt.Println(&x)
	x.Remove(42)
	fmt.Println(&x)
	x.Remove(0xFFFFFFFF)
	fmt.Println(&x)
	fmt.Println(&x, x.Len())
	x.Add(42)
	fmt.Println(&x, x.Len())
	fmt.Println(&y, y.Len())
	y.Clear()
	fmt.Println(&y, y.Len())
	y.Add(1)
	fmt.Println(&y, y.Len())

	r := x.Copy()

	fmt.Println(r)
	fmt.Printf("&x = %p, &r = %p\n", &x, &r)

	// Testing Intersection
	x.Clear()
	y.Clear()

	x.AddAll(1, 2, 3)
	y.AddAll(1, 4, 5)
	fmt.Println(&x, &y)
	x.IntersectWith(&y)
	fmt.Println(&x)

	// Testing Difference
	x.Clear()
	y.Clear()

	x.AddAll(1, 2, 3)
	y.AddAll(1, 4, 5)
	fmt.Println(&x, &y)
	x.DifferenceWith(&y)
	fmt.Println(&x)

	// Testing Symmetric Difference
	x.Clear()
	y.Clear()

	x.AddAll(1, 2, 3)
	y.AddAll(1, 4, 5)
	fmt.Println(&x, &y)
	x.SymmetricDifferenceWith(&y)
	fmt.Println(&x)

	fmt.Printf("%#v\n", x.Elems())
}
