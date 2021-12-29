package server

import "net"

type Command interface {
	handle(conn net.Conn)
}

var commands = make(map[string]Command)
