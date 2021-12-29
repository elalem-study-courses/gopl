package main

import (
	"flag"
	"github.com/mohamed-elalem/gopl/ch8/exercises/8.7/mirrors"
	"log"
)

var host = flag.String("host", "", "The host to be mirrored locally")

func main() {
	flag.Parse()

	if *host == "" {
		log.Fatal("Host can't be empty")
	}

	mirrorer, err := mirrors.NewMirrorer(*host)
	if err != nil {
		log.Fatal(err)
	}

	mirrorer.Run()
}
