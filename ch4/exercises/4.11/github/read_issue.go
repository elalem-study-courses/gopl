// package description
package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const ReadIssueURL = "https://api.github.com/repos/%s/%s/issues/%s"

func ReadIssue(owner, repo, issueNumber string) (*Issue, error) {
	url := fmt.Sprintf(ReadIssueURL, owner, repo, issueNumber)
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
