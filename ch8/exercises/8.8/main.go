package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	handleConn(conn)
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	inputReceived := make(chan struct{})
	done := make(chan struct{})

	go func() {
		select {
		case <-inputReceived:
		case <-time.After(10 * time.Second):
			fmt.Fprintln(c, "Terminating connection because it went idle for 10 seconds")
			done <- struct{}{}
		}
	}()

	go func() {
		for input.Scan() {
			go echo(c, input.Text(), 1*time.Second)
		}
		done <- struct{}{}
	}()

	<-done
	c.Close()
}
