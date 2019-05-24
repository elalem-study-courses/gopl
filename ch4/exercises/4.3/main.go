package main

import "fmt"

func reverse(arr *[]int) {
	arrV := *arr
	n := len(arrV)
	for i := 0; i*2 < n; i++ {
		arrV[i], arrV[n-i-1] = arrV[n-i-1], arrV[i]
	}
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	reverse(&x)
	fmt.Println(x)

}
