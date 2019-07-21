package mirrors

import "golang.org/x/net/html"

type Mirror struct {
	filename    string
	originalURL string
	html        *html.Node
}

func (m *Mirror) Save() error {
	return nil
}

func Run(site string) {
	crawler := NewCrawler(site)
	crawler.crawl()
}
