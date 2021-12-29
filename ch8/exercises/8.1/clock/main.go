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

var port = flag.Int("port", 80, "Port of the clock")

func main() {
	flag.Parse()

	timezone, ok := os.LookupEnv("TZ")
	if !ok {
		log.Fatal("TZ Environment variable is required")
	}

	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn, location)
	}
}

func handleConn(conn net.Conn, timezone *time.Location) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, fmt.Sprintf("%s: %s\n", timezone.String(), time.Now().In(timezone).Format("03:04:05PM")))
		if err != nil {
			break
		}

		time.Sleep(1 * time.Second)
	}
}
