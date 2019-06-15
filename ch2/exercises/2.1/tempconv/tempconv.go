// package description
package tempconv

import "fmt"

type Celsius float64
type Kelvin float64

const (
	AbsoluteZeroC = -273.15
	BoilingC      = 100
	FreezingC     = 0
)

func (c Celsius) String() { fmt.Sprintf("%g C\n", c) }
func (k Kelvin) String()  { fmt.Sprintf("%g K", k) }

func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k + AbsoluteZeroC)
}
