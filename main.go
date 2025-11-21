package main

import (
	"fmt"
	"log"
	"net/http"
)

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You visited: %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", UrlHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
