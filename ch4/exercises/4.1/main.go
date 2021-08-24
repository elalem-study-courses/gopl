package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n%d\n", c1, c2, c1 == c2, c1, countDiffBytes(c1, c2))
}

func countDiffBytes(c1, c2 [32]byte) int {
	diff := 0

	for i := 0; i < 32; i++ {
		d1, d2 := c1[i], c2[i]
		for d1 != 0 || d2 != 0 {
			if d1%2 != d2%2 {
				diff++
			}
			d1 /= 2
			d2 /= 2
		}
	}

	return diff
}
