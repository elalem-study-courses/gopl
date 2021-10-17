package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree

	for _, v := range values {
		root = add(root, v)
	}

	fmt.Println(root)

	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value <= t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func (t *tree) String() string {
	buf := &bytes.Buffer{}

	if t != nil {
		buf.WriteString(t.left.String())
		buf.WriteString(fmt.Sprintf("|%d|", t.value))
		buf.WriteString(t.right.String())
	}

	return buf.String()
}

func main() {
	arr := []int{534, 43, 43, 43, 34, 324355, 45, 54, 6, 665, 5, 657556, 65, 65, 6, 56, 56, 65}
	Sort(arr)
	fmt.Println(arr)
}
