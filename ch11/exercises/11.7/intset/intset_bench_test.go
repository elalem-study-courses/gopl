package intset

import (
	"math/rand"
	"testing"
)

func newIntSet() *IntSet { return &IntSet{} }
func newRandomIntSet() *IntSet {
	intset := &IntSet{}
	for i := 0; i < rand.Intn(1000); i++ {
		intset.Add(rand.Intn(1 << 24))
	}

	return intset
}

// Performance is crap for large numbers i.e. 1 << 32 - 1
func BenchmarkAdd(b *testing.B) {
	intset := newIntSet()
	for i := 0; i < b.N; i++ {
		intset.Add(rand.Intn(1 << 24))
	}
}

func BenchmarkHas(b *testing.B) {
	intset := newIntSet()
	for i := 0; i < 1000000; i++ {
		intset.Add(rand.Intn(1 << 24))
	}

	for i := 0; i < b.N; i++ {
		intset.Has(rand.Intn(1 << 24))
	}
}

func BenchmarkUnionWith(b *testing.B) {
	intset := newIntSet()

	for i := 0; i < b.N; i++ {
		intset.UnionWith(newRandomIntSet())
	}
}
