package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

type Timezone struct {
	timezone string
	address  string
	conn     net.Conn
}

var done chan int

func main() {
	timezones := make([]Timezone, 0)
	for _, keyValue := range os.Environ() {
		tokens := strings.Split(keyValue, "=")
		if !strings.Contains(tokens[1], "localhost") {
			continue
		}
		timezone := Timezone{timezone: tokens[0], address: tokens[1]}
		conn, err := net.Dial("tcp", timezone.address)
		if err != nil {
			log.Fatal(err)
		}

		timezone.conn = conn
		timezones = append(timezones, timezone)
	}

	done = make(chan int, len(timezones))

	for _, timezone := range timezones {
		done <- 1
		go mustCopy(os.Stdout, timezone)
	}

	for range timezones {
		done <- 1
	}

	close(done)
}

func mustCopy(dst io.Writer, timezone Timezone) {
	defer func() {
		<-done
	}()
	if _, err := io.Copy(dst, timezone.conn); err != nil {
		return
	}
}
