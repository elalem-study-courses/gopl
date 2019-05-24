package main

import "fmt"

func rotate(arr []int, offset int) []int {
	newStart := len(arr) - (offset % len(arr))

	rotatedArr := make([]int, len(arr))

	for i, v := range arr {
		rotatedArr[(newStart+i)%len(arr)] = v
	}

	return rotatedArr
}

func main() {
	arr := []int{1, 2, 3, 4, 5} // 3 4 5 1 2
	arr = rotate(arr, 2)
	fmt.Println(arr)
}
