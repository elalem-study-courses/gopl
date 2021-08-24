package main

import "fmt"

type Celesius float64
type Fahrenheit float64

func (c Celesius) String() string {
	return fmt.Sprintf("%g°C", c)
}

const (
	AbsoluteZero Celesius = -273.15
	FreezingC    Celesius = 0
	BoilingC     Celesius = 100
)

func CToF(c Celesius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celesius {
	return Celesius((f - 32) * 5 / 9)
}

func main() {
	fmt.Printf("%s\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
}
