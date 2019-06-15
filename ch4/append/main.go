package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	x := []int{}
	x = appendInt(x, 1)
	x = appendInt(x, 2)
	x = appendInt(x, 3)
	x = appendInt(x, 4)
	x = appendInt(x, 5)
	x = append(x, 1, 2, 3, 4, 5)
	x = append(x, x...)
	fmt.Printf("%v\n", x)

}
