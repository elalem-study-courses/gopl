package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"./github"
)

var indexTemplate = template.Must(template.New("index.html").ParseFiles("templates/index.html"))
var showTemplate = template.Must(template.New("show.html").ParseFiles("templates/show.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := indexTemplate.Execute(w, nil); err != nil {
		log.Printf("indexHandler: %v\n", err)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	terms := r.FormValue("terms")

	issues, err := github.SearchIssues(strings.Split(terms, " "))
	if err != nil {
		log.Printf("searchHandler: %v\n", err)
	}
	if err := showTemplate.Execute(w, issues); err != nil {
		log.Printf("searchHandler: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/search", searchHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
