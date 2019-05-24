package conv

import "fmt"

const FootMetreRate = 3.28084

type Foot float64
type Metre float64

func FeetToMetre(f Foot) Metre {
	return Metre(f / FootMetreRate)
}

func MetreToFeet(m Metre) Foot {
	return Foot(m * FootMetreRate)
}

func (f Foot) String() string {
	return fmt.Sprintf("%g Feet", f)
}

func (m Metre) String() string {
	return fmt.Sprintf("%g Metre", m)
}
