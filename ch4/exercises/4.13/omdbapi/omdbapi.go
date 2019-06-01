package omdbapi

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

var (
	ApiKey = os.Getenv("API_KEY")
)

const URL = "http://www.omdbapi.com/?t=%s&apikey=%s"

func SearchMovies(term string) {
	url := fmt.Sprintf(URL, term, ApiKey)

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalf("SearchMovies: %v\n", err)
	}
	defer resp.Body.Close()

	// if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	// 	fmt.Print(err)
	// }

	movie := new(Movie)

	if err := json.NewDecoder(resp.Body).Decode(movie); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Saving %s poster\n", movie.Title)

	saveImage(movie)
}

func saveImage(movie *Movie) {
	resp, err := http.Get(movie.Poster)

	if err != nil {
		log.Fatalf("saveImage: %v\n", err)
	}

	defer resp.Body.Close()

	os.Mkdir(movie.Title, 0755)

	imagePath := path.Join(".", movie.Title, path.Base(movie.Poster))

	imageFile, err := os.OpenFile(imagePath, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		log.Fatalf("saveImage: %v\n", err)
	}

	w := bufio.NewWriter(imageFile)

	if _, err := io.Copy(w, resp.Body); err != nil {
		log.Fatalf("saveImage: %v\n", err)
	}
}
