package main

import (
	"bufio"
	"fmt"
	"io"
)

type LimitReader struct {
	str   string
	limit int
}

func (lr *LimitReader) Read(p []byte) (int, error) {
	stopAt := lr.limit
	if len(lr.str) < lr.limit {
		stopAt = len(lr.str)
	}
	copy(p, []byte(lr.str[:stopAt]))
	return stopAt, io.EOF
}

func main() {
	lr := &LimitReader{str: "Hello world", limit: 4}
	scan := bufio.NewReader(lr)
	fmt.Println(scan.ReadLine())
}
