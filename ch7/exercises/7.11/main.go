package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/delete", db.delete)

	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}

}

func (db database) create(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, _ := strconv.ParseFloat(r.URL.Query().Get("price"), 64)

	db[item] = dollars(price)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok")
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, _ := strconv.ParseFloat(r.URL.Query().Get("price"), 64)

	if _, ok := db[item]; ok {
		db[item] = dollars(price)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	} else {
		http.Error(w, fmt.Sprintf("item %s not found", item), http.StatusNotFound)
	}
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	if _, ok := db[item]; ok {
		delete(db, item)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	} else {
		http.Error(w, fmt.Sprintf("item %s not found", item), http.StatusNotFound)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if price, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	} else {
		fmt.Fprintf(w, "%s\n", price)
	}

}
