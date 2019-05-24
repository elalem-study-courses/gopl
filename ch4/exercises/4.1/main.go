// package description
package main

import (
	"crypto/sha256"
	"fmt"
)

func xor(hash1, hash2 [32]byte) [32]byte {
	var diff [32]byte
	for index := range hash1 {
		diff[index] = hash1[index] ^ hash2[index]
	}
	return diff
}

func countArrayBits(arr [32]byte) int {
	cnt := 0
	for _, value := range arr {
		cnt += countBits(value)
	}
	return cnt
}

func countBits(n byte) int {
	cnt := 0
	for n > 0 {
		n &^= (n & -n)
		cnt++
	}
	return cnt
}

func main() {
	hash1 := sha256.Sum256([]byte("x"))
	hash2 := sha256.Sum256([]byte("X"))
	diff := xor(hash1, hash2)
	fmt.Println(countArrayBits(diff))
}
