package main

import "fmt"

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	fmt.Println(remove(data, 2))
}
