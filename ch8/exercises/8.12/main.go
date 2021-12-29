package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Client struct {
	ID      string
	channel chan<- string
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
			for id := range clients {
				activeClients = append(activeClients, id)
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
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()

	cli := &Client{
		ID:      who,
		channel: ch,
	}

	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- cli
	messages <- who + " has left"
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
