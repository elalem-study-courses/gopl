// package description
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func simpleErrorHandling(err error) {
	fmt.Printf("error while loading alexa %v\n", err)
}

func main() {
	start := time.Now()

	alexaMilSiteCsv, err := os.Open("/tmp/top-1m.csv")
	if err != nil {
		simpleErrorHandling(err)
	}
	lines, err := csv.NewReader(alexaMilSiteCsv).ReadAll()

	if err != nil {
		simpleErrorHandling(err)
	}

	ch := make(chan string)

	outputFile := createOutputFile()

	output := bufio.NewWriter(outputFile)

	for _, line := range lines {
		url := line[1]
		url = prepareURL(url)
		go fetch(url, ch)
	}

	for range lines {
		output.WriteString(<-ch)
		output.WriteByte('\n')
	}
	output.WriteString(fmt.Sprintf("%0.2fs elapsed", time.Since(start).Seconds()))
	output.Flush()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error occured for %s: %v", url, err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error occured for %s: %v", url, err)
		return
	}

	ch <- fmt.Sprintf("%0.2fs %7d %s", time.Since(start).Seconds(), nbytes, url)

}

func prepareURL(url string) string {
	if !hasHTTPKeyword(url) {
		url = "http://" + url
	}
	return url
}

func hasHTTPKeyword(url string) bool {
	if strings.Contains(url, "http://") ||
		strings.Contains(url, "https://") {
		return true
	}
	return false
}

func createOutputFile() *os.File {
	filename := strconv.FormatInt(time.Now().Unix(), 16)
	file, err := os.Create(fmt.Sprintf("%s.log", filename))
	if err != nil {
		fmt.Printf("Couldn't create output file %v\n", err)
		os.Exit(1)
	}
	return file
}
