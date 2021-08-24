package main

import (
	"bufio"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var bitSize = flag.Int("bitsize", 384, "Determines the size of the hash")

func main() {
	flag.Parse()
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if *bitSize == 384 {
			fmt.Printf("%x\n", sha512.Sum384(input.Bytes()))
		} else if *bitSize == 512 {
			fmt.Printf("%x\n", sha512.Sum512(input.Bytes()))
		}
	}
}
