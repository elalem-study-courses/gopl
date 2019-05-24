// package description
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

const (
	SHA256 = 1 << iota
	SHA384
	SHA512
)

var algorithmInUse = SHA256

var sha384Flag = flag.Bool("sha384", false, "Use SHA-384")
var sha512Flag = flag.Bool("sha512", false, "Use SHA-512")

func main() {
	flag.Parse()
	if *sha384Flag {
		algorithmInUse = 0
		algorithmInUse |= SHA384
	}

	if *sha512Flag {
		algorithmInUse = 0
		algorithmInUse |= SHA512
	}

	str := os.Args[1]

	var hashedStr string

	switch algorithmInUse {
	case SHA256:
		hashedStr = fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
	case SHA384:
		hashedStr = fmt.Sprintf("%x", sha512.Sum384([]byte(str)))
	case SHA512:
		hashedStr = fmt.Sprintf("%x", sha512.Sum512([]byte(str)))
	}
	fmt.Println(hashedStr)
}
