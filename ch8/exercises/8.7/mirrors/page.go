package mirrors

import "net/url"

type Page struct {
	*url.URL
}

func newPage(link string) (*Page, error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return nil, err
	}
	return &Page{
		parsedURL,
	}, nil
}
