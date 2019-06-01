package xkcd

type Comic struct {
	Month, Year string
	Num         int
	Link        string
	News        string
	SafeTitle   string `json:"safe_title"`
	Transcript  string
	Alt         string
	Image       string `json:"img"`
	Title       string
	Day         string
}
