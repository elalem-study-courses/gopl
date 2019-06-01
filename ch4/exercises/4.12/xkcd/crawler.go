package xkcd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/olivere/elastic"
)

var elasticClient *elastic.Client

func init() {
	var err error
	elasticClient, err = elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9300"),
	)

	if err != nil {
		log.Fatalf("init: %v\n", err)
	}

	_, err = elasticClient.CreateIndex("xkcd").Do(context.Background())
	if err != nil {
		log.Fatalf("init: %v\n", err)
	}
}

const (
	URL         = "https://xkcd.com/%d/info.0.json"
	UnkownError = -1
)

func Crawl() {
	for i := 1; true; i++ {
		comic, _, status := fetchComic(i)
		if status == http.StatusNotFound {
			log.Printf("Finished with all known comic exiting...")
			break
		} else if status == UnkownError {
			log.Println("Unknown error occured")
			break
		}
		err := indexComic(comic)
		if err != nil {
			log.Fatalf("Crawl: %v\n", err)
		}
	}
}

func fetchComic(id int) (*Comic, error, int) {
	url := fmt.Sprintf(URL, id)

	log.Printf("Quering %s\n", url)
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("fetchComic: %v\n", err)
		return nil, nil, UnkownError
	}
	defer resp.Body.Close()

	comic := new(Comic)

	if err := json.NewDecoder(resp.Body).Decode(comic); err != nil {
		log.Printf("fetchComic: %v\n", err)
		return nil, err, -1
	}

	return comic, nil, resp.StatusCode
}

func indexComic(comic *Comic) error {

	_, err := elasticClient.Index().
		Index("xkcd").
		Type("doc").
		Id(strconv.Itoa(comic.Num)).
		BodyJson(comic).
		Refresh("wait_for").
		Do(context.Background())

	if err != nil {
		return err
	}

	fmt.Printf("Indexed comic #%d successfully\n", comic.Num)
	return nil
}
