// package description
package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbol := [...]string{USD: "$", EUR: "euro", GBP: "dunno", RMB: "no idea"}
	fmt.Println(RMB, symbol[RMB])
}
