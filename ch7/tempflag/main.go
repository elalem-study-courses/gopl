// package description
package main

import (
	"flag"
	"fmt"

	"./tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
