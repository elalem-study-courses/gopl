package eval

import (
	"bytes"
	"testing"
)

func TestEval(t *testing.T) {
	buf := new(bytes.Buffer)
	buf.WriteString("((a + b + c) / 5) + 8 / 2.0 * aa")
	t.Log(Eval(buf, map[string]float64{"a": 1, "b": 2, "c": 2, "aa": 2}))
}