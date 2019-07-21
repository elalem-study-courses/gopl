package main

import (
	"fmt"
	"log"
	"os"

	"./links"
)

var (
	tokens = make(chan struct{}, 20)
	done   = make(chan struct{})
)

func crawl(url string, cancel <-chan struct{}) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url, cancel)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
		close(worklist)
		close(unseenLinks)
	}()

	go func() {
		worklist <- os.Args[1:]
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {

				foundLinks := crawl(link, done)
				go func() {
					select {
					case <-done:
						for range worklist {
						}
						return
					case worklist <- foundLinks:
					}
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
	InnerLoop:
		for _, link := range list {
			if !seen[link] && !cancelled() {
				seen[link] = true
				select {
				case unseenLinks <- link:
				case <-done:
					for range unseenLinks {
					}
					break InnerLoop
				}
			}
		}
	}
}
