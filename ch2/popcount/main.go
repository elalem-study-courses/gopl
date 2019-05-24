package main

// package description
import (
	"fmt"

	"./popcount"
)

func main() {
	fmt.Println(popcount.PopCount(0xFFFFFF))
}
