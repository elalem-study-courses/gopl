package tracks

import (
	"sort"
	"time"
)

type Track struct {
	Title  string        `json:"title"`
	Artist string        `json:"artist"`
	Album  string        `json:"album"`
	Year   int           `json:"year"`
	Length time.Duration `json:"length"`
}

type Tracks struct {
	Tracks []Track
	clicks [5]int
}

func (t *Tracks) Len() int      { return len(t.Tracks) }
func (t *Tracks) Swap(i, j int) { t.Tracks[i], t.Tracks[j] = t.Tracks[j], t.Tracks[i] }
func (t *Tracks) Less(i, j int) bool {
	var choices []int
	for range t.clicks {
		choices = columnChoices(-1, choices)
		currentChoice := choices[len(choices)-1]
		switch currentChoice {
		case 0:
			if t.Tracks[i].Title != t.Tracks[j].Title {
				return t.Tracks[i].Title < t.Tracks[j].Title
			}
		case 1:
			if t.Tracks[i].Artist != t.Tracks[j].Artist {
				return t.Tracks[i].Artist < t.Tracks[j].Artist
			}
		case 2:
			if t.Tracks[i].Album != t.Tracks[j].Album {
				return t.Tracks[i].Album < t.Tracks[j].Album
			}
		case 3:
			if t.Tracks[i].Year != t.Tracks[j].Year {
				return t.Tracks[i].Year < t.Tracks[j].Year
			}
		case 4:
			if t.Tracks[i].Length != t.Tracks[j].Length {
				return t.Tracks[i].Length < t.Tracks[i].Length
			}
		}
	}
	return false
}

var SelectedTracks Tracks

func init() {
	SelectedTracks = Tracks{Tracks: TracksArray}
}

func ArrayContains(arr []int, target int) bool {
	for _, elem := range arr {
		if elem == target {
			return true
		}
	}
	return false
}

func columnChoices(prevChoice int, state []int) []int {
	clicks := SelectedTracks.clicks
	max := 0
	for idx := range clicks {
		if ArrayContains(state, idx) {
			continue
		}
		if clicks[idx] > clicks[max] {
			max = idx
		}
	}
	state = append(state, max)
	return state
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}

var TracksArray = []Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func Sort(col int) []Track {
	SelectedTracks.clicks[col]++
	tracksCopy := SelectedTracks

	sort.Sort(&tracksCopy)
	return tracksCopy.Tracks
}
