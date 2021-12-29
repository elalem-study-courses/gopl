package server

import (
	"time"

	"github.com/google/uuid"
)

type Session map[string]Connection

var session = make(Session)

func openSession() Connection {
	id := uuid.NewString()
	connection := Connection{
		SessionID: id,
		CreatedAt: time.Now(),
	}
	session[id] = connection

	return connection
}

func closeSession(id string) {
	delete(session, id)
}
