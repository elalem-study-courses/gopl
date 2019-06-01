package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const CreateIssueURL = "https://api.github.com/repos/%s/%s/issues"

func CreateIssue(owner, repo, title, body string, labels []string) (*Issue, error) {
	url := fmt.Sprintf(CreateIssueURL, owner, repo)
	postBodyStruct := &struct {
		Title  string   `json:"title"`
		Body   string   `json:"body"`
		Labels []string `json:"labels"`
	}{
		Title:  title,
		Body:   body,
		Labels: labels,
	}

	postBody, _ := json.Marshal(postBodyStruct)

	log.Printf("Post %s\n", url)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(postBody))
	defer resp.Body.Close()

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		fmt.Print(err)
	}

	if err != nil {
		log.Fatalf("CreateIssue: %v\n", err)
	}

	issue := new(Issue)

	if err := json.NewDecoder(resp.Body).Decode(issue); err != nil {
		fmt.Printf("CreateIssue: %v\n", err)
	}

	fmt.Println(issue)

	return issue, nil
}

func stringArrayToLabels(labelsStr []string) *Labels {
	labels := make(Labels, 0, len(labelsStr))
	for _, label := range labelsStr {
		labels = append(labels, Label{Name: label})
	}

	return &labels
}
