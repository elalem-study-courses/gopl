// package description
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	file := initializeOutputFile()
	output := bufio.NewWriter(file)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		// fmt.Println(<-ch)
		output.WriteString(<-ch)
		output.WriteByte('\n')
	}

	fmt.Printf("%0.2fs elapsed\n", time.Since(start).Seconds())
	output.Flush()
}

func initializeOutputFile() *os.File {

	filename := strconv.FormatInt(time.Now().Unix(), 16)
	file, err := os.Create(fmt.Sprintf("%s.log", filename))
	if err != nil {
		fmt.Printf("Couldn't create output file %v\n", err)
		os.Exit(1)
	}
	return file
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("Error while fetching %s: %v", url, err)
		return
	}

	ch <- fmt.Sprintf("%0.2fs %7d %s", time.Since(start).Seconds(), nbytes, url)
}
