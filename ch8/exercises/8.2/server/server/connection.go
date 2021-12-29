package server

import "time"

type Connection struct {
	SessionID string
	CreatedAt time.Time
}
