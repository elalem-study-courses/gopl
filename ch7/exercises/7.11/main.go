package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32
type database map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%0.2f", d)
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		http.Error(w, fmt.Sprintf("no such item: %q\n", item), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) create(w http.ResponseWriter, r *http.Request) {
	for key := range r.URL.Query() {
		if _, ok := db[key]; ok {
			continue
		}
		value := r.URL.Query().Get(key)
		price, err := strconv.ParseFloat(value, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Could not parse value %v: %v\n", value, err)
			return
		}
		db[key] = dollars(price)
	}
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	for key := range r.URL.Query() {
		if _, ok := db[key]; !ok {
			continue
		}
		value := r.URL.Query().Get(key)
		price, err := strconv.ParseFloat(value, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Could not parse value %v: %v\n", value, err)
			return
		}
		db[key] = dollars(price)
	}
}

func (db database) read(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("item")
	price, ok := db[key]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 could not find item %s", key)
	} else {
		fmt.Fprintf(w, "%s = %s\n", key, price)
	}
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("item")
	if _, ok := db[key]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 could not find item %s", key)
	} else {
		delete(db, key)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
