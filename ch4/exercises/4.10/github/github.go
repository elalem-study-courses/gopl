package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssueURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return &result, nil
}

func SearchIssuesAgeCategorized(terms []string) (map[string]*IssuesSearchResult, error) {
	result, err := SearchIssues(terms)
	if err != nil {
		return nil, err
	}

	categorizedResult := map[string]*IssuesSearchResult{
		"less than a month": {},
		"less than a year":  {},
		"more than a year":  {},
	}

	for _, issue := range result.Items {
		if issue.CreatedAt.After(time.Now().AddDate(0, -1, 0)) {
			addIssueToIssueSearchResult(categorizedResult["less than a month"], issue)
		} else if issue.CreatedAt.After(time.Now().AddDate(-1, 0, 0)) {
			addIssueToIssueSearchResult(categorizedResult["less than a year"], issue)
		} else {
			addIssueToIssueSearchResult(categorizedResult["more than a year"], issue)
		}
	}

	return categorizedResult, nil
}

func addIssueToIssueSearchResult(result *IssuesSearchResult, issue *Issue) {
	result.Items = append(result.Items, issue)
	result.TotalCount++
}
