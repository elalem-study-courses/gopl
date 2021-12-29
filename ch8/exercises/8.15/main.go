package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	IdleTime = 5 * time.Minute
)

type Client struct {
	ID      string
	channel chan<- string
	name    string
}

var (
	entering = make(chan *Client)
	leaving  = make(chan *Client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[string]*Client)
	for {
		select {
		case msg := <-messages:
			for _, cli := range clients {
				cli.channel <- msg
			}
		case cli := <-entering:
			clients[cli.ID] = cli
			activeClients := make([]string, 0)
			for _, cli := range clients {
				activeClients = append(activeClients, cli.name)
			}
			for _, cli := range clients {
				cli.channel <- fmt.Sprintf("%d users available...", len(activeClients))
				for idx, id := range activeClients {
					cli.channel <- fmt.Sprintf("%d. %s", idx+1, id)
				}
			}
		case cli := <-leaving:
			delete(clients, cli.ID)
			close(cli.channel)
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	ch := make(chan string, 5)
	done := make(chan struct{})
	inputReceived := make(chan struct{})
	input := bufio.NewScanner(conn)

	go clientWriter(conn, ch)

	ch <- "Enter your name"
	input.Scan()
	username := input.Text()

	go func() {
	loop:
		for {
			select {
			case <-time.After(IdleTime):
				ch <- fmt.Sprintf("You've been disconnected because of being idle for %s", IdleTime.String())
				close(done)
				break loop
			case <-inputReceived:
			}
		}
	}()

	who := conn.RemoteAddr().String()

	cli := &Client{
		ID:      who,
		channel: ch,
		name:    username,
	}

	ch <- "You are " + who
	messages <- username + " has arrived"
	entering <- cli

	go func() {
		for input.Scan() {
			messages <- username + ": " + input.Text()
			inputReceived <- struct{}{}
		}
	}()

	<-done

	leaving <- cli
	messages <- username + " has left"
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}
