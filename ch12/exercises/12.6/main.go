package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	myJSON "./json"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
	Earnings        float64
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    true,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
		Earnings: 545451212.1515123,
	}

	var buf bytes.Buffer
	myJSON.Encode(&buf, strangelove)
	fmt.Println(buf.String())
	movie := Movie{}
	if err := json.Unmarshal(buf.Bytes(), &movie); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", movie)
}
