package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"./tracks"
)

var indexPage = template.Must(template.New("index.html").ParseFiles("templates/index.html"))
var sortParge = template.Must(template.New("sort.html").ParseFiles("templates/sort.html"))

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := indexPage.Execute(w, tracks.SelectedTracks.Tracks); err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/sort/", func(w http.ResponseWriter, r *http.Request) {
		col, _ := strconv.Atoi(r.URL.Query().Get("col"))
		tracks := tracks.Sort(col)

		if err := indexPage.Execute(w, tracks); err != nil {
			log.Println(err)
		}
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
