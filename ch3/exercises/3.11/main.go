package main

import (
	"fmt"
	"os"
	"strings"
)

func comma(s string) string {
	decimalPoint := strings.LastIndex(s, ".")
	wholeNumberPart := s
	fractionalPart := ""
	decimalPointChar := ""
	if decimalPoint >= 0 {
		wholeNumberPart = s[:decimalPoint]
		fractionalPart = s[decimalPoint+1:]
		decimalPointChar = "."
	}
	return commaRecursive(wholeNumberPart) + decimalPointChar + fractionalPart
}

func commaRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaRecursive(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Println(comma(os.Args[1]))
}
