package main

import "fmt"

type Palindrome struct {
	data string
}

func (p *Palindrome) Len() int           { return len(p.data) }
func (p *Palindrome) Less(i, _ int) bool { return p.data[i] != p.data[p.Len()-i-1] }
func (p *Palindrome) Swap(_, _ int)      {}

func main() {
	palin := Palindrome{"aibohphobia"}
	isPalindrome := true
	for i := 0; i < palin.Len(); i++ {
		if palin.Less(i, 0) {
			isPalindrome = false
			break
		}
	}
	fmt.Println(isPalindrome)
}
