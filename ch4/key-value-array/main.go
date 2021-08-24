package main

import "fmt"

type Currency uint8

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

var symbol = [...]string{EUR: "E", GBP: "G", USD: "U", RMB: "R"}
var dum = [...]int{99: -1}

func main() {
	fmt.Println(symbol)
	fmt.Println(len(dum))
}
