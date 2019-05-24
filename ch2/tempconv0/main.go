package main

import "fmt"

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	// fmt.Printf("%g\n", boilingF-FreezingC) // Compilation error

	var c Celsius
	var f Fahrenheit

	fmt.Println(f == 0)
	fmt.Println(c >= 0)
	// fmt.Println(f == c) // Compilation error
	fmt.Println(c == Celsius(f))

	c = FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64)

}
