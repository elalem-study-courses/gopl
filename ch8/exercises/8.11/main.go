package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

var done = make(chan struct{})

func main() {
	for _, link := range os.Args[1:] {
		wg.Add(1)
		go func(link string) {
			defer wg.Done()
			duration, err := benchmark(link)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			fmt.Printf("Link %s scored %v\n", link, duration)
		}(link)
	}

	wg.Wait()
}

func benchmark(link string) (time.Duration, error) {
	start := time.Now()
	req, err := http.NewRequest(http.MethodGet, link, nil)
	req.Cancel = done
	if err != nil {
		return 0, err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	select {
	case <-done:
	default:
		close(done)
	}

	return time.Since(start), nil
}
