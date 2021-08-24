package main

import (
	"bytes"
	"fmt"
)

func Join(delim string, tokens ...string) string {
	buf := bytes.Buffer{}
	for i, token := range tokens {
		if i > 0 {
			buf.WriteString(delim)
		}

		buf.WriteString(token)
	}

	return buf.String()
}

func main() {
	fmt.Println(Join(", ", "1", "2", "3", "4", "5"))
}
