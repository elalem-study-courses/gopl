package eval

// literal is a numeric constant e.g., 3.141
type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (l literal) Check(_ map[Var]bool) error {
	return nil
}
