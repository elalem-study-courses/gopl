package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"../constants"
	"../github"
)

var (
	indexTemplate = template.Must(template.ParseFiles(constants.PublicPath + "/index.html"))
	showTemplate  = template.Must(template.ParseFiles(constants.PublicPath + "/show.html"))
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result, err := github.SearchIssues(strings.Split(r.URL.Query().Get("q"), " "))
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		if err := indexTemplate.Execute(w, result); err != nil {
			fmt.Fprintln(w, err)
		}

	})

	mux.HandleFunc("/issues/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("api_url")
		issue, err := github.GetIssueDetails(url)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		if err := showTemplate.Execute(w, issue); err != nil {
			fmt.Fprintln(w, err)
		}
	})

	return mux
}
