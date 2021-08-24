package github

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
