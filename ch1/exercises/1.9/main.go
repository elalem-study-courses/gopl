// Package description
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
		}
		status := resp.Status
		resp.Body.Close()

		fmt.Printf("Response returned with status %s\n", status)
		fmt.Println("----------------------------------")
		fmt.Printf("%s\n", b)
	}
}
