package main

import (
	"flag"
	"fmt"
)

type Celesius float64
type Fahrenheit float64

type CelesiusFlagType struct {
	Celesius
}

func (f *CelesiusFlagType) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		f.Celesius = Celesius(value)
		return nil
	case "F":
		f.Celesius = FToC(Fahrenheit(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

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

func CelesiusFlag(name string, value Celesius, usage string) *Celesius {
	f := CelesiusFlagType{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celesius
}

var temp = CelesiusFlag("celesius", 20.0, "the temperature")

func main() {
	fmt.Printf("%s\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))

	flag.Parse()
	fmt.Printf("User input %v\n", *temp)
}
