package main

import (
	"flag"
	"fmt"
	"log"

	"./xkcd"
)

var action = flag.String("a", "", "Action to perform [index, fetch]")
var id = flag.String("id", "1", "ID of the comic")

func main() {
	flag.Parse()

	switch *action {
	case "index":
		xkcd.Index()
	case "fetch":
		comic, err := xkcd.FetchComic(*id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(comic)
	}
}
