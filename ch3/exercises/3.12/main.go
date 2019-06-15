// package description
package main

import "fmt"

func processString(str string, m map[rune]int, factor int) {
	for _, c := range str {
		m[c] += factor
	}
}

func isAnagram(str1, str2 string) bool {
	occMap := make(map[rune]int)
	processString(str1, occMap, 1)
	processString(str2, occMap, -1)

	areAnagram := true
	for _, v := range occMap {
		if v != 0 {
			areAnagram = false
			break
		}
	}
	return areAnagram
}

func main() {
	fmt.Println(isAnagram("earth", "heart"))
}
