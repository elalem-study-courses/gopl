// package description
package eval

import "fmt"

// Env is a map that maps the variable name to its value
type Env map[Var]float64

// Expr is an interface that has Eval method that computes a certain expression
type Expr interface {

	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	fmt.Stringer
}
