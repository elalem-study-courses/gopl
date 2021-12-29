package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer c.Close()

	input := bufio.NewScanner(c)
	inputReadChan := make(chan struct{})
	endChan := make(chan struct{})

	go func() {
		for {
			select {
			case <-inputReadChan:
			case <-time.After(10 * time.Second):
				c.Write([]byte("Closed due inactivity for 10 seconds\n"))
				endChan <- struct{}{}
			}
		}
	}()

	go func() {
		for input.Scan() {
			go echo(c, input.Text(), 4*time.Second)
		}

		endChan <- struct{}{}
	}()

	<-endChan
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go handleConn(conn)
	}
}
