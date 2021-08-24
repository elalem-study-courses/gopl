package main

import (
	"fmt"
	"log"
	"net/http"

	"./routes"
)

// /issues
// /issues/:number
func main() {
	http.Handle("/public", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.Handle("/", routes.Routes())
	// routes.Routes()
	fmt.Println("Running server on localhost:3000...")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
