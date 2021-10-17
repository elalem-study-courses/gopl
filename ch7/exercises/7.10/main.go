package main

import "fmt"

type ints []int

func (i ints) Less(a, b int) bool {
	return i[a] < i[b]
}

func isPalindrome(nums ints) bool {
	for i := range nums {
		if nums.Less(i, len(nums)-1-i) || nums.Less(len(nums)-1-i, i) {
			return false
		}
	}

	return true
}

func main() {
	p := []int{1, 2, 3, 3, 2, 1}
	np := []int{1, 2, 2, 2}

	fmt.Println(isPalindrome(p), isPalindrome(np))
}
