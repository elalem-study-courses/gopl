package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	Protocol      = "tcp"
	Host          = "localhost:8000"
	DirectoryRoot = "/tmp"
)

func Run() {
	listener, err := net.Listen(Protocol, Host)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server running on %s:%s\n", Protocol, Host)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("connection received")
	connection := openSession()
	writer := bufio.NewWriter(conn)
	reader := bufio.NewScanner(conn)
	fmt.Println(connection.SessionID)
	writer.WriteString(connection.SessionID + "\n")
	for reader.Scan() {
		command := reader.Text()
		commands[command].handle(conn)
	}
}
