package eval

import (
	"fmt"
	"strings"
)

type binary struct {
	op rune
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}

	panic(fmt.Sprintf("unsupported binary operation %q", b.op))
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-/*", b.op) {
		fmt.Errorf("unexpected binary operator %q", b.op)
	}

	if err := b.x.Check(vars); err != nil {
		return err
	}

	return b.y.Check(vars)
}
