package intset

import (
	"math"
	"testing"
)

func TestHas(t *testing.T) {
	tests := []struct {
		input   int
		context *IntSet
		want    bool
	}{
		{
			5, &IntSet{}, false,
		},
		{
			5, &IntSet{words: []uint64{32}}, true,
		},
		{
			5, &IntSet{words: []uint64{math.MaxUint64 &^ 32}}, false,
		},
	}

	for _, test := range tests {
		if got := test.context.Has(test.input); got != test.want {
			t.Errorf("%#v.Has(%d) got %t, want %t", test.context, test.input, got, test.want)
		}
	}
}

func TestAdd(t *testing.T) {
	set := &IntSet{}
	for i := 0; i < 1e5; i++ {
		if set.Has(i) {
			t.Errorf("set.Has(%d) = true, want false", i)
		}
		set.Add(i)
		if !set.Has(i) {
			t.Errorf("set.Has(%d) = false, want true", i)
		}
	}
}

func TestUnionWith(t *testing.T) {
	tests := []struct {
		input   *IntSet
		context *IntSet
		output  *IntSet
	}{
		{
			input:   &IntSet{words: []uint64{3}},
			context: &IntSet{words: []uint64{12}},
			output:  &IntSet{words: []uint64{15}},
		},
	}

	for _, test := range tests {
		test.context.UnionWith(test.input)
		if len(test.context.words) != len(test.output.words) {
			t.Errorf("Expected len(test.context.words) = len(test.output.words) = %d, got %d", len(test.output.words), len(test.context.words))
		}

		for i := range test.context.words {
			if test.context.words[i] != test.output.words[i] {
				t.Errorf("Expected test.contet.words[i] to equal %d got %d", test.output.words[i], test.context.words[i])
			}
		}
	}
}
