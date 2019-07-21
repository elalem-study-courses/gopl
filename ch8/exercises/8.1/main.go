// package description
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

var (
	Location *time.Location
	port     = flag.Int("port", 8080, "The port will listen too")
)

func init() {
	var err error
	Location, err = time.LoadLocation(os.Getenv("TZ"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on localhost:%d", *port)
	for {
		conn, err := listener.Accept()
		log.Printf("Accepted connection from %v", conn.RemoteAddr())
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(Location).Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
