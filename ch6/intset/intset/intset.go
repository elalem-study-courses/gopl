package intset

import (
	"bytes"
	"fmt"
	"math"
)

const (
	WordSize = 32 << (^uint(0) >> 63)
)

type IntSet struct {
	words []uint64
	size  int
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/WordSize, uint(x%WordSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/WordSize, uint(x%WordSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit
	s.size++
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}

	s.size += t.size
}

func (s *IntSet) Len() int {
	return s.size
}

func (s *IntSet) Remove(x int) {
	word, bit := x/WordSize, x%WordSize
	if word >= len(s.words) {
		return
	}

	if (s.words[word] & (1 << bit)) != 0 {
		s.size--
	}

	s.words[word] &^= (1 << bit)
}

func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

func (s *IntSet) Copy() *IntSet {
	c := &IntSet{}
	copy(c.words, s.words)

	return c
}

func (s *IntSet) AddAll(elems ...int) {
	for _, elem := range elems {
		s.Add(elem)
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		if i >= len(t.words) {
			break
		}

		s.words[i] &= t.words[i]
	}

	if len(s.words) >= len(t.words) {
		s.words = s.words[:len(t.words)]
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i >= len(t.words) {
			break
		}

		s.words[i] &^= t.words[i]
	}
}

func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i := range s.words {
		s.words[i] = (s.words[i] & (^t.words[i])) | ((^s.words[i]) & t.words[i])
	}

	if len(s.words) < len(t.words) {
		s.words = append(s.words, t.words[len(s.words):]...)
	}
}

func (s *IntSet) Elems() []int64 {
	elems := make([]int64, 0, s.size)
	for i, word := range s.words {
		for word != 0 {
			bit := word & -word
			elems = append(elems, int64(int64(i)*WordSize+int64(math.Log2(float64(bit)))))
			word &^= bit
		}
	}

	return elems
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < WordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", WordSize*i+j)
			}
		}
	}

	buf.WriteByte('}')
	return buf.String()
}
