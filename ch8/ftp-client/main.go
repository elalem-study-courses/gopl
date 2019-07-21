package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%v", os.Getenv("PORT")))
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go send(conn, done)
	go receive(conn, done)

	<-done
}

func send(conn net.Conn, done chan struct{}) {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Print(err)
		}

		conn.Write(line)
		fmt.Printf("Sent command %s\n", line)
	}

	conn.Close()
	done <- struct{}{}
}

func receive(conn net.Conn, done chan struct{}) {
	for {
		recv := make([]byte, 128)
		_, err := conn.Read(recv)
		if err == io.EOF {
			conn.Close()
			break
		} else if err != nil {
			log.Print(err)
		}

		fmt.Printf("%s\n", recv)
	}

	fmt.Println("Connection closed by server")
	done <- struct{}{}
}
