// package description
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, nil
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "." || local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)

	if closeError := f.Close(); err == nil {
		err = closeError
	}

	return local, n, err
}

func main() {
	filename, n, err := fetch(os.Args[1])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Written %d in file %s\n", n, filename)
	}
}
