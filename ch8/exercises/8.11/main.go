package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	startSignal = make(chan struct{})
	ctx, cancel = context.WithCancel(context.Background())
)

func fireRequest(url string) (int64, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return -1, err
	}

	<-startSignal
	startTime := time.Now().UnixMilli()

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return -1, err
	}

	cancel()

	endTime := time.Now().UnixMilli()

	return endTime - startTime, nil
}

func main() {
	flag.Parse()

	var wg sync.WaitGroup

	for _, arg := range flag.Args() {
		wg.Add(1)
		go func(url string) {
			time, err := fireRequest(url)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Request %q took %d ms\n", url, time)
			}
			wg.Done()
		}(arg)
	}

	close(startSignal)
	wg.Wait()
}
