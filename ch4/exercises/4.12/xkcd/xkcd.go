package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	LastId = 2500
)

var (
	fpath string
)

var registry Registry

func init() {
	registry = newRegistry()
}

func init() {
	fpath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	fpath = filepath.Join(fpath, "comics.json")
}

func Index() {
	for i := 1; i <= LastId; i++ {
		comic, err := fetchById(i)
		if err != nil {
			log.Println(err)
			continue
		}

		registry.add(comic)

		if i%100 == 0 {
			fmt.Printf("\rIndexed %d comics....", i)
		}
	}

	fmt.Printf("\nWriting contents to %s\n", fpath)

	fileContent, err := json.MarshalIndent(registry, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(fpath, fileContent, 0644)
}

func fetchById(id int) (*Comic, error) {
	res, err := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", id))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Got status %s", res.Status)
	}

	comic := new(Comic)

	if err := json.NewDecoder(res.Body).Decode(comic); err != nil {
		return nil, err
	}

	return comic, nil
}

func FetchComic(id string) (*Comic, error) {
	loadIndex()
	if comic, ok := registry.get(id); ok {
		return comic, nil
	} else {
		return nil, fmt.Errorf("Unable to find comic " + id)
	}
}

func loadIndex() {
	fileContent, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(fileContent, &registry); err != nil {
		log.Fatal(err)
	}
}
