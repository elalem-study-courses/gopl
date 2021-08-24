package github

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
