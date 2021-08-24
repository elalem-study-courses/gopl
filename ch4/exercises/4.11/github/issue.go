package github

import "time"

type Issue struct {
	URL       string
	Id        int64
	Number    int64
	Title     string
	State     string
	CreatedAt time.Time `json:"created_at"`
	ClosedAt  time.Time `json:"closed_at"`
	User      User
	Assignee  User
	Assignees []User
	Labels    []Label
}
