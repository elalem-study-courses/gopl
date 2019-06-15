// package description
package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func sort(values []int) []int {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	values = appendValues(values[:0], root)
	return values
}

func appendValues(values []int, root *tree) []int {
	if root != nil {
		values = appendValues(values, root.left)
		values = append(values, root.value)
		values = appendValues(values, root.right)
	}
	return values
}

func add(root *tree, value int) *tree {
	if root == nil {
		// shorter
		// root = &tree{value: value}
		root = new(tree)
		root.value = value
	} else if root.value > value {
		root.left = add(root.left, value)
	} else {
		root.right = add(root.right, value)
	}
	return root
}

func main() {
	x := sort([]int{4, 3, 2, 1, 5, 6, 7, 8, 6, 5, 4, 3, 2, 565, 6, 54, 6, 54, 6})
	fmt.Println(x)
}
