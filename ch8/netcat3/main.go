// package description
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	// go mustCopy(os.Stdout, conn)
	go func() {
		io.Copy(os.Stdout, conn)
		// conn.CloseWrite()
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	// conn.CloseRead()
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
