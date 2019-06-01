package github

import (
	"fmt"
	"strings"
)

type Labels []Label

type Issue struct {
	HTMLURL string `json:"html_url", omitempty`
	User    *User  `omitempty`
	Labels  *Labels
	Title   string
	Body    string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Label struct {
	Name string
}

func (issue *Issue) String() string {
	template := `
title: %s
url: %s
user_name: %s
user_url: %s
labels: %v
Body: %s
	`

	return fmt.Sprintf(template, issue.Title, issue.HTMLURL, issue.User.Login, issue.User.HTMLURL, issue.Labels, issue.Body)
}

func (labels Labels) String() string {
	labelsString := make([]string, 0, len(labels))
	for _, label := range labels {
		labelsString = append(labelsString, label.Name)
	}
	return "[" + strings.Join(labelsString, ", ") + "]"
}
