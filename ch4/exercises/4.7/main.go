package main

import "fmt"

func rev(bytes []byte) []byte {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return bytes
}

func main() {
	str := "This is sparta"
	fmt.Printf("%q\n", rev([]byte(str)))
}
