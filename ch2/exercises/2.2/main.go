// package description
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./conv"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		for _, arg := range args {
			value, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "2.2: %v\n", err)
				os.Exit(1)
			}
			printValue(value)
		}
	} else {
		fmt.Println("Enter list of numbers:")
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			arg := input.Text()
			value, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "2.2: %v\n", err)
				os.Exit(1)
			}
			printValue(value)
		}
	}
}

func printValue(v float64) {
	f := conv.Fahrenheit(v)
	c := conv.Celsius(v)
	k := conv.Kilo(v)
	p := conv.Pound(v)
	m := conv.Metre(v)
	feet := conv.Foot(v)

	fmt.Printf("Values for %v\n", v)
	fmt.Println("-----------------")

	fmt.Printf("%s = %s, %s = %s\n", conv.FToC(f), f, conv.CToF(c), c)
	fmt.Printf("%s = %s, %s = %s\n", conv.KiloToPound(k), k, conv.PoundToKilo(p), p)
	fmt.Printf("%s = %s, %s = %s\n", conv.MetreToFeet(m), m, conv.FeetToMetre(feet), feet)
	fmt.Println("------------------------------------------------------")
}
