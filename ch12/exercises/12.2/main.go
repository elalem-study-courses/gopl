package main

import (
	"./display"
)

type dummy [1]int

var dummyMap map[dummy]int

func main() {
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	display.Display("c", c)
}
