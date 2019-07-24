package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helpHandler)
	http.HandleFunc("/collections", collectionsHandler)
	http.HandleFunc("/collections/", collectionHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
