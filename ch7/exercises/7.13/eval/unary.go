package eval

import (
	"fmt"
	"strings"
)

// unary represents a unary operator expression e.g., -x
type unary struct {
	op rune
	x  Expr
}

// unary Evaluate its operand recursively e.g., -(1 + (3 * 4))
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsuppoerted unary operator: %q", u.op))
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary operator %q", u.op)
	}
	return u.x.Check(vars)
}

func (u unary) String() string {
	return fmt.Sprintf("%c %s", u.op, u.x.String())
}
