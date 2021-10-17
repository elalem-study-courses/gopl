package main

import (
	"fmt"
	"os"
	"sort"
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
	sortingColumns = append(sortingColumns, column)
	if len(sortingColumns) > 4 {
		sortingColumns = sortingColumns[1:]
	}

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
	var column int
	printTracks(tracks)
	fmt.Println("-------------------------------------------------------------")
	for {
		fmt.Print("Sort by Column (0-4): ")
		fmt.Scanf("%d", &column)
		tracks.Sort(column)
		printTracks(tracks)
		fmt.Println("-------------------------------------------------------------")
	}
}
