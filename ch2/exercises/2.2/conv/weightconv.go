package conv

import "fmt"

type Kilo float64
type Pound float64

const PoundKiloRate = 2.20462

func KiloToPound(k Kilo) Pound {
	return Pound(k * PoundKiloRate)
}

func PoundToKilo(p Pound) Kilo {
	return Kilo(p / PoundKiloRate)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g Pound", p)
}

func (k Kilo) String() string {
	return fmt.Sprintf("%g Kilo", k)
}
