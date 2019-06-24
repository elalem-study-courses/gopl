// package description
package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

var c ByteCounter

func main() {
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Mohamed"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

}
