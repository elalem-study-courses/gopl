package github

type User struct {
	Id      int64
	Login   string
	URL     string
	HTMLURL string `json:"html_url"`
	Type    string
}
