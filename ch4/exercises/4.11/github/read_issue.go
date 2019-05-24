// package description
package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const URL = "https://api.github.com/repos/%s/%s/issues/%s"

type Labels []Label

type Issue struct {
	HTMLURL string `json:"html_url"`
	User    *User
	Labels  *Labels
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
url: %s
user_name: %s
user_url: %s
labels: %v
	`

	return fmt.Sprintf(template, issue.HTMLURL, issue.User.Login, issue.User.HTMLURL, issue.Labels)
}

func (labels Labels) String() string {
	labelsString := make([]string, 0, len(labels))
	for _, label := range labels {
		labelsString = append(labelsString, label.Name)
	}
	return "[" + strings.Join(labelsString, ", ") + "]"
}

func ReadIssue(owner, repo, issueNumber string) (*Issue, error) {
	url := fmt.Sprintf(URL, owner, repo, issueNumber)
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		log.Printf("readIssue: %v\n", err)
		return nil, err
	}

	// issue := &Issue{}
	issue := new(Issue)

	if err := json.NewDecoder(resp.Body).Decode(issue); err != nil {
		log.Printf("readIssue: %v\n", err)
		return nil, err
	}

	return issue, nil
}
