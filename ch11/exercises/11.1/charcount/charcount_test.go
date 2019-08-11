package charcount

import (
	"testing"
	"unicode/utf8"
)

const (
	UTFLenSize = utf8.UTFMax + 1
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		input string
		want  struct {
			invalid int
			utfLen  [UTFLenSize]int
		}
	}{
		{input: "mohamed", want: struct {
			invalid int
			utfLen  [UTFLenSize]int
		}{
			0,
			[UTFLenSize]int{0, 7, 0, 0, 0},
		}},
	}

	for _, test := range tests {
		Reset()
		for _, r := range test.input {
			CharCount(r, utf8.RuneLen(r))
		}
		if test.want.invalid != Invalid || test.want.utfLen != UTFLen {
			t.Errorf("for %q got (%d, %v) want (%d, %v)", test.input, Invalid, UTFLen, test.want.invalid, test.want.utfLen)
		}
	}
}
