// package description
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for n > 0 {
		sliceLength := 3
		if n%3 != 0 {
			sliceLength = n % 3
		}
		buf.WriteString(s[len(s)-n : len(s)-n+sliceLength])
		n -= sliceLength
		if n > 0 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("123456789"))
}
