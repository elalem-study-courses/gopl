package mirrors

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"regexp"
	"sync"
)

const (
	ConcurrencyFactor               = 100
	fileOperationsConcurrencyFactor = 1000
)

var (
	hostRegex                  *regexp.Regexp
	fileOperationsSynchronizer = make(chan struct{}, fileOperationsConcurrencyFactor)
)

type Mirrorer struct {
	visited      map[string]bool
	linksVisited int
	Host         *url.URL
}

func (m *Mirrorer) MarkAsReceived(link string) {
	m.visited[link] = true
}

func (m *Mirrorer) IsVisited(link string) bool {
	_, ok := m.visited[link]
	return ok
}

func (m *Mirrorer) Run() {
	hostRegex = regexp.MustCompile(fmt.Sprintf(`http(s)?:\/\/%s`, m.Host.Host))
	workList := make(chan *Links)
	go func() {
		link, err := url.Parse(m.Host.String())
		if err != nil {
			log.Fatal(err)
		}
		links := &Links{urls: []*url.URL{link}}
		workList <- links
	}()

	n := 1

	var wg sync.WaitGroup

	for ; n > 0; n-- {
		links := <-workList

		if links.src != nil {
			wg.Add(1)

			go func(links *Links) {
				folderPath := generatePath(links.src)
				absoluteFolderPath, _ := filepath.Abs(folderPath)
				fmt.Println(absoluteFolderPath)
				contents := hostRegex.ReplaceAll([]byte(links.pageContent), []byte(absoluteFolderPath))
				fileOperationsSynchronizer <- struct{}{}
				createFolderForLink(folderPath)
				writeFile(filepath.Join(folderPath, "index.html"), contents)
				<-fileOperationsSynchronizer
				wg.Done()
			}(links)
		}

		for _, link := range links.urls {
			if link.Host == m.Host.Host && !m.IsVisited(link.String()) {
				m.MarkAsReceived(link.String())
				n++
				go func(link string) {
					links, err := crawl(link)
					if err != nil {
						log.Println(err)
						return
					}

					workList <- links
				}(link.String())
			}
		}
	}

}

func crawl(link string) (*Links, error) {
	// fmt.Printf("Crawling %v\n", link)
	links, err := extractLinks(link)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return links, nil
}

func NewMirrorer(host string) (*Mirrorer, error) {
	parsedHost, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse host %q: %v", host, err)
	}

	if len(parsedHost.Scheme) == 0 {
		parsedHost.Scheme = "https"
	}

	return &Mirrorer{
		visited:      make(map[string]bool),
		linksVisited: 0,
		Host:         parsedHost,
	}, nil
}
