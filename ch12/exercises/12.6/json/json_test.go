package json

import (
	"bytes"
	"testing"
)

func TestEncodeSimpleStruct(t *testing.T) {
	simpleStruct := struct {
		x int
		y int
	}{1, 2}
	var buf bytes.Buffer

	Encode(&buf, simpleStruct)

	if buf.String() != `{"x":1,"y":2}` {
		t.Errorf("Encode(%v, %#v) expected %q got %q", &buf, simpleStruct, `{"x":1,"y":2}`, buf.String())
	}
}
