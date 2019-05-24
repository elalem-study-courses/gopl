package main

import "fmt"

func removeDup(arr []int) []int {
	uniqueIndex := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			arr[uniqueIndex] = arr[i]
			uniqueIndex++
		}
	}
	return arr[:uniqueIndex]
}

func main() {
	arr := []int{1, 2, 2, 2, 3, 3, 4, 5}
	fmt.Println(removeDup(arr))
}
