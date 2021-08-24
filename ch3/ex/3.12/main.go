package main

import "fmt"

func isAnagram(s1 string, s2 string) bool {
	freq := make(map[rune]int)
	for _, r := range s1 {
		freq[r]++
	}

	for _, r := range s2 {
		freq[r]--
		if freq[r] == 0 {
			delete(freq, r)
		}
	}

	return len(freq) == 0
}

func main() {
	fmt.Println(isAnagram("heart", "earth"))
}
