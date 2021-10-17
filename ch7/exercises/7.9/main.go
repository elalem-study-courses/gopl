package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var (
	indexPage = template.Must(template.ParseFiles("./index.html"))
)

var (
	lessMethods []func(t Tracks, i, j int) bool = []func(t Tracks, i, j int) bool{
		func(t Tracks, i, j int) bool { return t[i].Title < t[j].Title },
		func(t Tracks, i, j int) bool { return t[i].Artist < t[j].Artist },
		func(t Tracks, i, j int) bool { return t[i].Album < t[j].Album },
		func(t Tracks, i, j int) bool { return t[i].Year < t[j].Year },
		func(t Tracks, i, j int) bool { return t[i].Length < t[j].Length },
	}
	equalMethods []func(t Tracks, i, j int) bool = []func(t Tracks, i, j int) bool{
		func(t Tracks, i, j int) bool { return t[i].Title == t[j].Title },
		func(t Tracks, i, j int) bool { return t[i].Artist == t[j].Artist },
		func(t Tracks, i, j int) bool { return t[i].Album == t[j].Album },
		func(t Tracks, i, j int) bool { return t[i].Year == t[j].Year },
		func(t Tracks, i, j int) bool { return t[i].Length == t[j].Length },
	}

	sortingColumns = []int{}
)

type Tracks []*Track

func (t Tracks) Less(i, j int) bool {
	for _, column := range sortingColumns {
		if equalMethods[column](t, i, j) {
			continue
		}

		return lessMethods[column](t, i, j)
	}

	return lessMethods[0](t, i, j)
}
func (t Tracks) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t Tracks) Len() int      { return len(t) }
func (t Tracks) Sort(column int) {
	l := len(sortingColumns)
	if l > 5 {
		l = 5
	}
	sortingColumns = append([]int{column}, sortingColumns[0:l]...)

	sort.Sort(t)
}

var tracks = Tracks{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}

	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', tabwriter.StripEscape)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}

	tw.Flush()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := indexPage.Execute(w, tracks); err != nil {
			fmt.Fprintf(w, "Error occured %v", err)
		}
	})

	http.HandleFunc("/sort", func(w http.ResponseWriter, r *http.Request) {
		column, _ := strconv.ParseInt(r.URL.Query().Get("column"), 10, 32)
		fmt.Println(column)
		tracks.Sort(int(column))
		http.Redirect(w, r, "/", http.StatusFound)
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
