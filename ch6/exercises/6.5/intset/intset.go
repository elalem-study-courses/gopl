package intset

import (
	"bytes"
	"fmt"
	"math"
)

const (
	// WordLength based on platform.
	// Basically If you can shift an unsigned int by 63 position
	// So this platform can hold 64-bit and take the result of the bitwise shift right
	// which will be 0 or 1 for 32-bit and 64-bit architecture respectively
	// and apply a shift left which is basically multiplying by 1 or 2 respectively.
	WordLength = 32 << (uint(0) >> 63)
)

type IntSet struct {
	words []int
	count int
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/WordLength, uint(x%WordLength)
	return word < len(s.words) && (s.words[word]&(1<<bit)) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/WordLength, uint(x%WordLength)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << bit)
	s.count++
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < WordLength; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", WordLength*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	return s.count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/WordLength, uint(x%WordLength)
	if word < len(s.words) {
		s.words[word] &^= (1 << bit)
	}
}

func (s *IntSet) Clear() {
	s.words = s.words[:0]
	s.count = 0
}

func (s *IntSet) Copy() *IntSet {
	replica := IntSet{words: make([]int, len(s.words))}
	copy(replica.words, s.words)
	replica.count = s.count
	return &replica
}

func (s *IntSet) AddAll(elems ...int) {
	for _, elem := range elems {
		s.Add(elem)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	if len(s.words) > len(t.words) {
		s.words = s.words[:len(t.words)]
	}
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := 0; i < len(s.words); i++ {
		s.words[i] &^= t.words[i]
	}
}

func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i := 0; i < len(t.words); i++ {
		if i < len(s.words) {
			s.words[i] ^= t.words[i]
		} else {
			s.words = append(s.words, t.words[i])
		}
	}
}

func (s *IntSet) Elems() []int {
	elements := make([]int, 0)
	for word, bits := range s.words {
		for bits > 0 {
			lsb := bits & -bits
			bits -= lsb
			bitPosition := int(math.Log2(float64(lsb)))
			elements = append(elements, int(word)*WordLength+bitPosition)
		}
	}
	return elements
}
