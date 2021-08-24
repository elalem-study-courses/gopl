package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	dotIndex := strings.Index(s, ".")
	decimal, fractional := s[:dotIndex], s[dotIndex:]
	for i := 0; i < len(decimal); i++ {
		if i > 0 && i%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(decimal[i])
	}

	return buf.String() + fractional
}

func main() {
	fmt.Println(comma("323223223.3221"))
}
