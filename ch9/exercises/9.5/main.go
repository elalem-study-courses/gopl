package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

/**
Around 2894171 op/sec
*/

var sent, received int

func pingPong(ch chan struct{}, first bool) {
	isReceived := false
	for {
		if first || isReceived {
			ch <- struct{}{}
		}
		sent++
		<-ch
		received++
		isReceived = true
	}
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ch := make(chan struct{})
	go pingPong(ch, true)
	go pingPong(ch, false)
	go func() {
		now := time.Now()
		for {
			time.Sleep(1 * time.Second)
			since := time.Since(now)
			fmt.Printf("\r%d op/sec", int(math.Ceil(float64(sent)/since.Seconds())))
		}
	}()
	wg.Wait()
}
