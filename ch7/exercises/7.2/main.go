package main

import (
	"fmt"
	"io"
	"os"
)

type WrappedWriter struct {
	writer io.Writer
	count  int64
}

func (ww *WrappedWriter) Write(p []byte) (int, error) {
	ww.count += int64(len(p))
	return ww.writer.Write(p)
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	wrappedWriter := &WrappedWriter{writer: w}
	return wrappedWriter, &wrappedWriter.count
}

func main() {
	cw, cnt := CountingWriter(os.Stdout)
	fmt.Fprintf(cw, "Hello from the other side\n")
	fmt.Println(*cnt)
}
