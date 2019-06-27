package main

import (
	"fmt"
	"log"

	"./eval"
)

func main() {
	expr, err := eval.Parse("(x+y)*5+sin(x)+pow(x,2)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(expr)
}
