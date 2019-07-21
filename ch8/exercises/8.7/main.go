package main

import (
	"flag"
	"log"

	"./mirrors"
)

var site = flag.String("site", "", "The site to be locally mirrored")

func main() {
	flag.Parse()

	if *site == "" {
		log.Fatal("Site cannot be empty")
	}
	mirrors.Run(*site)
}
