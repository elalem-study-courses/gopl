package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"./ftp"
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", os.Getenv("PORT")))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	// reader := bufio.NewScanner(conn)
	session := ftp.NewSession(conn)
	session.Start()
	defer session.Close()
	for {
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Print(err)
		}

		commandEnd := strings.IndexByte(string(buf), 0x00)

		line := string(buf[:commandEnd])

		fmt.Printf("Received line %q\n", line)

		go func(session *ftp.Session, line string) {
			err := session.Handle(line)
			if err != nil {
				log.Print(err)
				fmt.Fprint(conn, err)
			}
		}(session, line)
	}
}
