package main

import (
	"fmt"
	"strings"
)

func stringsJoin(delim string, parts ...string) string {
	return strings.Join(parts, delim)
}

func main() {
	fmt.Println(stringsJoin("/", "a", "b", "c"))
}
