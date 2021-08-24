package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	defer resp.Body.Close()

	local := strings.TrimSpace(path.Base(resp.Request.URL.Path))
	if local == "." || local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	defer func() {
		if closeError := f.Close(); closeError != nil {
			err = closeError
		}
	}()

	n, err = io.Copy(f, resp.Body)

	return local, n, err
}

func main() {
	fmt.Println(fetch(os.Args[1]))
}
