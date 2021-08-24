package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseUrl = "https://api.github.com"
)

var (
	client = &http.Client{}
)

func CreateIssue(username, accessToken, repo, title, description string) (*Issue, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/repos/%s/%s/issues", baseUrl, username, repo), bytes.NewBuffer([]byte(
		fmt.Sprintf(`{"title": "%s"}`, title))))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	res, err := client.Do(req)
	if err != nil {
		res.Body.Close()
		return nil, err
	}

	issue := Issue{}
	if err := json.NewDecoder(res.Body).Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func CloseIssue(username, accessToken, repo, issueNumber string) (*Issue, error) {
	req, err := http.NewRequest(
		"Patch",
		fmt.Sprintf("%s/repos/%s/%s/issues/%s", baseUrl, username, repo, issueNumber),
		bytes.NewBuffer([]byte(fmt.Sprintf(`{"state": "closed"}`))))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+accessToken)

	res, err := client.Do(req)
	resBody, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(resBody))
	if err != nil {
		res.Body.Close()
		return nil, err
	}

	issue := Issue{}
	if err := json.NewDecoder(res.Body).Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func ReadIssue(username, repo, issueNumber string) (*Issue, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/repos/%s/%s/issues/%s", baseUrl, username, repo, issueNumber), nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		res.Body.Close()
		return nil, err
	}

	issue := Issue{}
	if err := json.NewDecoder(res.Body).Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}

func pathFor(path string) string {
	return baseUrl + path
}
