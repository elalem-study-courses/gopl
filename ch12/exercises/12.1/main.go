package main

import (
	"./display"
)

type dummy [1]int

var dummyMap map[dummy]int

func main() {
	dummyMap = make(map[dummy]int)
	dummyMap[dummy{1}] = 1
	dummyMap[dummy{2}] = 2
	dummyMap[dummy{3}] = 3
	dummyMap[dummy{4}] = 4
	dummyMap[dummy{5}] = 5

	display.Display("dummyMap", dummyMap)
}
