package eval

// Var identifies a variable e.g., x
type Var string

// Eval for Var type simple does a simple lookup with the variable name
func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}
