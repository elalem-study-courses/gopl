package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	fmt.Printf("%T\n", f)
	// c := w.(*bytes.Buffer)
	// fmt.Printf("%T\n", c)
	rw := w.(io.ReadWriter)
	fmt.Printf("%T\n", rw)

	f, ok := w.(*os.File)
	fmt.Printf("%T %t\n", f, ok)
	c, ok := w.(*bytes.Buffer)
	fmt.Printf("%T %t\n", c, ok)

	_, err := os.Open("kdnfkfndf")
	fmt.Println(os.IsNotExist(err))

	dummy(w)
}

func dummy(w io.Writer) {
	fmt.Printf("%T %[1]v\n", w)
	nw, ok := w.(*os.File)
	fmt.Printf("%T %[1]v %t\n", nw, ok)
	bw, ok := w.(*bytes.Buffer)
	fmt.Printf("%T %[1]v %t\n", bw, ok)
}
