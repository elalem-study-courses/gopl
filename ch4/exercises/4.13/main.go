package main

import (
	"os"

	"./omdbapi"
)

func main() {
	omdbapi.SearchMovies(os.Args[1])
}
