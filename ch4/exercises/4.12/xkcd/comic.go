package xkcd

import "fmt"

type Comic struct {
	Month      string
	Number     int64 `json:"num"`
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Image      string `json:"img"`
	Title      string
	Day        string
}

func (comic *Comic) String() string {
	return fmt.Sprintf(`Number: %d
Title: %s
Safe Title: %s
Transcript: %s
Date: %s/%s/%s
Link: %s
News: %s
Image: %s
Alt: %s
`, comic.Number, comic.Title, comic.SafeTitle, comic.Transcript, comic.Month, comic.Day, comic.Year, comic.Link, comic.News, comic.Image, comic.Alt)
}
